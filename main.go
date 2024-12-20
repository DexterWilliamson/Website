package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Blog struct {
	Title   string
	Body    string
	Picture string
}

type Categorys struct {
	CategoryName string
}

var tmpl *template.Template

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	tmpl = template.Must(template.ParseGlob(wd + "/templates/*.html"))

	h1 := func(w http.ResponseWriter, r *http.Request) {

		blogs := map[string][]Blog{
			"Blogs": {
				{Title: "Go", Body: "Today feet good", Picture: "/assets/img/Go-Logo.png"},
				{Title: "Java", Body: "Today feet good", Picture: "/assets/img/Java-Logo.png"},
				{Title: "C++", Body: "Today feet good", Picture: "/assets/img/C++-Logo.png"},
				{Title: "Python", Body: "Today feet good", Picture: "/assets/img/Python-Logo.png"},
			},
		}
		tmpl.ExecuteTemplate(w, "index.html", blogs)

	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")
		picture := r.PostFormValue("picture")
		err := tmpl.ExecuteTemplate(w, "blog-list-element", Blog{Title: title, Body: body, Picture: picture})
		if err != nil {
			log.Fatalln(err)
		}
	}

	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, wd+"/templates/css/styles.css")
	})
	http.HandleFunc("/templates/js/scripts.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, wd+"/templates/js/scripts.js")
	})
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, wd+r.URL.Path)
	})
	http.HandleFunc("/assets/download-resume", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, wd+"/assets/Dexter_Williamson_Resume.pdf")
	})

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-blog/", h2)

	//c := cron.New()
	//c.AddFunc("@every 10s", func() { http.HandleFunc("/spawnSVG/", spawnSVG) })
	//c.Start()
	log.Fatal(http.ListenAndServe(":8080",
		nil))

}
