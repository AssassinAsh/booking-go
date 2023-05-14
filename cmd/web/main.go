package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AssassinAsh/booking-go/pkg/cache"
	"github.com/AssassinAsh/booking-go/pkg/config"
	"github.com/AssassinAsh/booking-go/pkg/handlers"
	"github.com/AssassinAsh/booking-go/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig

const portNumber = ":8080"

// main function of the application
func main() {

	// Set to true for production
	app.IsProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Secure = app.IsProduction
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	tc, err := cache.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	fmt.Printf("Service started on port: %s", portNumber)

	app.UserCache = true

	http.ListenAndServe(portNumber, routes(&app))
}
