package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pulu pulu pulu"))
	}

	mux.HandleFunc("/", omah)
	mux.HandleFunc("/ola", ola)
	mux.HandleFunc("/wasu", wasu)
	mux.HandleFunc("/about", aboutHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func omah(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return
	}

	w.Write([]byte("Njajal gae website golang"))
}

func ola(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Apose kokondao yarabe sorendoreri muflenso paninema panipase"))
}

func wasu(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Sue ora jamu, jamu godong kates, sue ora ketemu, ketemu pisan jaluk pites"))
}
