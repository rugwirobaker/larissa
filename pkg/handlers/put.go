package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/rugwirobaker/larissa/pkg/larissa"
	"github.com/rugwirobaker/larissa/pkg/log"

	"github.com/gorilla/mux"
)

//PathPut is the http Upload path
const PathPut = "/put/{bucket}"

const maxUploadSize = 50 * 1024 * 1024 // 50 mb

// Put handles file http file uploads
func Put(proctl larissa.Protocol, lggr log.Entry) http.Handler {
	const op errors.Op = "handlers.Put"

	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bucket := vars["bucket"]

		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			lggr.SystemErr(errors.E(op, err))
			w.WriteHeader(errors.KindBadRequest)
			return
		}

		file, header, err := r.FormFile("image")
		if err != nil {
			lggr.SystemErr(errors.E(op, err))
			w.WriteHeader(errors.KindBadRequest)
			return
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			lggr.SystemErr(errors.E(op, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// get content type
		_, err = ext(content)
		if err != nil {
			lggr.SystemErr(errors.E(op, err))
			w.WriteHeader(errors.KindBadRequest)
			return
		}

		var filename = header.Filename

		if err := proctl.Put(filename, bucket, content); err != nil {
			severityLevel := errors.Expect(err, errors.KindNotFound)
			err = errors.E(op, err, severityLevel)
			lggr.SystemErr(err)
			w.WriteHeader(errors.Kind(err))
			return
		}

		message:=map[string]string{"message":"file created"}

		w.WriteHeader(http.StatusCreated)
		encodeRes(w, message)
	}
	return http.HandlerFunc(f)
}
