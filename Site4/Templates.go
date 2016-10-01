package main

import (
	"encoding/json"
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

// Post type
type Post struct {
	Title       string
	Author      string
	Body        string
	PublishDate string
}

// BlogViewModel type
type BlogViewModel struct {
	Blog  Blog
	Posts []Post
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func loadPosts() []Post {
	bytes, _ := ioutil.ReadFile("Site4/posts.json")
	var posts []Post
	json.Unmarshal(bytes, &posts)

	return posts
}

func handler(w http.ResponseWriter, r *http.Request) {
	blog := Blog{Title: "My Amazing Blog", Author: "Derek Jensen", Header: "Welcome to my amazing blog"}
	posts := loadPosts()
	viewModel := BlogViewModel{Blog: blog, Posts: posts}
	t, _ := template.ParseFiles("Site4/blog.html")
	t.Execute(w, viewModel)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
