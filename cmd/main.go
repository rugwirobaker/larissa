package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/rugwirobaker/larissa/pkg/larissa"
)

const port = "9090"

func main() {
	//get configuration

	// set the storage backend
	storage := larissa.NewBackend()

	// create service
	service := larissa.New(storage)

	// create HTTPHandler
	httpHandler := larissa.NewHTTPHandler(service)

	// create a router
	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(larissa.ErrorHandler)
	router.Handle("/", httpHandler)
	router.HandleFunc("/build", httpHandler.Build).Methods("GET")
	router.HandleFunc("/get", httpHandler.Get).Methods("GET")
	router.HandleFunc("/put", httpHandler.Put).Methods("PUT")
	router.HandleFunc("/del", httpHandler.Del).Methods("DELETE")
	router.HandleFunc("/exists", httpHandler.Exists).Methods("GET")

	// start the server
	server := &http.Server{Addr: ":" + port, Handler: router}

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
