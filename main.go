package main

import (
	"auth/app"
	"auth/pkg/database"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitializeDB()
	appServer := app.NewApplication()

	go func() {
		if err := appServer.Listen(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	_ = appServer.Shutdown()

	fmt.Println("Running cleanup tasks...")

	database.DB.Close()
	fmt.Println("Fiber was successful shutdown.")
}
