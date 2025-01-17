package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

func connectToCDN() {
	endpoint := "https://personal-website.nyc3.cdn.digitaloceanspaces.com"
	accessKey := os.Getenv("DO801ZZNBCPTZVEDQMGY")
	secKey := os.Getenv("bOjBwcrRPXsCe8Ju8Ka+Bn4l6QJFOdANOau3qEg/LDE")

	// Initiate a client using DigitalOcean Spaces.

	newSession, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("us-east-1"),
		S3ForcePathStyle: aws.Bool(false), // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
	},
	)

	_, err = newSession.Config.Credentials.Get()

	if err != nil {
		fmt.Println("Error:", err)
	}

	s3Client := s3.New(newSession)

	result, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		// Handle error
		panic(err)
	}

	// Print out bucket names
	for _, b := range result.Buckets {
		println(*b.Name)
	}

	resp, _ := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{})

	for _, key := range resp.Contents {
		fmt.Println(*key.Key)
		fmt.Println("da")
		fmt.Println(key)
	}

	fmt.Println(resp.Contents)
	fmt.Println("blah")
}

func main() {
	connectToCDN()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	tmpl = template.Must(template.ParseGlob(wd + "/templates/*.html"))

	//h1 := func(w http.ResponseWriter, r *http.Request) {
	//
	//	blogs := map[string][]Blog{
	//		"Blogs": {
	//			{Title: "Go", Body: "Today feet good", Picture: "/assets/img/Go-Logo.png"},
	//			{Title: "Java", Body: "Today feet good", Picture: "/assets/img/Java-Logo.png"},
	//			{Title: "C++", Body: "Today feet good", Picture: "/assets/img/C++-Logo.png"},
	//			{Title: "Python", Body: "Today feet good", Picture: "/assets/img/Python-Logo.png"},
	//		},
	//	}
	//	tmpl.ExecuteTemplate(w, "index.html", blogs)
	//
	//}

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

	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "resume.html", nil)

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)

	})

	http.HandleFunc("/add-blog/", h2)

	//c := cron.New()
	//c.AddFunc("@every 10s", func() { http.HandleFunc("/spawnSVG/", spawnSVG) })
	//c.Start()
	log.Fatal(http.ListenAndServe(":8080",
		nil))

}
