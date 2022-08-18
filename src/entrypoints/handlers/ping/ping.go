package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Ping struct {
}

func NewPing() Ping {
	return Ping{}
}

func (p Ping) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
