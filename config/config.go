package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var Mode string
var Host string

var GitUrl string
var ForumUrl string

func stringOrDefault(value, defaultValue string) string {
	if len(strings.TrimSpace(value)) == 0 {
		return defaultValue
	}

	return value
}

func LoadConfig() {
	godotenv.Load()

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	Mode = stringOrDefault(os.Getenv("MODE"), "production")
	log.Printf("MODE = %v", Mode)

	Host = stringOrDefault(os.Getenv("HOST"), "localhost:8080")
	log.Printf("HOST = %v", Host)

	GitUrl = stringOrDefault(os.Getenv("GIT_URL"), "git.phc.dm.unipi.it")
	log.Printf("GIT_URL = %v", GitUrl)

	ForumUrl = stringOrDefault(os.Getenv("FORUM_URL"), "forum.phc.dm.unipi.it")
	log.Printf("FORUM_URL = %v", ForumUrl)
}
