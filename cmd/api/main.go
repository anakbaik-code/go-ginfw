package main

import (
	"log"
	"go-fwgin/internal/bootstrap"
)

func main() {
	app,cleaup, err := bootstrap.InitializeApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	
	defer cleaup()

	// 2. Nyalakan server Gin
	if err := app.Start(); err != nil {
		log.Fatalf("Server stopped with error: %v", err)
	}
	
}