package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there %s, isn't Go amazing?!", r.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
