package serverfx

import (
	"context"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func lifecycle(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":8080"}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping Server")
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

var ModuleServer = fx.Module("server", fx.Invoke(lifecycle))
