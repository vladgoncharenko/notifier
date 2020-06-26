package main

import (
	"encoding/json"
	"fmt"
	"github.com/vladgoncharenko/notifier/actions/gate"
	"github.com/vladgoncharenko/notifier/actions/solidGate"
	"github.com/vladgoncharenko/notifier/actions/vmpi"
	"github.com/vladgoncharenko/notifier/common"
	"github.com/vladgoncharenko/notifier/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var bodyToShow []string
var notificationToShow []string
var testNotifications []string
var testNotificationsTemp []models.Status

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
	http.HandleFunc("/notification", gate.Notification)
	http.HandleFunc("/notificationH", gate.NotificationHeader)
	http.HandleFunc("/shownotification", gate.ShowNotification)

	http.HandleFunc("/savenotification", gate.SaveNotifications)
	http.HandleFunc("/backnotification", gate.BackNotifications)

	http.HandleFunc("/saveSolid", solidGate.SaveSolidGateProd)
	http.HandleFunc("/backSolid", solidGate.BackSolidGateProd)

	//VMPI
	http.HandleFunc("/empty", vmpi.Empty)
	http.HandleFunc("/vmpiResponseAsMerchant", vmpi.VmpiResp)
	http.HandleFunc("/vmpiCheckRequestFromVisa", vmpi.VmpiCheckRequestFromVisa)
	http.HandleFunc("/vmpi/transaction-id", vmpi.VmpiResponseExtended)

	err := http.ListenAndServe(":9099", nil)

	if err != nil {
		fmt.Println(err)
	}
}

func start(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		common.ErrorHandler(err)
		w.Write(body)
	}
	defer r.Body.Close()
}

func ui(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		common.ErrorHandler(err)
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
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, str)
}

func delayRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		d := &RequestData{}
		err = json.Unmarshal(body, &d)
		common.ErrorHandler(err)
		log.Println(string(body))
		common.ErrorHandler(err)
		if d.Delay > 0 {
			time.Sleep(time.Duration(d.Delay) * time.Second)
		}
		w.Write([]byte(d.RespBody))
	}
	defer r.Body.Close()
}

func delayEleven(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		common.ErrorHandler(err)
		time.Sleep(11 * time.Second)
		w.Write(body)
	}
	defer r.Body.Close()
}
