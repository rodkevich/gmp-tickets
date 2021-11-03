package msg

import (
	"log"
	"net/http"
)

// ReturnServerError 500
func ReturnServerError(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	if _, err := w.Write([]byte(`Internal Server Error`)); err != nil {
		log.Println(err)
	}
}

// ReturnClientError 400
func ReturnClientError(w http.ResponseWriter, text string) {
	log.Println(text)
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write([]byte(text)); err != nil {
		log.Println(err)
	}
}
