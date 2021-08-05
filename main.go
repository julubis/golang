package main

import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	port := "8080"
	http.HandleFunc("/", hello)
	log.Print("listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
