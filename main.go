package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	r.SkipClean(true)

	r.HandleFunc("/", start)
	r.HandleFunc("/ui", ui)
	r.HandleFunc("/show", show)
	r.HandleFunc("/delay", delayRequest)
	r.HandleFunc("/delay11", delayEleven)
	r.PathPrefix("/notification").HandlerFunc(gate.Notification)
	r.HandleFunc("/notificationH", gate.NotificationHeader)
	r.HandleFunc("/shownotification", gate.ShowNotification)
	r.HandleFunc("/notificationRedirect", gate.NotificationRedirect)

	r.PathPrefix("/lastNotification").HandlerFunc(gate.LastNotification)
	r.HandleFunc("/showLastNotifications", gate.ShowLastNotification)

	r.HandleFunc("/savenotification", gate.SaveNotifications)
	r.HandleFunc("/backnotification", gate.BackNotifications)

	r.HandleFunc("/saveSolid", solidGate.SaveSolidGateProd)
	r.HandleFunc("/backSolid", solidGate.BackSolidGateProd)

	r.HandleFunc("/saveSolidLast", solidGate.SaveSolidGateProdLast)
	r.HandleFunc("/backSolidLast", solidGate.BackSolidGateProdLast)

	//VMPI
	r.HandleFunc("/empty", vmpi.Empty)
	r.HandleFunc("/vmpiResponseAsMerchant", vmpi.VmpiResp)
	r.HandleFunc("/vmpiCheckRequestFromVisa", vmpi.VmpiCheckRequestFromVisa)
	r.HandleFunc("/vmpi/transaction-id", vmpi.VmpiResponseExtended)

	//http.Handle("/",r)
	err := http.ListenAndServe(":9099", r)

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

func toRevert() {
	fmt.Println("to_revert")
}
