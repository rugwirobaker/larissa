package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func encodeRes(w io.Writer, i interface{}) {
	if headered, ok := w.(http.ResponseWriter); ok {
		headered.Header().Set("Cache-Control", "no-cache")
		headered.Header().Set("Content-type", "application/json")
	}

	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		log.WithFields(log.Fields{
			"route": "mustEncode",
		}).Errorf("error encoding response to json: %s", err.Error())
	}
}

func encodeData(w io.Writer, contentType string, data []byte) {
	if headered, ok := w.(http.ResponseWriter); ok {
		headered.Header().Set("Cache-Control", "no-cache")
		headered.Header().Set("Content-Type", contentType)
	}
	w.Write(data)
}

func encodeErr(w io.Writer, err error) {
	if headered, ok := w.(http.ResponseWriter); ok {
		headered.Header().Set("Cache-Control", "no-cache")
		headered.Header().Set("Content-Type", "application/json")
	}
}
