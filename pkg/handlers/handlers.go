package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rugwirobaker/larissa/pkg/larissa"

	"github.com/rugwirobaker/larissa/pkg/build"

	"github.com/rugwirobaker/larissa/pkg/log"
)

//ProtocolHandler returns an http.Handler
type ProtocolHandler func(proctl larissa.Protocol, lggr log.Entry) http.Handler

// HandlerOpts are the generic options
// for a ProtocolHandler
type HandlerOpts struct {
	Protocol larissa.Protocol
	Logger   *log.Logger
}

// LogEntryHandler pulls a log entry from the request context.
func LogEntryHandler(h ProtocolHandler, opts *HandlerOpts) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		ent := log.EntryFromContext(r.Context())
		handler := h(opts.Protocol, ent)
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(f)
}

// RegisterHandlers is a convenience method that registers
func RegisterHandlers(r *mux.Router, opts *HandlerOpts) {
	// If true, this would only panic at boot time, static nil checks anyone?
	if opts == nil || opts.Protocol == nil || opts.Logger == nil {
		panic("absolutely unacceptable handler opts")
	}

	r.Handle(PathPut, LogEntryHandler(Put, opts)).Methods(http.MethodPut)
	r.Handle(PathGet, LogEntryHandler(Get, opts)).Methods(http.MethodGet)
	r.Handle(PathDel, LogEntryHandler(Del, opts)).Methods(http.MethodDelete)
	r.Handle(PathList, LogEntryHandler(List, opts)).Methods(http.MethodGet)
	r.Handle(PathExists, LogEntryHandler(Exists, opts)).Methods(http.MethodGet)
}

// Build returns larissa build information
func Build(w http.ResponseWriter, r *http.Request) {
	encodeRes(w, build.Data())
}

// Health indicates the health of the server
func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
