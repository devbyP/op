package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8082"
	}
	mux := chi.NewMux()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	serv := http.Server{
		Addr:    port,
		Handler: mux,
	}
	log.Println("server start ", port)
	log.Fatalln(serv.ListenAndServe())
}
