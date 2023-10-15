package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles(("index.html")))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
	// w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error laoding .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
