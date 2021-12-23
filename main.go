package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phc-dm/phc-server/config"
	"github.com/phc-dm/phc-server/util"
)

func main() {
	config.Load()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static content
	r.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	r.Handle("/blog/*", http.StripPrefix("/blog", http.FileServer(http.Dir("./blog/public"))))

	// Templates & Renderer
	renderer := NewTemplateRenderer("base.html")
	articleRegistry := NewArticleRegistry()
	articleRegistry.LoadAll()

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := renderer.Render(w, "home.html", util.H{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/link", func(w http.ResponseWriter, r *http.Request) {
		err := renderer.Render(w, "link.html", util.H{})
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

	r.Get("/news", func(w http.ResponseWriter, r *http.Request) {
		articles, err := articleRegistry.LoadAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := renderer.Render(w, "news.html", util.H{
			"Articles": articles,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/news/{article}", func(w http.ResponseWriter, r *http.Request) {
		article := chi.URLParam(r, "article")

		htmlSource, err := articleRegistry.Render(article)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := renderer.Render(w, "news-base.html", util.H{
			"ContentHTML": template.HTML(htmlSource),
		}); err != nil {
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
