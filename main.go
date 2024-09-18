package main

import (
	"html/template"
	"log"
	"net/http"
)

type Blog struct {
	Title string
	Body  string
}

type Categorys struct {
	CategoryName string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

}

func main() {

	h1 := func(w http.ResponseWriter, r *http.Request) {

		blogs := map[string][]Blog{
			"Blogs": {
				{Title: "sdf", Body: "Today feet good"},
				{Title: "sdf", Body: "Today feet good"},
				{Title: "zgdg", Body: "Today feet good"},
			},
		}

		tmpl.ExecuteTemplate(w, "index.html", blogs)

	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")
		err := tmpl.ExecuteTemplate(w, "blog-list-element", Blog{Title: title, Body: body})
		if err != nil {
			log.Fatalln(err)
		}
	}

	http.HandleFunc("/templates/css/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/css/styles.css")
	})
	http.HandleFunc("/templates/js/scripts.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/js/scripts.js")
	})
	http.HandleFunc("/img/bg-masthead.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/assets/img/bg-masthead.jpg")
	})
	http.HandleFunc("/img/bg-callout.jpg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/assets/img/bg-callout.jpg")
	})
	http.HandleFunc("/assets/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/assets/favicon.ico")
	})
	http.HandleFunc("/download-resume", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/assets/Dexter Williamson - Resume.pdf")
	})
	http.HandleFunc("/resume", h1)
	http.HandleFunc("/add-blog/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
