package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/rugwirobaker/larissa/pkg/larissa"
	"github.com/rugwirobaker/larissa/pkg/log"
)

// PathList is the http path to download files
const PathList = "/list/{bucket}"

// List ...
func List(proctl larissa.Protocol, lggr log.Entry) http.Handler {
	const op errors.Op = "handlers.List"

	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bucket := vars["bucket"]

		objects, err := proctl.List(r.Context(), bucket)
		if err != nil {
			severityLevel := errors.Expect(err, errors.KindNotFound)
			err = errors.E(op, err, severityLevel)
			lggr.SystemErr(err)
			w.WriteHeader(errors.Kind(err))
			return
		}
		encodeRes(w, objects)
	}

	return http.HandlerFunc(f)
}
