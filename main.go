package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathroot = "."
	const port = "8080"

	log.Println("Starting server...")
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathroot)))

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Println("Server listening on :8080")
	log.Printf("Serving files from %s on port: %s\n", filepathroot, port)
	log.Fatal(server.ListenAndServe())

}
