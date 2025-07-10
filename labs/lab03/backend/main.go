package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"lab03-backend/api"
	"lab03-backend/storage"
)

func main() {
	log.Println("Setting server up...")

	memory := storage.NewMemoryStorage()
	handler := api.NewHandler(memory)
	router := handler.SetupRoutes()

	wrappedRouter := api.ChainMiddleware(
		api.RecoveryMiddleware,
		api.LoggingMiddleware,
		api.CorsMiddleware,
	)(router)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      wrappedRouter,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	log.Printf("Server successfully started at http://localhost:%d", addr.Port)
	log.Printf("Full server address: http://%s", listener.Addr())

	log.Fatal(server.Serve(listener))
}
