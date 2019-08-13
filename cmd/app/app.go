package app

import (
	"fmt"
	"net/http"

	"github.com/rugwirobaker/larissa/pkg/handlers"

	"github.com/rugwirobaker/larissa/pkg/larissa"
	"github.com/rugwirobaker/larissa/pkg/log"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/larissa/pkg/config"
	mw "github.com/rugwirobaker/larissa/pkg/middleware"
)

// Bootstrap is where all routes and middleware for the server
func Bootstrap(conf *config.Config) (http.Handler, error) {

	// get storage backend
	storage, err := GetStorage(conf.StorageType, conf.Storage)
	if err != nil {
		err = fmt.Errorf("error getting storage configuration (%s)", err)
		return nil, err
	}

	logLvl, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		return nil, err
	}
	lggr := log.New(conf.CloudRuntime, logLvl)

	// create service
	protocol := larissa.New(storage)

	//new router
	r := mux.NewRouter()
	r.Use(mw.LogEntryMiddleware(lggr))

	if conf.GoEnv == "development" {
		r.Use(mw.RequestLogger)
	}

	//r.NotFoundHandler = http.HandlerFunc(handlers.NotFoundHandler)
	r.HandleFunc("/build", handlers.Build).Methods("GET")
	r.HandleFunc("/health", handlers.Health).Methods("GET")

	handlerOpts := &handlers.HandlerOpts{Protocol: protocol, Logger: lggr}
	handlers.RegisterHandlers(r, handlerOpts)

	return r, nil
}
