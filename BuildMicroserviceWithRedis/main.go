package main

import (
	application "GOTEST/BuildMicroserviceWithRedis/Application"
	"context"
	"fmt"
	"os"
	"os/signal"
)

// learning: https://www.youtube.com/watch?v=qCv-q37qjZU&t=828s
// Package chi is a small, idiomatic and composable router for building HTTP services .
func main() {
	app := application.New()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
