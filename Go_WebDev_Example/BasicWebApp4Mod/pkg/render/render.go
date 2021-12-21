package render

import (
	"bytes"
	"fmt"
	"github.com/Go_workspace/Go_WebDev_Example/BasicWebApp4Mod/pkg/config"
	"github.com/Go_workspace/Go_WebDev_Example/BasicWebApp4Mod/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplates is rendering
func RenderTemplates(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from app config
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()
	}

	/*	tc, err := CreateTemplateCache()
		if err != nil{
			log.Fatal(err)
		}*/

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get the template from template cache")
	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser")
	}

	/*parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil{
		fmt.Println("error parsing template:", err)
		return
	}*/

}

// RenderTemplatesNew is new render template func
//func RenderTemplatesNew(w http.ResponseWriter, tmpl string) (map[string]*template.Template, error){

//CreateTemplateCache create template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}
