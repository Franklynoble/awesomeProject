package main

import (
	"html/template"
	"net/http"
)
/*
This listing starts by creating a map to hold the templates. Each template is stored
separately from the others. The map is populated with the template instances by
using a key for the template name. When the templates user.html and page.html are
loaded, the base.html file is loaded with each of them. This allows for the inheritance
in each case.
 */


/*A map to store
templates in a
map of named
templates */
var t map[string]*template.Template


func init(){
	t = make(map[string]*template.Template)
	/*       Sets up the
	template map     */
	temp := template.Must(template.ParseFiles("base.html", "user.html"))
	t["user.html"] = temp
	/*      Loads templates along
	with base into the map line 18 - 23 */
	temp = template.Must(template.ParseFiles("base.html","page.html"))
	t["page.html"] = temp
}

type User struct{      //Data Object to Pass into templates
	Username, Name string
}
type Page struct{
	Title, Content string

}

// populate the dataSet for the Page
func displayPage(w http.ResponseWriter, r * http.Request) {
	p := &Page{
		Title: "An Example for Page",
		Content: "Have fun Stormin' da castle",
	}
	t["page.html"].ExecuteTemplate(w,"base",p) // invokes the template for the page

}
func displayUser(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "Swordsmith",
     	Name:"Inigo Monotoya",
	}
	t["user.html"].ExecuteTemplate(w,"base",u)
}
func main() {
	// serves Pages via the Built-in Web Server
	http.HandleFunc("/user",displayUser)
	http.HandleFunc("/",displayPage)
    http.ListenAndServe(":8080",nil)


}
