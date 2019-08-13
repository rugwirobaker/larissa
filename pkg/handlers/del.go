package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/larissa/pkg/errors"
	"github.com/rugwirobaker/larissa/pkg/larissa"
	"github.com/rugwirobaker/larissa/pkg/log"
)

// PathDel is the http path to download files
const PathDel = "/del/{bucket}/{name}"

// Del ...
func Del(proto larissa.Service, lggr log.Entry) http.Handler {
	const op errors.Op = "handlers.Del"
	f := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bucket := vars["bucket"]
		name := vars["name"]

		encodeRes(w, struct {
			Message string `json:"message"`
		}{fmt.Sprintf("delete `%s.png` from `%s` bucket", name, bucket)},
		)
	}
	return http.HandlerFunc(f)
}
