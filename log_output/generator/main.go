package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

func main() {
	os.MkdirAll("/usr/src/app/files", 0755)

	for range time.Tick(5 * time.Second) {

		b := make([]byte, 16)
		rand.Read(b)
		id := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

		f, _ := os.OpenFile("/usr/src/app/files/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		f.WriteString(fmt.Sprintf("%s: %s\n", time.Now().Format(time.RFC3339), id))
		f.Close()

		fmt.Printf("Wrote: %s\n", id)
	}
}
