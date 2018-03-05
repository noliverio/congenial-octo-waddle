package main

//package wiki

import (
	//	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestSave(t *testing.T) {
	testPage := &Page{Title: "TestPage", Body: []byte("Sample test page")}
	_ = testPage.Save()

	retrievedBody, _ := ioutil.ReadFile("TestPage.txt")

	if !reflect.DeepEqual(testPage.Body, retrievedBody) {
		t.Errorf("Body was incorrect.\n Expected: %s \n Got: %s \n", testPage.Body, retrievedBody)
	}
	_ = os.Remove("TestPage.txt")
}

func TestLoadPage(t *testing.T) {
	testPage := &Page{Title: "TestPage", Body: []byte("Sample test page")}
	_ = testPage.Save()

	loadedPage, _ := LoadPage("TestPage")

	// Verify that the title and body of the Page created from a file are correct.
	if !reflect.DeepEqual(testPage.Title, loadedPage.Title) {
		t.Errorf("Title was incorrect.\n Expected: %s \n Got: %s \n", testPage.Title, loadedPage.Title)
	}

	if !reflect.DeepEqual(testPage.Body, loadedPage.Body) {
		t.Errorf("Body was incorrect.\n Expected: %s \n Got: %s \n", testPage.Body, loadedPage.Body)
	}
	_ = os.Remove("TestPage.txt")
}

func TestViewHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost:8080/view/geraffes", nil)
	w := httptest.NewRecorder()
	ViewHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("View test page did not return %v", http.StatusOK)
	}

	testString := "<div>dumb long horses"
	if !strings.Contains(w.Body.String(), testString) {
		t.Errorf("View test page did not contain %v", testString)
	}

}

func TestSaveHandler(t *testing.T) {

}

func TestEditHandler(t *testing.T) {

}

func TestRenderTemplate(t *testing.T) {

}
