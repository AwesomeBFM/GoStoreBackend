package main

import (
	"fmt"
	"github.com/awesomebfm/go-store-backend/internal/database"
	"github.com/awesomebfm/go-store-backend/internal/router"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("[ENV] [ERROR] Failed to load dotenv! Error: %v\n", err)
		return
	}
	fmt.Println("[ENV] [NORMAL] Env loaded successfully!")

	// Connect to MongoDB
	err = database.Init(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE_NAME"))
	if err != nil {
		fmt.Printf("[DB] [ERROR] Failed to connect to MongoDB! Error: %v\n", err)
		return
	}
	fmt.Println("[DB] [NORMAL] MongoDB connection initialized!")

	// Shutdown handling (this comes before router as router is intentionally blocking the main thread)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalCh // Wait for Ctrl+C signal
		fmt.Println("Received shutdown signal. Properly shutting down the server!")

		// Close Database
		if err := database.Close(); err != nil {
			fmt.Printf("[DB] [ERROR] Failed to close MongoDB: %v\n", err)
		} else {
			fmt.Println("[DB] [NORMAL] MongoDB closed properly!")
		}

		// Add more shutdown tasks here

		os.Exit(0)
	}()

	// Setup Router
	apiRouter := router.NewRouter()

	err = apiRouter.Start()
	if err != nil {
		fmt.Printf("[ROUTER] [ERROR] Failed to start router! Error: %v\n", err)
		return
	}
}
