package handlers

import (
	"net/http"

	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/rugwirobaker/larissa/pkg/larissa"
	"github.com/rugwirobaker/larissa/pkg/log"

	"github.com/gorilla/mux"
)

// PathExists is the http path to download files
const PathExists = "/exists/{bucket}/{name}"

// Exists ...
func Exists(proctl larissa.Protocol, lggr log.Entry) http.Handler {
	const op errors.Op = "handlers.Exists"
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		name := vars["name"]

		if err := proctl.Exists(name, bucket); err != nil {
			severityLevel := errors.Expect(err, errors.KindNotFound)
			err = errors.E(op, err, severityLevel)
			lggr.SystemErr(err)
			w.WriteHeader(errors.Kind(err))
			return
		}
		w.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(f)
}
