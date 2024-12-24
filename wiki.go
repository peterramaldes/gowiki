//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/peterramaldes/gowiki/internal/page"
)

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := page.LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]

	p, err := page.LoadPage(title)
	if err != nil {
		p = &page.Page{Title: title}
	}

	body := `
	<h1>Editing %s</h1>
	<form action="/save/%s" method="POST">
		<textarea name="body">%s</textarea>
		</br>
		<input type="submit" value="Save">
	</form>
	`

	fmt.Fprintf(w, body, p.Title, p.Title, p.Body)
}
