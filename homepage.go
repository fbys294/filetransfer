package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	defaultport        = "9000"
	Static_Url  string = "/static/"
	Static_Root string = "static/"
)

type Context struct {
	Title  string
	Static string
	Date   string
	Time   string
}

// Home homepage index
func Home(w http.ResponseWriter, req *http.Request) {
	now := time.Now() // find the time right now

	context := Context{ //store the date and time in a struct
		Title: "Crunchy Transfer Application",
		Date:  now.Format("02-01-2006"),
		Time:  now.Format("15:04:05"),
	}
	render(w, "index", context)
}

// Send Page
func Send(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Send"}
	render(w, "send", context)
}

// Receive Page
func Received(w http.ResponseWriter, req *http.Request) {
	context := Context{Title: "Received"}
	render(w, "received", context)
}

func render(w http.ResponseWriter, tmpl string, context Context) {
	context.Static = Static_Url
	tmpl_list := []string{"template/base.html",
		fmt.Sprintf("template/%s.html", tmpl)}

	t, err := template.ParseFiles(tmpl_list...) //parse the html files
	if err != nil {                             // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, context) //execute the template and pass it the  struct to fill in the gaps
	if err != nil {             // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(Static_Url):]
	if len(static_file) != 0 {
		f, err := http.Dir(Static_Root).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+vn", defaultport)
		port = defaultport
	}
	http.HandleFunc("/", Home)
	http.HandleFunc("/send/", Send)
	http.HandleFunc("/received/", Received)
	http.HandleFunc(Static_Url, StaticHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
