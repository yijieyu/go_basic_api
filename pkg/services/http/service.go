package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yijieyu/go_basic_api/pkg/configuration"
)

var httpServer *http.Server

func Run(conf configuration.HTTP, handler http.Handler) {
	address := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	log.Printf("listening and serving HTTP on %s\n", address)
	logrus.WithFields(logrus.Fields{
		"address": address,
		"state":   "running",
	}).Info("service is starting...")

	httpServer = &http.Server{
		Addr:        address,
		Handler:     handler,
		IdleTimeout: time.Minute,
	}

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("fail to start server: %v", err)
	}
}

func RegisterOnShutdown(f func()) {
	httpServer.RegisterOnShutdown(f)
}

func OnServiceStop(timeout time.Duration) {
	if timeout <= 0 {
		timeout = 10 * time.Minute
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	logrus.WithFields(logrus.Fields{
		"state":    "stop",
		"shutdown": httpServer.Shutdown(ctx),
	}).Info("service is stopping...")
}

func Catch(then func(), signals ...os.Signal) {
	c := make(chan os.Signal)
	signal.Notify(c, signals...)
	<-c
	if then != nil {
		then()
	}
}
