package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phc-dm/server-poisson/utils"
)

type object map[string]interface{}

func init() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Printf("MODE = %v (default: production)", os.Getenv("MODE"))
	log.Printf("PORT = %v (default: 8000)", os.Getenv("PORT"))

}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Templates & Renderer
	e.Renderer = utils.NewTemplateRenderer("base.html")

	// Static assets
	e.Static("/assets", "./assets")

	// Blog from "content-poisson"
	blogRoute := e.Group("/blog", middleware.Static("./blog"))
	{
		blogRoute.GET("", func(c echo.Context) error {
			return c.Redirect(http.StatusPermanentRedirect, "/blog/")
		})
		blogRoute.GET("/", func(c echo.Context) error {
			return c.File("./blog/index.html")
		})
	}

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home.html", object{})
	})
	e.GET("/utenti", func(c echo.Context) error {
		return c.Render(http.StatusOK, "utenti.html", object{})
	})
	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", object{})
	})

	port, ok := strconv.Atoi(os.Getenv("PORT"))
	if ok != nil {
		port = 8000
	}

	log.Print(port)

	// Start server
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(port)))
}
