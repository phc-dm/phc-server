package main

import (
	"log"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/home.html")
}

func main() {

	http.HandleFunc("/", serveHome)

	log.Println("Listening on :8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
