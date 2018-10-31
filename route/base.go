package route

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/LoadBalance", loadbalanceController)
	return router
}
