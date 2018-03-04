//// Following wiki example from go site article.
package main

// Godoc only uses the top line when I say package main.
// I need to structure this as a package to get the function by function docs.
// Also need to make func public to get godocs to appear
//package wiki

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		fmt.Println("Encountered an error:")
		fmt.Println(err)
	}
}

type Page struct {
	Title string
	Body  []byte
}

// Save writes the body of a page to a file named p.Title.txt.
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage read a file named title.txt and returns a Page.
// Returns an error if the file could not be read.
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	check(err)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	check(err)
	t.Execute(w, p)
}

// ViewHandler takes a url, finds the part after "/view/" and uses that as a
// title for a page to load. It creates a Page object with that name, the tries to
// load the page. It returns html with the title and body of the page.
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := LoadPage(title)
	check(err)

	RenderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	check(err)
	if err != nil {
		p = &Page{Title: title}
	}

	RenderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/view/", ViewHandler)
	http.HandleFunc("/edit/", EditHandler)
	//http.HandleFunc("/save/", SaveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
