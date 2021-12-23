package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phc-dm/phc-server/config"
	"github.com/phc-dm/phc-server/templates"
	"github.com/phc-dm/phc-server/util"
)

func main() {
	config.Load()

	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	// Static content
	r.Handle("/public/*", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// Templates & Renderer
	renderer := templates.NewRenderer(
		"./views/",
		templates.File("./views/base.html"),
		templates.Pattern("./views/partials/*.html"),
	)
	newsArticlesRegistry := NewArticleRenderer("./news")

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
		articles, err := newsArticlesRegistry.LoadAll()
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
		articleName := chi.URLParam(r, "article")

		article, err := newsArticlesRegistry.Load(articleName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		html, err := article.Render()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := renderer.Render(w, "news-base.html", util.H{
			"Article":     article,
			"ContentHTML": template.HTML(html),
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
