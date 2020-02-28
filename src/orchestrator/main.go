package main

import (
	"log"
	"net/http"
)

func main() {
	handler := makeRouter()

	makeCluster()

	log.Println("listening to port 8080")
	http.ListenAndServe(":8080", handler)
}
