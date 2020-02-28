package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	handler := makeRouter()
	cluster := newCluster()
	cluster.init()

	bytes, err := ioutil.ReadFile("pod.json")
	if err != nil {
		panic(err)
	}

	go func() {
		log.Println("attempting to make a pod")
		cluster.makePod(bytes)
	}()

	log.Println("listening to port 8080")
	http.ListenAndServe(":8080", handler)
}
