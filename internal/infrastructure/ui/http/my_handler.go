package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetArticlesBoostedByTags(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}
