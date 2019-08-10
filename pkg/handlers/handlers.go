package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rugwirobaker/larissa/pkg/larissa"

	"github.com/rugwirobaker/larissa/pkg/build"

	log "github.com/sirupsen/logrus"
)

const maxUploadSize = 50 * 1024 * 1024 // 50 mb

// HTTPHandler describes an HTTP API to larissa.Service
type HTTPHandler struct {
	svc larissa.Service
}

// GetObjectRes describes larissa object response
type GetObjectRes struct{}

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

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "{\"message\": \"could not create file: "+err.Error()+"\"}", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "{\"message\": \"could not save: "+err.Error()+"\"}", 400)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "{\"message\": \"could save file: "+err.Error()+"\"}", 400)
		return
	}

	// get content type
	_, err = ext(content)
	if err != nil {
		http.Error(w, "{\"message\": \"could save file: "+err.Error()+"\"}", 400)
		return
	}

	var filename = header.Filename

	if err := handler.svc.Put(filename, bucket, content); err != nil {
		http.Error(w, "{\"message\": \"could not create file: "+err.Error()+"\"}", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "{\"message\": \"new file: %s/%s created\"}", bucket, filename)
}

// Get ...
func (handler HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	name := vars["name"]

	object, err := handler.svc.Get(name, bucket)
	if err != nil {
		http.Error(w, "{\"message\": \"could not get file: "+err.Error()+"\"}", http.StatusNotFound)
		return
	}

	contentType := mime(object.Content)
	encodeData(w, contentType, object.Content)
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

func encodeData(w io.Writer, contentType string, data []byte) {
	if headered, ok := w.(http.ResponseWriter); ok {
		headered.Header().Set("Cache-Control", "no-cache")
		headered.Header().Set("Content-Type", contentType)
	}
	w.Write(data)
}
