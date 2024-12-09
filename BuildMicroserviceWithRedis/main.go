package main

import (
	application "GOTEST/BuildMicroserviceWithRedis/Application"
	"context"
	"fmt"
)

// Package chi is a small, idiomatic and composable router for building HTTP services .
func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}

}
