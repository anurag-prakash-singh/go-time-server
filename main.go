package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
)

func timeRequestHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	fmt.Fprintf(w, "Current time is %s", now.Format(time.RFC822))
}

func main() {
	

	s := &http.Server {
		Addr: ":8080",
		Handler: http.HandlerFunc(timeRequestHandler),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

