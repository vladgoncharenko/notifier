package solidGate

import (
	"encoding/json"
	"github.com/vladgoncharenko/notifier/models"
	"io/ioutil"
	"log"
	"net/http"
)

var listNotificationsProd []models.StatusSolidGate
var listNotificationsStage []models.StatusSolidGate

func SaveSolidGateProd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)

		body, err := ioutil.ReadAll(r.Body)

		errorHandeler(err)

		if len(listNotificationsProd) > 100 {
			listNotificationsProd = nil
		}

		notific := &models.StatusSolidGate{}

		err = json.Unmarshal(body, &notific)

		errorHandeler(err)

		listNotificationsProd = append(listNotificationsProd, *notific)

		w.Write([]byte("{\"status\":\"ok\"}"))
	}
	defer r.Body.Close()
}

func BackSolidGateProd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		list, _ := json.Marshal(listNotificationsProd)

		w.Write([]byte(list))
		listNotificationsProd = nil
	}
	defer r.Body.Close()
}

func SaveSolidGateStage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)

		body, err := ioutil.ReadAll(r.Body)

		errorHandeler(err)

		if len(listNotificationsStage) > 100 {
			listNotificationsStage = nil
		}

		notific := &models.StatusSolidGate{}

		err = json.Unmarshal(body, &notific)

		errorHandeler(err)

		listNotificationsStage = append(listNotificationsStage, *notific)

		w.Write([]byte("{\"status\":\"ok\"}"))
	}
	defer r.Body.Close()
}

func BackSolidGateStage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		list, _ := json.Marshal(listNotificationsStage)

		w.Write([]byte(list))
		listNotificationsStage = nil
	}
	defer r.Body.Close()
}


func errorHandeler(err error) {
	if err != nil {
		log.Print(err)
	}
}
