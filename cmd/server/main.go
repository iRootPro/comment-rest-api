package main

import (
	"context"
	"fmt"

	"github.com/iRootPro/comment-rest-api/internal/db"
)

func Run() error {
	fmt.Println("Run services")

	database, err := db.NewDB()
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	ctx := context.Background()
	if err := database.Ping(ctx); err != nil {
		return err
	}

	fmt.Println("Successfully connected to database")

	return nil
}

func main() {
	fmt.Println("Start server")
	if err := Run(); err != nil {
		fmt.Println(err.Error())
	}
}
