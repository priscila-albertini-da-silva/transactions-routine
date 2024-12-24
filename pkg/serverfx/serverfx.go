package serverfx

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

type Route struct {
	Path    string
	Method  string
	Handler gin.HandlerFunc
}

type params struct {
	fx.In
	Rt []Route
}

func StartServer(lifecycle fx.Lifecycle, p params) {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api := router.Group("/")

	for _, r := range p.Rt {
		handlers := []gin.HandlerFunc{r.Handler}
		api.Handle(r.Method, r.Path, handlers...)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info("Starting server on :8080")
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Error("Server error: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping Server...")
			return server.Close()
		},
	})
}

var ModuleServer = fx.Options(
	fx.Invoke(StartServer),
)
