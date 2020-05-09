package common

import "log"

func ErrorHandler(err error) {
	if err != nil {
		log.Println(err)
	}
}

func ClearSlice(slice []interface{}) {
	if len(slice) > 40 {
		slice = nil
	}

}
