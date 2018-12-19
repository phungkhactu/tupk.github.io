package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public"))) //public is the name of the view folder
	log.Fatal(http.ListenAndServe(":9090", nil))
}
