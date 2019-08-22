package gate

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var notificationGate [] interface{}

func SaveNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)

		body, err := ioutil.ReadAll(r.Body)

		errorHandeler(err)

		if len(notificationGate) > 100 {
			notificationGate = nil
		}

		var notific interface{}

		err = json.Unmarshal(body, &notific)

		errorHandeler(err)

		notificationGate = append(notificationGate, notific)

		w.Write([]byte("{\"status\":\"ok\"}"))
	}
	defer r.Body.Close()
}

func BackNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content", "")
		list, _ := json.Marshal(notificationGate)

		w.Write([]byte(list))
		notificationGate = nil
	}
	defer r.Body.Close()
}

func errorHandeler(err error) {
	if err != nil {
		log.Print(err)
	}
}
