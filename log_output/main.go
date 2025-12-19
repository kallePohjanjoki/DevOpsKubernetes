package main

import (
	"crypto/rand"
	"fmt"
	"log"
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

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339Nano), randomID)

	for range ticker.C {
		fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339Nano), randomID)
	}
}
