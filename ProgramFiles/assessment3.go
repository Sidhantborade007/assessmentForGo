package main

import (
	"log"
	"net/http"
)

func main() {
	var c = make(chan string)
	go createServer(c)
	log.Println("Listening on :3000...")
	c <- "./view"

}
func createServer(c chan string) {
	filePath := <-c
	fs := http.FileServer(http.Dir(filePath))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}