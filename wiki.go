// basing around golang wiki example, modified to fit my purposes
package main

import (
	"fmt"
	"io/ioutil"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	check(err)
	return &Page{Title: title, Body: body}

}

func main() {
}
