package main

import (
	"html/template"
	"log"
	"net/http"

	inspiratio "github.com/log-ed/orazio/pkg/muse"
)

func main() {
	s := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	log.Fatal(s.ListenAndServe())
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.New("satirae")                  // Create a template.
	t, _ = t.ParseFiles("tmpl/satirae.html", nil) // Parse template file.
	content := inspiratio.GetInspiratio()         // Get content.
	t.Execute(w, content)                         // merge.
}
