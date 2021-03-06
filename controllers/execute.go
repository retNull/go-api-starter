package controllers

import (
	"github.com/gin-gonic/gin"
)

// ExecuteControllerMethod is called on every route. It validates the request body, calls the controller method, and returns an http response
func ExecuteControllerMethod(bindingModel interface{}, method func(interface{}) ControllerResult) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := context.ShouldBindJSON(bindingModel); err != nil {
			context.AbortWithError(400, err)
			return
		}

		r := method(bindingModel)

		if !r.Success {
			if r.Error != nil {
				context.AbortWithError(r.Code, r.Error)
			} else if r.ErrorMessage != "" {
				context.AbortWithStatusJSON(r.Code, gin.H{
					"error": r.ErrorMessage,
				})
			} else {
				context.AbortWithStatusJSON(r.Code, gin.H{
					"error": "Unknown error",
				})
			}
		} else {
			context.JSON(r.Code, gin.H{
				"data": r.Data,
			})
		}
	}
}
