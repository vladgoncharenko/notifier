package vmpi

import (
	"encoding/json"
	"github.com/vladgoncharenko/notifier/common"
	"github.com/vladgoncharenko/notifier/models"
	"io/ioutil"
	"log"
	"net/http"
)

func Empty(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		common.ErrorHandler(err)
		w.Write(nil)
	}
	defer r.Body.Close()
}

func VmpiResp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		var requestFromVmpi models.VmpiRequest
		err = json.Unmarshal(body, &requestFromVmpi)
		common.ErrorHandler(err)
		response := requestFromVmpi.GetResponseForVmpi()
		data, _ := json.Marshal(response)
		w.Write(data)
	}
	defer r.Body.Close()
}

func VmpiCheckRequestFromVisa(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var responseAsClient models.ResponseAsVmpiClient
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		body, err := ioutil.ReadAll(r.Body)
		common.ErrorHandler(err)
		responseAsClient.AddVisaRequest(string(body))
		data, _ := json.Marshal(responseAsClient)
		w.Write(data)
	}
	defer r.Body.Close()
}
