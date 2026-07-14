package container

import "github.com/gin-gonic/gin"

type Module interface {
	RegisterRoutes(public *gin.RouterGroup, protected *gin.RouterGroup)
}
