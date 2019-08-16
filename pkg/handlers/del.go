package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/rugwirobaker/larissa/pkg/larissa"
	"github.com/rugwirobaker/larissa/pkg/log"
)

// PathDel is the http path to download files
const PathDel = "/del/{bucket}/{name}"

// Del ...
func Del(proctl larissa.Protocol, lggr log.Entry) http.Handler {
	const op errors.Op = "handlers.Del"
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		name := vars["name"]

		if err := proctl.Del(r.Context(), name, bucket); err != nil {
			severityLevel := errors.Expect(err, errors.KindNotFound)
			err = errors.E(op, err, severityLevel)
			lggr.SystemErr(err)
			w.WriteHeader(errors.Kind(err))
			return
		}
	}
	return http.HandlerFunc(f)
}
