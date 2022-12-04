package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/hello", func(writer http.ResponseWriter, request *http.Request) {
		log.Info("/api/v1/hello invoke")
		writer.Write([]byte("hello, I am mux server"))
		writer.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			writer.Header().Set(k, v[0])
		}
		version := "default version"
		writer.Header().Set("VERSION", version)
		writer.WriteHeader(http.StatusOK)
	})

	server := http.Server{
		Addr:              ":8888",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Errorf("err: %v\n", err)
	}
}
