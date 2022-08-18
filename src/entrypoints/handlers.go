package entrypoints

import "github.com/gin-gonic/gin"

type Handlers interface {
	Handle(c *gin.Context)
}
