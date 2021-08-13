package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("simple1.html"))
}
type Page struct {
	Title, Content string
}
func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin da castle. buffer",
	}
	var b bytes.Buffer
	err := t.Execute(&b, p)
	if err != nil{
		fmt.Fprint(w, "A error occured.")

	}
	b.WriteTo(w)
}
func main() {
	http.HandleFunc("/dates",displayPage)
	http.ListenAndServe(":8080",nil)
}
