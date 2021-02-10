package cmd

import (
	"fmt"
	logg "log"
	"net/http"
	"os"
	"searchservice/config"
	"searchservice/endpoints"
	"searchservice/service"
	"searchservice/transport"
	"time"

	"github.com/go-kit/kit/log"
)

func startServer() error {
	// Defining Logger
	fmt.Println("Server starting...")
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	service := service.NewService(logger)
	endpoint := endpoints.MakeEndpoints(service)
	handler := transport.MakeHTTPHandlers(endpoint, logger)

	l := logg.New(os.Stdout, "SEARCH-SERVICE", 0)
	// Set handler and listen given port
	s := http.Server{
		Addr:         config.C.Server.Address, // configure the bind address
		Handler:      handler,                 // set the default handler
		ErrorLog:     l,                       // set the logger for the server
		ReadTimeout:  5 * time.Second,         // max time to read request from the client
		WriteTimeout: 10 * time.Second,        // max time to write response to the client
		IdleTimeout:  120 * time.Second,       // max time for connections using TCP Keep-Alive
	}

	return s.ListenAndServe()
}
