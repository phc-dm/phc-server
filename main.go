package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phc-dm/phc-server/articles"
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
	newsArticlesRegistry := articles.NewRegistry("./news")

	// Routes

	actuallyStaticRoutes := map[string]string{
		"/":       "home.html",
		"/link":   "link.html",
		"/utenti": "utenti.html",
		"/login":  "login.html",
	}

	for route, view := range actuallyStaticRoutes {
		localView := view
		r.Get(route, func(w http.ResponseWriter, r *http.Request) {
			err := renderer.Render(w, localView, util.H{})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		})
	}

	r.Get("/appunti", func(w http.ResponseWriter, r *http.Request) {
		searchQuery := ""

		keys, present := r.URL.Query()["q"]
		if present {
			searchQuery = keys[0]
		}

		err := renderer.Render(w, "appunti.html", util.H{
			"Query": searchQuery,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Get("/news", func(w http.ResponseWriter, r *http.Request) {
		articles, err := newsArticlesRegistry.GetArticles()
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
		articleID := chi.URLParam(r, "article")

		article, err := newsArticlesRegistry.GetArticle(articleID)
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

	log.Printf(`Starting server on %v...`, config.Host)
	http.ListenAndServe(config.Host, r)
}
