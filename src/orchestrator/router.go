package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// makeRouter creates a new chi mux router
func makeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))

	r.Get("/pod", func(w http.ResponseWriter, r *http.Request) {
		cluster := newCluster()
		cluster.init()

		b, err := ioutil.ReadFile("pod.json")
		if err != nil {
			panic(err)
		}

		log.Println("attempting to make a pod")
		err = cluster.makePod(b)
		if err != nil {
			render.Status(r, 500)
			return
		}
		render.PlainText(w, r, "OK")
	})

	return r
}
