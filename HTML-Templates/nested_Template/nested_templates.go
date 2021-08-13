package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

func init(){
	t = template.Must(template.ParseFiles("index.html","head.html")) //Loads the two templates into a template Object
}
type Page struct{
	Title, Content string
}
func displayPage(w http.ResponseWriter, r * http.Request){
	 p := &Page{
	 	Title:"An Example With index and head html",
	 	Content: "Have fun Stormin' da castle",
	 }
	 /*
	 ExecuteTemplate provides control over the template file
	 when multiple ones are available.
	  */
	 t.ExecuteTemplate(w,"index.html",p) // invokes the Template with the Page Data
}
func main() {
	http.HandleFunc("/",displayPage)      // Serves the page on the built-in Web Server
	http.ListenAndServe(":8080",nil)

}
