package bootstrap

import (
	"go-fwgin/internal/config"
	"go-fwgin/internal/modules/user"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config      *config.Config
	UserHandler *user.HandlerUser
}

func (a *App) Start() error {
	if a.Config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// gin engine
	r := gin.Default()
	apiV1 := r.Group("/api/v1")
	cfg := a.Config
	a.UserHandler.RoutesUser(apiV1, cfg)

	// run
	log.Printf("Application [%s] is running on port :%s in %s mode", a.Config.AppName, a.Config.AppPort, a.Config.AppEnv)
	return r.Run(a.Config.AppPort)
}
