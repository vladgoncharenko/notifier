package solidGate

import (
	"encoding/json"
	"github.com/vladgoncharenko/notifier/common"
	"io/ioutil"
	"net/http"
)

var listNotifications []interface{}

func SaveSolidGateProd(w http.ResponseWriter, r *http.Request) {
	var notific interface{}
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		common.ErrorHandler(err)
		common.ClearSlice(listNotifications)
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
