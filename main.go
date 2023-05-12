package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error occurred while executing the template: ", err)
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Service started on port: %s", portNumber))

	http.ListenAndServe(portNumber, nil)
}
