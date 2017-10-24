package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	inspiratio "github.com/log-ed/orazio/pkg/muse"
)

func main() {
	log.Println("starting orazio...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("tmpl/satirae.html"))
	var out string
	muse := inspiratio.GetInspiratio()
	for _, m := range muse {
		// Each muse must be an HTTP endpoint in the format http://...
		log.Println("getting inspiration from muse: " + m)
		resp, err := http.Get(m)
		if err != nil {
			log.Println("an error occured: " + err.Error())
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("an error occured: " + err.Error())
			return
		}
		out = out + " " + string(body)
	}
	log.Println("render web page with text: " + out)
	t.Execute(w, out)
}
