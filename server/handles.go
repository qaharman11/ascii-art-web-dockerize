package server

import (
	"ascii-art-web/ascii-logic/funcs"
	"html/template"
	"net/http"
)

type Banner struct {
	OutputText string
}

type Urls struct {
	Home     string
	AsciiArt string
}

type Data struct {
	Text   Banner
	Routes Urls
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != Routes.Home {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	html, err := template.ParseFiles("templates/view.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405 not allowed", http.StatusMethodNotAllowed)
		return
	}
	data := Data{
		Text:   Banner{""},
		Routes: Routes,
	}

	err = html.Execute(w, data)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != Routes.AsciiArt {

		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {

		http.Error(w, "405 not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("signature")
	text2 := r.FormValue("value")
	if text2 != "standard" && text2 != "shadow" && text2 != "thinkertoy" {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}
	html, err := template.ParseFiles("templates/view.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	output, err := funcs.Logic(text, text2)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	banner := Banner{
		OutputText: output,
	}

	data := Data{
		Text:   banner,
		Routes: Routes,
	}

	err = html.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	html, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AuthorsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/authors" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	html, err := template.ParseFiles("templates/authors.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = html.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
