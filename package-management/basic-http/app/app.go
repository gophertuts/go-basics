package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"

	"github.com/gophertuts/go-basics/package-management/basic-http/config"
	"github.com/gophertuts/go-basics/package-management/basic-http/controllers"
	"github.com/gophertuts/go-basics/package-management/basic-http/logging"
)

type App struct {
	Config config.Manager
	Router *httprouter.Router
	Server http.Server
}

func New() *App {
	err := logging.InitLogger()
	if err != nil {
		log.Fatal(err)
	}
	manager, err := config.New()
	if err != nil {
		logging.Logger.Error("could not initialize config")
	}
	router := controllers.New()
	return &App{
		Config: manager,
		Router: router,
	}
}

func (app *App) Start() error {
	logging.Logger.Info(
		"application started",
		zap.String("listen", app.Config.Listen()),
	)
	server := http.Server{
		Addr: app.Config.Listen(),
		Handler: app.Router,
	}
	app.Server = server
	err := server.ListenAndServe()
	if err == http.ErrServerClosed {
		logging.Logger.Info("http server is closed")
		return nil
	}
	return err
}

func (app *App) Shutdown() {
	// Need to call Sync() before exiting. See Sync docs
	defer logging.Logger.Sync()
	ch := make(chan struct{})
	go func() {
		logging.Logger.Info("shutting down http server")
		if err := app.Server.Shutdown(context.Background()); err != nil {
			logging.Logger.Error("error on server shutdown", zap.Error(err))
		}
		close(ch)
	}()
	select {
	case <-ch:
		logging.Logger.Info("application was shut down")
	case <-time.After(2 * time.Second):
		logging.Logger.Error("could not shut down in", zap.Duration("timeout", 2 * time.Second))
	}
}

type Shutdowner interface {
	Shutdown()
}

func ListenForSignals(signals []os.Signal, servers ...Shutdowner) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	sig := <-c
	logging.Logger.Info("received shutdown signal", zap.String("signal", sig.String()))

	for _, server := range servers {
		server.Shutdown()
	}
	os.Exit(0)
}
