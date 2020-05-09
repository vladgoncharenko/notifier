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

func SaveNotifications(w http.ResponseWriter, r *http.Request) {
	var notific interface{}
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		common.ErrorHandler(err)
		common.ClearSlice(notificationGate)
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
		w.Header().Add("content", "")
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
		common.ClearSlice(notificationGate)

		notificationGate = append(notificationGate, string(body))
		w.Write([]byte(common.JsonStatusOk))
	}
	defer r.Body.Close()
}

func ShowNotification(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		var str string
		for i, res := range notificationGate {
			str += "<span>" + "_______________________________________________________________" + "</span>" + "<p>"
			str += "<span>" + strconv.Itoa(i+1) + ")" + fmt.Sprint(res) + "</span>" + "<p>"
			str += "<span>" + "_______________________________________________________________" + "</span>" + "<p>"
		}
		notificationGate = nil
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, str)
	}
}
