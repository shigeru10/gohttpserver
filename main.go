package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func clockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		`<!DOCTYPE html>
    <html>
    <body>
      It's %d o'clock now.
    </body>
    </html>
  `, time.Now().Hour())
}

func main() {
	http.HandleFunc("/clock", clockHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./doc"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
