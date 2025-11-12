package main

import (
	"log"

	"transcription-service/server"
)

func main() {
	log.Print("Starting Transcription Service")

	s, err := server.InitializeServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	err = s.Start()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
