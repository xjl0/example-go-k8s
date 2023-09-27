package main

import (
	"context"
	"errors"
	"example-go-k8s/app/handlers"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const waitShutdown = 30

func main() {

	handler := handlers.NewHandler()

	svr := new(handlers.Server)

	idleConnClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, syscall.SIGTERM, os.Interrupt)

		<-sigint

		slog.Info("service interrupt received")

		ctx, cancel := context.WithTimeout(context.Background(), waitShutdown*time.Second)
		defer cancel()

		slog.Info("soft shutdown start")

		go func() {
			for i := waitShutdown; i > 0; i-- {
				select {
				case <-ctx.Done():
					return
				default:
					slog.Info("shutdown in", slog.Int("seconds", i))
					time.Sleep(time.Second)
				}
			}
		}()

		if err := svr.Shutdown(ctx); err != nil {
			slog.Error("http server shutdown", slog.String("error", err.Error()))
		}

		slog.Info("soft shutdown complete")

		close(idleConnClosed)
	}()

	slog.Info("service started")

	if err := svr.Run("", "8080", handler.InitRoutes()); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("fatal http server failed to start", slog.String("error", err.Error()))
		}
	}

	<-idleConnClosed

	slog.Info("service stopped")
}
