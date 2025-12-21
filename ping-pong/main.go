package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/pingpong", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		currentValue := counter
		counter++
		mutex.Unlock()

		fmt.Fprintf(w, "pong %d", currentValue)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		currentValue := counter
		counter++
		mutex.Unlock()

		fmt.Fprintf(w, "pong %d", currentValue)
	})

	fmt.Printf("Ping-pong server started in port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
