package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rugwirobaker/larissa/pkg/larissa"

	"github.com/rugwirobaker/larissa/pkg/build"

	log "github.com/sirupsen/logrus"
)

// HTTPHandler describes an HTTP API to larissa.Service
type HTTPHandler struct {
	svc larissa.Service
}

// NewHTTPHandler creates a new instance of HTTPHandler
func NewHTTPHandler(svc larissa.Service) HTTPHandler {
	return HTTPHandler{svc}
}

func (handler HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "{\"message\": \"Welcome to Larissa Image Server\"}")
}

// Build returns larissa build information
func (handler HTTPHandler) Build(w http.ResponseWriter, r *http.Request) {
	encodeRes(w, build.Data())
}

// Put ..
func (handler HTTPHandler) Put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	name := vars["name"]

	encodeRes(w, struct {
		Message string `json:"message"`
	}{fmt.Sprintf("save `%s.png` to `%s` bucket", name, bucket)},
	)
}

// Get ...
func (handler HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	name := vars["name"]

	encodeRes(w, struct {
		Message string `json:"message"`
	}{fmt.Sprintf("get `%s.png` from `%s` bucket", name, bucket)},
	)
}

// Del ...
func (handler HTTPHandler) Del(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	name := vars["name"]

	encodeRes(w, struct {
		Message string `json:"message"`
	}{fmt.Sprintf("delete `%s.png` from `%s` bucket", name, bucket)},
	)
}

// Exists ...
func (handler HTTPHandler) Exists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	name := vars["name"]

	encodeRes(w, struct {
		Message string `json:"message"`
	}{fmt.Sprintf("find `%s.png` from `%s` bucket", name, bucket)},
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
