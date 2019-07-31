package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var bodyToShow string

func main() {
	http.HandleFunc("/", start)
	http.HandleFunc("/ui", ui)
	http.HandleFunc("/show", show)

	err := http.ListenAndServe(":9099", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		if err != nil {
			log.Print(err)
		}
		w.Write(body)
	}
	defer r.Body.Close()
}

func ui(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		if err != nil {
			log.Print(err)
		}
		log.Println(string(body))
		bodyToShow = string(body)
	}
	w.Write([]byte("127.0.0.1:9099/show"))
	defer r.Body.Close()

}

func show(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	showStr := "<span>" +
		bodyToShow +
		"</span>"
	io.WriteString(w, showStr)
}
