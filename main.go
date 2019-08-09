package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var bodyToShow [] string
var notificationToShow [] string
var testNotifications [] string

type RequestData struct {
	RespBody string `json:"respBody"`
	Delay    int    `json:"delay"`
}

func main() {
	http.HandleFunc("/", start)
	http.HandleFunc("/ui", ui)
	http.HandleFunc("/show", show)
	http.HandleFunc("/delay", delayRequest)
	http.HandleFunc("/delay11", delayEleven)
	http.HandleFunc("/notification", notification)
	http.HandleFunc("/shownotification", showNotification)
	http.HandleFunc("/savenotification", saveNotifications)
	http.HandleFunc("/backnotification", backNotifications)

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
		if len(bodyToShow) > 100 {
			bodyToShow = nil
		}
		bodyToShow = append(bodyToShow, string(body))
	}
	w.Write([]byte("127.0.0.1:9099/show"))
	defer r.Body.Close()

}

func show(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var str string

	for _, res := range bodyToShow {
		str += "<span>" + res + "</span>" + "<p>"
	}
	bodyToShow = nil

	fmt.Println(str)

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, str)
}

func delayRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		d := &RequestData{}
		err = json.Unmarshal(body, &d)

		if err != nil {
			log.Print(err)
		}
		log.Println(string(body))
		if err != nil {
			log.Print(err)
		}

		if d.Delay > 0 {
			time.Sleep(time.Duration(d.Delay) * time.Second)
		}
		w.Write([]byte(d.RespBody))
	}
	defer r.Body.Close()
}

func delayEleven(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		if err != nil {
			log.Print(err)
		}
		time.Sleep(11 * time.Second)
		w.Write(body)
	}
	defer r.Body.Close()
}

func notification(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		if err != nil {
			log.Print(err)
		}

		if len(notificationToShow) > 100 {
			notificationToShow = nil
		}
		notificationToShow = append(notificationToShow, string(body))
		w.Write([]byte("{\"status\":\"ok\"}"))
	}
	defer r.Body.Close()
}

func showNotification(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var str string

	for _, res := range notificationToShow {
		str += "<span>" + res + "</span>" + "<p>"
	}
	bodyToShow = nil
	notificationToShow = nil
	fmt.Println(str)

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, str)
}

func saveNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		if err != nil {
			log.Print(err)
		}

		if len(testNotifications) > 100 {
			testNotifications = nil
		}
		testNotifications = append(testNotifications, string(body))
		w.Write([]byte("{\"status\":\"ok\"}"))
	}
	defer r.Body.Close()
}

func backNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		list, _ := json.Marshal(testNotifications)
		w.Write([]byte(list))
	}
	defer r.Body.Close()
}
