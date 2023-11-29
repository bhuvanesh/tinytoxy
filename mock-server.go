package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from mock server!!!\n"))
}

func mockServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", mockHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:9898",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
