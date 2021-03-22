package solidGate

import (
	"encoding/json"
	"github.com/vladgoncharenko/notifier/common"
	"io/ioutil"
	"net/http"
)

var listNotifications []interface{}
var lastNotificationOnGate []interface{}

func SaveSolidGateProd(w http.ResponseWriter, r *http.Request) {
	var notific interface{}
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		common.ErrorHandler(err)
		common.ClearSlice(&listNotifications)
		err = json.Unmarshal(body, &notific)
		common.ErrorHandler(err)
		listNotifications = append(listNotifications, notific)
		w.Write([]byte(common.JsonStatusOk))
	}
	defer r.Body.Close()
}

func BackSolidGateProd(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		list, _ := json.Marshal(listNotifications)
		w.Write([]byte(list))
		listNotifications = nil
	}
	defer r.Body.Close()
}

func SaveSolidGateProdLast(w http.ResponseWriter, r *http.Request) {
	var notific interface{}
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		common.ErrorHandler(err)
		common.ClearSlice(&lastNotificationOnGate)
		err = json.Unmarshal(body, &notific)
		common.ErrorHandler(err)
		lastNotificationOnGate = append(lastNotificationOnGate, notific)
		w.Write([]byte(common.JsonStatusOk))
	}
	defer r.Body.Close()
}

func BackSolidGateProdLast(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		list, _ := json.Marshal(lastNotificationOnGate)
		w.Write([]byte(list))
	}
	defer r.Body.Close()
}
