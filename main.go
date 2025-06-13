package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
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

func cdnConnect(){
    // Step 2: Define the parameters for the session you want to create.
    key := os.Getenv("AWS_ACCESS_KEY_ID") // Access key pair. You can create access key pairs using the control panel or API.
    secret := os.Getenv("AWS_SECRET_ACCESS_KEY") // Secret access key defined through an environment variable.

    s3Config := &aws.Config{
        Credentials: credentials.NewStaticCredentials(key, secret, ""), // Specifies your credentials.
        Endpoint:    aws.String("https://nyc3.digitaloceanspaces.com"), // Find your endpoint in the control panel, under Settings. Prepend "https://".
        S3ForcePathStyle: aws.Bool(false), // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
        Region:      aws.String("us-east-1"), // Must be "us-east-1" when creating new Spaces. Otherwise, use the region in your endpoint, such as "nyc3".
    }

    // Step 3: The new session validates your request and directs it to your Space's specified endpoint using the AWS SDK.
    newSession := session.New(s3Config)
    s3Client := s3.New(newSession)

    // Step 4: Define the parameters of the object you want to upload.
    object := s3.PutObjectInput{
        Bucket: aws.String("personalphotos"), // The path to the directory you want to upload the object to, starting with your Space name.
        Key:    aws.String("hello-world.txt"), // Object key, referenced whenever you want to access this file later.
        Body:   strings.NewReader("Hello, World!"), // The object's contents.
        ACL:    aws.String("private"), // Defines Access-control List (ACL) permissions, such as private or public.
        Metadata: map[string]*string{ // Required. Defines metadata tags.
                                 "x-amz-meta-my-key": aws.String("your-value"),
                         },
    }

    // Step 5: Run the PutObject function with your parameters, catching for errors.
    _, err := s3Client.PutObject(&object)
    if err != nil {
        fmt.Println(err.Error())
    }
}

func main() {
	cdnConnect()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	tmpl = template.Must(template.ParseGlob(wd + "/templates/*.html"))


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
	http.HandleFunc("/cdn/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://personalphotos.nyc3.cdn.digitaloceanspaces.com"+r.URL.Path[4:], http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/assets/download-resume", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, wd+"/assets/Dexter_Williamson_Resume.pdf")
	})

	http.HandleFunc("/resume", func(w http.ResponseWriter, r *http.Request) {
	
		progLangs := map[string][]Blog{
			"Progs": {
				{Title: "Go", Picture: "/cdn/Go-Logo.png"},
				{Title: "Java", Picture:  "/cdn/Java-Logo.png"},
				{Title: "C++", Picture:  "/cdn/C++-Logo.png"},
				{Title: "Python", Picture: "/cdn/Python-Logo.png"},
			},
		}
		tmpl.ExecuteTemplate(w, "resume.html", progLangs)
	
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

	http.HandleFunc("/add-blog/", h2)

	log.Fatal(http.ListenAndServe(":8080",
		nil))

}
