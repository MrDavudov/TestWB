package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/MrDavudov/TestWB/internal/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.TextFormatter))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// Запуск сервера
	srv := new(handler.Server)
	if err := srv.Start(ctx); err != nil {
		logrus.Fatalf("errors occured while running http server: %s", err)
	}
}