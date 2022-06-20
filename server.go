package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	// method to save file
	filename := (*p).Title + ".txt"
	return os.WriteFile(filename, (*p).Body, 0600)
}

func loadPage(title string) (*Page, error) {
	//function to load pages
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test page %s", (*r).URL.Path)
}

func wikiViewHandler(w http.ResponseWriter, r *http.Request) {
	title := (*r).URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	http.HandleFunc("/test", testHandler)      // http routing
	http.HandleFunc("/view/", wikiViewHandler) // http routing

	log.Fatal(http.ListenAndServe(":8080", nil))
}
