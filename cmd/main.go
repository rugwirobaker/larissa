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

	"github.com/rugwirobaker/larissa/cmd/app"

	"github.com/rugwirobaker/larissa/pkg/config"

	log "github.com/sirupsen/logrus"

	"github.com/rugwirobaker/larissa/pkg/build"
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

	handler, err := app.Bootstrap(conf)
	if err != nil {
		log.Fatalf("could not load config file: %v", err)
	}

	// start the server
	server := &http.Server{Addr: ":" + conf.Port, Handler: handler}

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
