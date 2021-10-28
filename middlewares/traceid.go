package middlewares

import (
	"boilerplate/models/contract"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type TraceIDHandler struct {
}

func NewTraceIDHandler(r *gin.RouterGroup) {
	handler := &TraceIDHandler{}

	r.Use(handler.traceIDHandler)
}

func (eh *TraceIDHandler) traceIDHandler(c *gin.Context) {

	parentid := c.GetHeader(contract.HeaderReqID)
	if parentid != ""{
		c.Request.Header.Set(contract.HeaderParentReqID, parentid)
		c.Writer.Header().Set(contract.HeaderParentReqID,parentid)
	}

	id := uuid.NewV4().String()
	c.Request.Header.Set(contract.HeaderReqID, id)
	c.Writer.Header().Set(contract.HeaderReqID,id)
	c.Next()
}