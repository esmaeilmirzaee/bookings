package renders

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/esmaeilmirzaee/bookings/internal/config"
	"github.com/esmaeilmirzaee/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions template.FuncMap
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func DefaultTemplate(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "Warning")
	td.CSRFToken = nosurf.Token(r)

	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[templateName]
	if !ok {
		log.Fatal("Cannot read template", tc[templateName])
	}

	buf := new(bytes.Buffer)
	td = DefaultTemplate(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println("Error executing templates", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing logs to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		log.Fatal("Cannot read template pages.")
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Fatal("Could not parse templates", page)
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

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
