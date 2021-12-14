package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phc-dm/server-poisson/config"
	"github.com/phc-dm/server-poisson/util"
)

func main() {
	config.Load()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static assets
	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	r.Handle("/blog/*", http.StripPrefix("/blog", http.FileServer(http.Dir("./blog/public"))))

	// Templates & Renderer
	renderer := NewTemplateRenderer("base.html")

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := renderer.Render(w, "home.html", util.H{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/utenti", func(w http.ResponseWriter, r *http.Request) {
		err := renderer.Render(w, "utenti.html", util.H{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		err := renderer.Render(w, "login.html", util.H{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Printf(`Starting server on %v...`, config.Host)
	http.ListenAndServe(config.Host, r)
}
