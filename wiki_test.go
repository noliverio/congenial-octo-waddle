package main

//package wiki

import (
	"io/ioutil"
	"os"
	//	"net/http/httptest"
	//	"fmt"
	"reflect"
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

}
