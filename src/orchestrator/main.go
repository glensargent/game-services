package main

import (
	"log"
	"net/http"
)

func main() {
	handler := makeRouter()

	log.Println("listening to port 8080")
	http.ListenAndServe(":8080", handler)
}
