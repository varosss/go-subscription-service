package builder

import (
	"fmt"
	httpcontroller "go-subscription-service/internal/infrastructure/adapter/controller/http"
	gormrepo "go-subscription-service/internal/infrastructure/adapter/gorm/repo"
	"go-subscription-service/internal/infrastructure/config"
	"go-subscription-service/internal/infrastructure/db"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
)

type App struct {
	httpServer *http.Server
	closers    []func() error
}

func (a *App) Run() error {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	errCh := make(chan error, 1)
	go func() {
		fmt.Printf("HTTP server running on %s\n", a.httpServer.Addr)
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case sig := <-stop:
		fmt.Printf("Received signal %s, shutting down...\n", sig)
		return a.Shutdown()
	case err := <-errCh:
		return err
	}
}

func (a *App) Shutdown() error {
	for _, closer := range a.closers {
		if err := closer(); err != nil {
			return err
		}
	}
	return nil
}

func BuildApp(cfg *config.Config) (*App, error) {
	app := &App{}
	var closers []func() error

	conn, err := db.ConnectGorm(cfg.Database.DSN)
	if err != nil {
		return nil, err
	}
	closers = append(closers, func() error {
		return db.CloseGorm(conn)
	})

	router := gin.Default()
	router.Use(gin.Recovery())

	uc := BuildUseCases(gormrepo.NewGormSubscriptionRepo(conn))

	subsHandler := httpcontroller.NewSubscriptionHandler(
		uc.CreateSub,
		uc.GetSub,
		uc.UpdateSub,
		uc.DeleteSub,
		uc.ListSubs,
	)
	subsHandler.RegisterRoutes(router)

	app.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.HttpServer.Port),
		Handler: router,
	}
	app.closers = closers

	return app, nil
}
