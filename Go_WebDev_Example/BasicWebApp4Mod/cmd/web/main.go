package main

import (
	"fmt"
	"github.com/Go_workspace/Go_WebDev_Example/BasicWebApp4Mod/pkg/config"
	"github.com/Go_workspace/Go_WebDev_Example/BasicWebApp4Mod/pkg/handlers"
	"github.com/Go_workspace/Go_WebDev_Example/BasicWebApp4Mod/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting the application at port: %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
