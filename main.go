package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	inspiratio "github.com/log-ed/orazio/pkg/muse"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("starting orazio...")
	t := template.New("satirae")
	t, _ = t.ParseFiles("tmpl/satirae.html")
	var out string
	muse := inspiratio.GetInspiratio()
	for _, m := range muse {
		log.Println("getting inspiration from muse: " + m)
		resp, err := http.Get(m)
		if err != nil {
			log.Println("an error occured: " + err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			out = out + " " + string(body)
		}
	}
	t.Execute(w, out)
}
