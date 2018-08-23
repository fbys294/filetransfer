package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	default_port = "9000"
)

// PageVariables variable for the webpage goes here
type PageVariables struct {
	Date string
	Time string
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+vn", default_port)
		port = default_port
	}
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// HomePage homepage of site
func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now() // find the time right now

	HomePageVars := PageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	if err != nil {                                // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
