package middlewares

import (
	"boilerplate/usecase"
	"github.com/gin-gonic/gin"
)

type ErrorHandler struct {
	ErrorHandlerUsecase usecase.ErrorHandlerUsecase
}

type ResponseErrSnap struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

func NewErrorHandler(r *gin.RouterGroup, ehus usecase.ErrorHandlerUsecase) {
	handler := &ErrorHandler{
		ErrorHandlerUsecase: ehus,
	}

	r.Use(handler.errorHandler)
}

func (eh *ErrorHandler) errorHandler(c *gin.Context) {
	c.Next()

	errorToPrint := c.Errors.Last()
	if errorToPrint != nil {
		c.JSON(eh.ErrorHandlerUsecase.ResponseError(errorToPrint))
		c.Abort()
	}
	return
}
