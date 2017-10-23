package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/moby/moby/pkg/namesgenerator"
)

func main() {
	log.Println("starting calliope...")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, namesgenerator.GetRandomName(0))
}
