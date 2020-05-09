package vmpi

import (
	"github.com/vladgoncharenko/notifier/common"
	"io/ioutil"
	"log"
	"net/http"
)

func Empty(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		common.ErrorHandler(err)
		w.Write(nil)
	}
	defer r.Body.Close()
}
