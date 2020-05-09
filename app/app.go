package app

import (
	"encoding/gob"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	// Store is an api to get session storage
	Store *sessions.CookieStore
)

// Init configures Store and Cookies
func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}

	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	gob.Register(map[string]interface{}{})
	return nil
}
