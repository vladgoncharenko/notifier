package gate

import (
	"encoding/json"
	"fmt"
	"github.com/vladgoncharenko/notifier/common"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var notificationGate []interface{}
var notificationToShow []interface{}
var lastNotificationToShow []interface{}

func SaveNotifications(w http.ResponseWriter, r *http.Request) {
	var notific interface{}
	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		common.ErrorHandler(err)
		common.ClearSlice(&notificationGate)
		err = json.Unmarshal(body, &notific)
		common.ErrorHandler(err)
		notificationGate = append(notificationGate, notific)

		w.Write([]byte(common.JsonStatusOk))
	}
	defer r.Body.Close()
}

func BackNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		list, _ := json.Marshal(notificationGate)

		w.Write([]byte(list))
		notificationGate = nil
	}
	defer r.Body.Close()
}

func Notification(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)

		common.ErrorHandler(err)
		common.ClearSlice(&notificationToShow)

		notificationToShow = append(notificationToShow, string(body))
		w.Write([]byte(common.JsonStatusOk))
	}
	defer r.Body.Close()
}

func LastNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)

		common.ErrorHandler(err)
		common.ClearSlice(&lastNotificationToShow)

		lastNotificationToShow = append(lastNotificationToShow, string(body))
		w.Write([]byte(common.JsonStatusOk))
	}
	defer r.Body.Close()
}

func NotificationRedirect(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "127.0.0.1/notification", http.StatusSeeOther)
	}
	defer r.Body.Close()
}

func ShowNotification(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var str string
		for i, res := range notificationToShow {
			str += "<span>" + "_______________________________________________________________" + "</span>" + "<p>"
			str += "<span>" + strconv.Itoa(i+1) + ")" + fmt.Sprint(res) + "</span>" + "<p>"
			str += "<span>" + "_______________________________________________________________" + "</span>" + "<p>"
		}
		notificationToShow = nil
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, str)
	}
}

func ShowLastNotification(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var str string
		for i, res := range lastNotificationToShow {
			str += "<span>" + "_______________________________________________________________" + "</span>" + "<p>"
			str += "<span>" + strconv.Itoa(i+1) + ")" + fmt.Sprint(res) + "</span>" + "<p>"
			str += "<span>" + "_______________________________________________________________" + "</span>" + "<p>"
		}
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, str)
	}
}

func NotificationHeader(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		header := r.Header
		common.ErrorHandler(err)
		common.ClearSlice(&notificationToShow)
		notificationToShow = append(notificationToShow, header.Get("signature"))
		notificationToShow = append(notificationToShow, string(body))

		w.Write([]byte(common.JsonStatusOk))
	}
	defer r.Body.Close()
}
