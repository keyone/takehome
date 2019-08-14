package main

import (
	"fmt"
	"log"
	"net/http"

	"takehome/handlers"
)

func main() {
	fmt.Println(handlers.Header)
	err := http.ListenAndServe(":8080", handlers.Handler())
	if err != nil {
		log.Fatalf("Cannot start server, %s", err.Error())
	}
}
