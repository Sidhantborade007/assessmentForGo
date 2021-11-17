package main

import (
	"log"
	"net/http"
)

func main() {
	var c = make(chan string)
	c <- "./view"
	go createServer2(c)
	log.Println("Listening on :3000...")

}
func createServer2(c chan string) {
	filePath := <-c
	fs := http.FileServer(http.Dir(filePath))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}