package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/conf"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/router"
	"github.com/oresdev/tbcc-wallet-api-v3/internal/server/store"
)

const appName = "tbcc-wallet-api-v3"

func main() {
	conf, err := conf.ParseConfig(appName)
	if err != nil {
		logrus.Fatalf("parsing config; %v", err)
	}

	conn, err := store.СreateDB(conf)
	if err != nil {
		logrus.Errorf("opening connection: %v", err)
	}
	logrus.RegisterExitHandler(func() {
		conn.Close()
	})

	handler, err := router.CreateHTTPHandler(conn)
	if err != nil {
		logrus.Fatalf("creating http handler: %v", err)
	}

	listenErr := make(chan error, 1)
	server := &http.Server{
		Addr:    ":" + os.Args[1],
		Handler: handler,
	}

	go func() {
		logrus.Println("server started at port:", os.Args[1], time.Now().Format(time.RFC3339))
		listenErr <- server.ListenAndServe()
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-listenErr:
		logrus.Fatal(err)
		logrus.Exit(1)
	case <-osSignals:
		server.SetKeepAlivesEnabled(false)
		timeout := time.Second * 20

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			logrus.Fatal(err)
		}
		logrus.Println("stop server")
		logrus.Exit(0)
	}
}
