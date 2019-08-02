package larissa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HTTPHandler describes an HTTP API to larissa.Service
type HTTPHandler struct {
	svc Service
}

// NewHTTPHandler creates a new instance of HTTPHandler
func NewHTTPHandler(svc Service) HTTPHandler {
	return HTTPHandler{svc}
}

func (handler HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "{\"message\": \"Welcome to Larissa Image Server\"}")
}

// Build returns larissa build information
func (handler HTTPHandler) Build(w http.ResponseWriter, r *http.Request) {
	encodeRes(w, Data())
}

// Put ..
func (handler HTTPHandler) Put(w http.ResponseWriter, r *http.Request) {
	encodeRes(w, struct {
		Message string `json:"message"`
	}{"you have reached the put handler"},
	)
}

// Get ...
func (handler HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	encodeRes(w, struct {
		Message string `json:"message"`
	}{"you have reached the get handler"},
	)
}

// Del ...
func (handler HTTPHandler) Del(w http.ResponseWriter, r *http.Request) {
	encodeRes(w, struct {
		Message string `json:"message"`
	}{"you have reached the delete handler"},
	)
}

// Exists ...
func (handler HTTPHandler) Exists(w http.ResponseWriter, r *http.Request) {
	encodeRes(w, struct {
		Message string `json:"message"`
	}{"you have reached the exists handler"},
	)
}

// ErrorHandler handles non-existing route
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"handler not found for path: " + r.URL.Path + "\"}"))
}

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
