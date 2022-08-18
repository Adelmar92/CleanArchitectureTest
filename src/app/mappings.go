package app

import (
	"github.com/adelmar92/GoApiClean/src/infraestructure/dependencies"
	"github.com/gin-gonic/gin"
)

const (
	pingURL = "/ping"
)

func ConfigureMappings(router *gin.Engine, handlers dependencies.HandlerContainer) {
	router.GET(pingURL, handlers.Ping.Handle)
}
