package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AssassinAsh/going-go/pkg/cache"
	"github.com/AssassinAsh/going-go/pkg/config"
	"github.com/AssassinAsh/going-go/pkg/handlers"
	"github.com/AssassinAsh/going-go/pkg/render"
)

const portNumber = ":8080"

// main function of the application
func main() {

	var app config.AppConfig

	tc, err := cache.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	repo := handlers.NewRepository(&app)
	handlers.NewHandler(repo)

	fmt.Println(fmt.Sprintf("Service started on port: %s", portNumber))

	app.UserCache = true

	http.ListenAndServe(portNumber, routes())
}
