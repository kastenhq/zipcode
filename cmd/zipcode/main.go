package main

import (
	"log"

	"github.com/kastenhq/zipcode/pkg/zipcode"
)

func main() {
	s, err := zipcode.New()
	if err != nil {
		log.Fatalf("Server failed to start: %s", err.Error())
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Server returned error on Shutdown: %s", err.Error())
	}
	log.Println("Server shutdown")
}
