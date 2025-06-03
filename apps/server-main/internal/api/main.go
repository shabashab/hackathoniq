package api

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/shabashab/hackathoniq/apps/server-main/internal/api/controllers"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/api/middleware"
	"github.com/shabashab/hackathoniq/apps/server-main/internal/config/configs"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("api",
	controllers.Module,
	middleware.Module,

	fx.Provide(
		newServer,
	),
	fx.Invoke(setupRoutes),
)

func newServer(lc fx.Lifecycle, apiConfig *configs.ApiConfig) (*http.Server, *gin.Engine) {
	router := gin.Default()

	server := &http.Server{
		Addr: apiConfig.Host,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Handler = router.Handler()

			listener, err := net.Listen("tcp", server.Addr)

			if err != nil {
				return err
			}

			fmt.Println("Starting http server at", server.Addr)

			go server.Serve(listener)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})

	return server, router
}
