package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<!doctype html>\n"+
		"<html lang=\"en\">\n"+
		"<head>\n"+
		"<meta charset=\"utf-8\">\n"+
		"<title>My Amazing Blog!</title>\n"+
		"<meta name=\"description\" content=\"A really great blog about Go\">\n"+
		"<meta name=\"author\" content=\"Derek Jensen\">\n"+
		"</head>\n"+
		"<body>\n"+
		"<h1>Welcome to my amazing blog</h1>\n"+
		"</body>\n"+
		"</html>\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
