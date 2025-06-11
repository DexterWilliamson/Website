package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"math/rand"

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


	//h2 := func(w http.ResponseWriter, r *http.Request) {
	//	title := r.PostFormValue("title")
	//	body := r.PostFormValue("body")
	//	picture := r.PostFormValue("picture")
	//	err := tmpl.ExecuteTemplate(w, "blog-list-element", Blog{Title: title, Body: body, Picture: picture})
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//}

	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, wd+"/templates/css/styles.css")
	})
	http.HandleFunc("/templates/js/scripts.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, wd+"/templates/js/scripts.js")
	})
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, wd+r.URL.Path)
	})
	http.HandleFunc("/cdn/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://personalphotos.nyc3.cdn.digitaloceanspaces.com"+r.URL.Path[4:], http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/assets/download-resume", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, wd+"/assets/Dexter_Williamson_Resume.pdf")
	})

	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
	
		blogs := map[string][]Blog{
			"Blogs": {
				{Title: "Go", Body: "Today feet good", Picture: "/cdn/Go-Logo.png"},
				{Title: "Java", Body: "Today feet good", Picture:  "/cdn/Java-Logo.png"},
				{Title: "C++", Body: "Today feet good", Picture:  "/cdn/C++-Logo.png"},
				{Title: "Python", Body: "Today feet good", Picture: "/cdn/Python-Logo.png"},
			},
		}
		tmpl.ExecuteTemplate(w, "resume.html", blogs)
	
	})

	//http.HandleFunc("/gallery", func(w http.ResponseWriter, r *http.Request) {
//
	//	entries, err := os.ReadDir(wd + "/assets" + r.URL.Path)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	newPhotoList := []Photo{}
//
	//	for _, entry := range entries {
	//		something := Photo{Title: strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())), Picture: entry.Name()}
	//		newPhotoList = append(newPhotoList, something)
	//	}
	//	photos := map[string][]Photo{
	//		"Photos": newPhotoList,
	//	}
	//	tmpl.ExecuteTemplate(w, "gallery.html", photos)
//
	//})

	//http.HandleFunc("/add-blog/", h2)

	http.HandleFunc("/spawnSVG/", func(w http.ResponseWriter, r *http.Request) {
		filepaths := []string{"/assets/resume/zig_zag.svg", "/assets/resume/v.svg", "/assets/resume/just_o.svg", "/assets/resume/x.svg"}
		randomPicker := rand.Intn(len(filepaths))
		http.ServeFile(w, r, wd+filepaths[randomPicker])

	})
	log.Fatal(http.ListenAndServe(":8080",
		nil))

}
