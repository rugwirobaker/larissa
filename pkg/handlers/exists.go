package handlers

import (
	"fmt"
	"net/http"

	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/rugwirobaker/larissa/pkg/larissa"
	"github.com/rugwirobaker/larissa/pkg/log"

	"github.com/gorilla/mux"
)

// PathExists is the http path to download files
const PathExists = "/exists/{bucket}/{name}"

// Exists ...
func Exists(proto larissa.Service, lggr log.Entry) http.Handler {
	const op errors.Op = "handlers.Exists"
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		name := vars["name"]

		encodeRes(w, struct {
			Message string `json:"message"`
		}{fmt.Sprintf("find `%s.png` from `%s` bucket", name, bucket)},
		)
	}
	return http.HandlerFunc(f)
}
