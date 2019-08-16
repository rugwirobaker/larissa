package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/rugwirobaker/larissa/pkg/larissa"
	"github.com/rugwirobaker/larissa/pkg/log"
)

// PathGet is the http path to download files
const PathGet = "/get/{bucket}/{name}"

// Get ...
func Get(proctl larissa.Protocol, lggr log.Entry) http.Handler {
	const op errors.Op = "handlers.Get"

	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		name := vars["name"]

		object, err := proctl.Get(r.Context(), name, bucket)
		if err != nil {
			severityLevel := errors.Expect(err, errors.KindNotFound)
			err = errors.E(op, err, severityLevel)
			lggr.SystemErr(err)
			w.WriteHeader(errors.Kind(err))
			return
		}

		contentType := mime(object.Content)
		encodeData(w, contentType, object.Content)
	}
	return http.HandlerFunc(f)
}
