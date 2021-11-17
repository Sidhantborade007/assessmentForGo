package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	message := query.Get("message")
	if  len(message) == 0 {
		fmt.Fprintf(w, "<h1>Hello World</h1>")
	}
	w.WriteHeader(200)
	w.Write([]byte("<h1>" +message + "</h1>"))
}


func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
