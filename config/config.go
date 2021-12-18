package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/phc-dm/server-poisson/util"
)

var Mode string
var Host string

var GitUrl string
var ForumUrl string
var Email string

func loadEnv(target *string, name, defaultValue string) {
	value := os.Getenv(name)
	if len(strings.TrimSpace(value)) == 0 {
		*target = defaultValue
	} else {
		*target = value
	}
	log.Printf("%s = %v", name, *target)
}

func Load() {
	godotenv.Load()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	loadEnv(&Mode, "MODE", "production")
	loadEnv(&Host, "HOST", "localhost:8080")
	loadEnv(&GitUrl, "GIT_URL", "https://git.example.org")
	loadEnv(&ForumUrl, "FORUM_URL", "https://forum.example.org")
	loadEnv(&Email, "EMAIL", "mail@example.org")
}

func Object() util.H {
	return util.H{
		"Mode":     Mode,
		"Host":     Host,
		"GitUrl":   GitUrl,
		"ForumUrl": ForumUrl,
		"Email":    Email,
	}
}
