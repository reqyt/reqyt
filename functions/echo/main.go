package main

import (
	"io"
	"log"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
	io.Copy(w, r.Body)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
