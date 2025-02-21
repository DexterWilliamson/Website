package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Photo struct {
	Title   string
	Picture string
}

var tmpl *template.Template

func main() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	tmpl = template.Must(template.ParseGlob(wd + "/templates/*.html"))

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

	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "resume.html", nil)

	})

	http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {

		entries, err := os.ReadDir(wd + "/assets" + r.URL.Path)
		if err != nil {
			log.Fatal(err)
		}
		newPhotoList := []Photo{}

		for _, entry := range entries {
			something := Photo{Title: strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())), Picture: entry.Name()}
			newPhotoList = append(newPhotoList, something)
		}
		photos := map[string][]Photo{
			"Photos": newPhotoList,
		}
		tmpl.ExecuteTemplate(w, "gallery.html", photos)

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)

	})

	//c := cron.New()
	//c.AddFunc("@every 10s", func() { http.HandleFunc("/spawnSVG/", spawnSVG) })
	//c.Start()
	log.Fatal(http.ListenAndServe(":8080",
		nil))

}
