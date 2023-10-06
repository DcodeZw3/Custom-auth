package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

//Home page handler
func Home(w http.ResponseWriter, r *http.Request){
	RenderTemplate(w, "index.html")
}

//about page handler
func About(w http.ResponseWriter, r *http.Request){
	RenderTemplate(w, "index.html")
}


//render template
func RenderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+ html)
	err := parsedTemplate.Execute(w,nil)
	if err!= nil {
        fmt.Println("Error parsing template", err)
		return
    }
}