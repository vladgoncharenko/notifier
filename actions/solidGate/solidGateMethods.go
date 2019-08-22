package solidGate

import (
	"encoding/json"
	"github.com/vladgoncharenko/notifier/models"
	"io/ioutil"
	"log"
	"net/http"
)

var listNotifications []models.StatusSolidGate
var tempList [] string

func SaveSolidGateNotific(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)

		body, err := ioutil.ReadAll(r.Body)

		errorHandeler(err)

		if len(listNotifications) > 100 {
			listNotifications = nil
		}

		notific := &models.StatusSolidGate{}

		err = json.Unmarshal(body, &notific)

		errorHandeler(err)

		listNotifications = append(listNotifications, *notific)

		w.Write([]byte("{\"status\":\"ok\"}"))
	}
	defer r.Body.Close()
}

func BackSolidGateNotific(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		list, _ := json.Marshal(listNotifications)

		w.Write([]byte(list))
		listNotifications = nil
	}
	defer r.Body.Close()
}

func SavekHeader(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"ok\"}"))
		header := r.Header

		for v := range header {
			tempList = append(tempList, v)
		}

	}
}

func BackHeader(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		list, _ := json.Marshal(tempList)

		w.Write([]byte(list))
		tempList = nil
	}
	defer r.Body.Close()
}

func errorHandeler(err error) {
	if err != nil {
		log.Print(err)
	}
}
