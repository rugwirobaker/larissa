package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rugwirobaker/larissa/pkg/config"

	"github.com/rugwirobaker/larissa/pkg/handlers"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/larissa/pkg/build"
	"github.com/rugwirobaker/larissa/pkg/larissa"
)

var (
	configFile = flag.String("config_file", "", "The path to the config file")
	version    = flag.Bool("version", false, "Print version information and exit")
)

func main() {
	flag.Parse()
	if *version {
		fmt.Println(build.String())
		os.Exit(0)
	}
	//get configuration
	conf, err := config.Load(*configFile)
	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	// get storage backend
	storage, err := GetStorage(conf.StorageType, conf.Storage)
	if err != nil {
		log.Fatalf("larissa failed to start: %s", err.Error())
	}

	// create service
	service := larissa.New(storage)

	//logger
	logger := log.New()

	// create HTTPHandler
	httpHandler := handlers.NewHTTPHandler(service, logger)

	// create a router
	router := mux.NewRouter()
	router.Use(httpHandler.LogRequest)
	router.NotFoundHandler = http.HandlerFunc(handlers.ErrorHandler)
	router.Handle("/", httpHandler)
	router.HandleFunc("/build", httpHandler.Build).Methods("GET")
	router.HandleFunc("/health", httpHandler.Health).Methods("GET")
	router.HandleFunc("/get/{bucket}/{name}", httpHandler.Get).Methods("GET")
	router.HandleFunc("/put/{bucket}", httpHandler.Put).Methods("PUT")
	router.HandleFunc("/del/{bucket}/{name}", httpHandler.Del).Methods("DELETE")
	router.HandleFunc("/exists/{bucket}/{name}", httpHandler.Exists).Methods("GET")

	// start the server
	server := &http.Server{Addr: ":" + conf.Port, Handler: router}

	go func() {
		if err := server.ListenAndServe(); err == nil {
			log.Error(err)
		}
	}()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)

		done <- true
	}()

	log.Info("awaiting signal...")
	<-done

	if err := server.Shutdown(ctx); err != nil {
		log.Error(err)
	}
	log.Info("exiting...")
}
