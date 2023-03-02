package handlers

import (
	"github.com/gin-gonic/gin"
	httpHandlers "main/internal/infrastructure/ui/http"
)

func Routes(route *gin.RouterGroup) {
	route.GET("/my-url", httpHandlers.MyHandler)
}
