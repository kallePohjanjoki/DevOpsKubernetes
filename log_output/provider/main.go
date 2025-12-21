package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, _ := os.ReadFile("/usr/src/app/files/log.txt")
		w.Write(data)
	})
	http.ListenAndServe(":8000", nil)
}
