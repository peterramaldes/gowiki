package main

import (
	"log"
	"net/http"

	"github.com/peterramaldes/gowiki/internal/handler"
)

func main() {
	http.HandleFunc("/view/", handler.View())
	http.HandleFunc("/edit/", handler.Edit())
	http.HandleFunc("/save/", handler.Save())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
