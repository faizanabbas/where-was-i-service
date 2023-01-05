package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type configuration struct {
	port int
	env  string
}

func main() {
	var config configuration

	flag.IntVar(&config.port, "port", 3000, "API server port")
	flag.StringVar(&config.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "status: available")
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", config.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
