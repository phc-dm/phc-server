package main

import (
	"log"
	"net/http"

	"git.phc.dm.xxxxx.xx/server-poisson/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type object map[string]interface{}

func main() {
	config := LoadConfig()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static assets
	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	r.Handle("/blog/*", http.StripPrefix("/blog", http.FileServer(http.Dir("./blog/public"))))

	// Templates & Renderer
	renderer := utils.NewTemplateRenderer("base.html")

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := renderer.Render(w, "home.html", object{"Config": config})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/utenti", func(w http.ResponseWriter, r *http.Request) {
		err := renderer.Render(w, "utenti.html", object{"Config": config})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		err := renderer.Render(w, "login.html", object{"Config": config})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	log.Printf(`Starting server...`)
	http.ListenAndServe(config.Host, r)
}
