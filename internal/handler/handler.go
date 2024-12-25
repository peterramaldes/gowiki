package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"github.com/peterramaldes/gowiki/internal/page"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		title := m[2]
		fn(w, r, title)
	}
}

func View() http.HandlerFunc {
	return makeHandler(viewHandler)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.New(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func Edit() http.HandlerFunc {
	return makeHandler(editHandler)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := page.New(title)
	if err != nil {
		p = &page.Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func Save() http.HandlerFunc {
	return makeHandler(saveHandler)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &page.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("_content/edit.tmpl", "_content/view.tmpl"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *page.Page) {
	filename := fmt.Sprintf("%s.tmpl", tmpl)
	err := templates.ExecuteTemplate(w, filename, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
