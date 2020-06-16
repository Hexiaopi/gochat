package api

import (
	"context"
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var port string

func init() {
	flag.StringVar(&port, "http.port", ":8080", "set http port")
}

var server *http.Server

func StartServer() {
	router := RegisterRouter()

	server = &http.Server{
		Addr:    port,
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Errorf("start http server fail:%v", err)
		panic(err)
	}
}

func StopServer(ctx context.Context) {
	if err := server.Shutdown(ctx); err != nil {
		log.Errorf("http server shut down server fail %v", err)
	} else {
		log.Info("http server shut down success.")
	}
}
