package vmpi

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Empty(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		body, err := ioutil.ReadAll(r.Body)
		log.Println(string(body))
		if err != nil {
			log.Print(err)
		}
		w.Write(nil)
	}
	defer r.Body.Close()
}
