package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"git.phc.dm.xxxxx.xx/server-poisson/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Static("/blog", "./blog/public")

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
