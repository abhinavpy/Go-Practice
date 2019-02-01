package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, there, I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	//Note the use of _ to ignore the return value from loadPage
	//First, this function extracts the page title from r.URL.Path, the
	//path component of the request URL. The Path is re-sliced with
	//[len("/view/"):] to drop the leading "/view/" component of the request pat.
	//This is because the path will invariable begin with "/view/", which is
	//not part of the page's title.
	//The function then loads the page data, formats the page with a string of simple HTML, and writes it to w, the http.ResponseWriter.
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w,p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
