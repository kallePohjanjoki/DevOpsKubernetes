package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"time"
)

func generateRandomString() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
	randomID := generateRandomString()

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339Nano), randomID)

		for range ticker.C {
			fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339Nano), randomID)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s: %s", time.Now().Format(time.RFC3339Nano), randomID)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
