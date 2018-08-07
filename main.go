package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func clockHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/clock.html.tpl"))

	if err := t.ExecuteTemplate(w, "clock.html.tpl", time.Now()); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/clock", clockHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./doc"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
