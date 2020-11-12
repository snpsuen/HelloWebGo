package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public_html")))

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
                http.ServeFile(w, r, "./public_html/hello.html")
        })

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there! ...\n")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
