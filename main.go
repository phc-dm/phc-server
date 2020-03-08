package main

import (
	"html/template"
	"log"
	"net/http"
)

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./views/home.html")
// }

var templates = template.Must(template.ParseFiles("views/base.html"))

func newTemplateHandler(file string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		t, _ := (template.Must(templates.Clone()).ParseFiles("views/" + file))
		t.ExecuteTemplate(w, "base", nil)
	})
}

func main() {

	mux := http.NewServeMux()

	// Serve homepage
	homeTmpl := newTemplateHandler("home.html")
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/" {
			homeTmpl.ServeHTTP(w, req)
		} else {
			http.NotFound(w, req)
		}
	})

	mux.Handle("/utenti", newTemplateHandler("utenti.html"))
	mux.Handle("/login", newTemplateHandler("login.html"))

	// Serve assets files
	fs := http.FileServer(http.Dir("./assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fs))

	log.Println("Listening on :8000...")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
