package middleware

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ValidateRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: true}
		doc, _ := loader.LoadFromFile("api/docs/openapi.yml")

		// Validate document
		_ = doc.Validate(ctx)

		// Instantiate router
		router, _ := gorillamux.NewRouter(doc)

		// Find route
		route, pathParams, _ := router.FindRoute(ctx.Request)

		// Validate incoming request against the found OpenAPI route
		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    ctx.Request,
			PathParams: pathParams,
			Route:      route,
		}
		err := openapi3filter.ValidateRequest(ctx, requestValidationInput)

		if nil != err {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"message": http.StatusText(http.StatusUnprocessableEntity),
				"details": gin.H{
					"error": getErrorMessage(err),
				},
			})
		}
	}
}

func getErrorMessage(err error) string {
	errorSlices := strings.Split(err.Error(), "\n")
	errorMessage := strings.ReplaceAll(errorSlices[0], "\"", "'")

	return errorMessage
}
