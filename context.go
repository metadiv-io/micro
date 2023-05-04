package micro

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/err_map"
	"github.com/metadiv-io/sql"
)

type Context[T any] struct {
	GinContext *gin.Context
	TraceID    string
	Page       *sql.Pagination
	Sort       *sql.Sort
	Request    *T
	Response   interface{}
}

func (ctx *Context[T]) ClientIP() string {
	return ctx.GinContext.ClientIP()
}

func (ctx *Context[T]) UserAgent() string {
	return ctx.GinContext.Request.UserAgent()
}

func (ctx *Context[T]) OK(data interface{}, traceID string, traces []Trace, page ...*sql.Pagination) {
	var p *sql.Pagination
	if len(page) > 0 {
		p = page[0]
	}
	var credit float64
	var duration uint
	for _, t := range traces {
		credit += t.Credit
		duration += t.Duration
	}
	resp := &Response{
		Success:    true,
		Data:       data,
		Duration:   duration,
		Credit:     credit,
		Pagination: p,
		TraceID:    traceID,
		Traces:     traces,
	}
	ctx.Response = resp // for testing
	ctx.GinContext.JSON(200, resp)
}

func (ctx *Context[T]) Error(err err_map.Error, traceID string, traces []Trace) {
	var credit float64
	var duration uint
	for _, t := range traces {
		credit += t.Credit
		duration += t.Duration
	}
	resp := &Response{
		Success:  false,
		TraceID:  traceID,
		Duration: duration,
		Credit:   credit,
		Traces:   traces,
		Error: &Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
	}
	ctx.Response = resp // for testing
	ctx.GinContext.JSON(200, resp)
}
