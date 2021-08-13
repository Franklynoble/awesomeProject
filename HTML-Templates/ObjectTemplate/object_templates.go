package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML


func init() {
	t = template.Must(template.ParseFiles("index.html","quote.html"))
}
/*
Types to store data
for templates with
differing and specific
properties
line for Page, and line for Quote
 */
type Page struct {
	Title string
	Content template.HTML
}
type Quote struct {
	Quote, Person string
}

func main() {
	q := &Quote{
		Quote: `You keep using that word. i do not think it means what you think it means`,
		Person: "Inigo Montayo",
	}
	//write template and data
	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html",q)
    qc = template.HTML(b.String())   // store quotes as HTML IN Global variable
	//Serves handler using built-in web server
	http.HandleFunc("/",displayPage)
	http.ListenAndServe(":8080",nil)

}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "A User",  //create dataset
		Content: qc ,    //  with quote  html
	}
	t.ExecuteTemplate(w, "index.html",p) // Write quote and page to web server output
}
