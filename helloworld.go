package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func SetTarget(target string) {
	os.Setenv("TARGET", target)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	fmt.Fprintf(w, "Hello there %s!\n", target)
}

func main() {
	log.Print("Hello world sample started.")

	http.HandleFunc("/", Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	SetTarget("Cloud Foundry Attendees")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
