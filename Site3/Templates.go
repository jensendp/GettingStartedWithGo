package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

// Blog type
type Blog struct {
	Title  string
	Author string
	Header string
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	blog := Blog{Title: "My Amazing Blog", Author: "Derek Jensen", Header: "Welcome to my amazing blog"}
	t, _ := template.ParseFiles("Site3/blog.html")
	t.Execute(w, blog)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
