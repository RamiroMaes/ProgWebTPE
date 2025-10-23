package main

import (
	"log"
	"os"

	"ejemplo.com/mi-proyecto-go/server"
)

func main() {
	connStr := "user=postgres password=XYZ dbname=tpespecial sslmode=disable"
	addr := ":8080"
	if env := os.Getenv("HTTP_ADDR"); env != "" {
		addr = env
	}
	if err := server.StartServer(connStr, addr); err != nil {
		log.Fatalf("error en servidor: %v", err)
	}
}
