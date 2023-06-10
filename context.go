package micro

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/err_map"
	"github.com/metadiv-io/micro/header"
	"github.com/metadiv-io/micro/types"
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

func (ctx *Context[T]) GetTraceID() string {
	return header.GetTraceID(ctx.GinContext)
}

func (ctx *Context[T]) SetTraceID(traceID string) {
	header.SetTraceID(ctx.GinContext, traceID)
}

func (ctx *Context[T]) GetTraces() []types.Trace {
	return header.GetTraces(ctx.GinContext)
}

func (ctx *Context[T]) SetTraces(traces []types.Trace) {
	header.SetTraces(ctx.GinContext, traces)
}

func (ctx *Context[T]) GetWorkspace() string {
	return header.GetWorkspace(ctx.GinContext)
}

func (ctx *Context[T]) SetWorkspace(workspace string) {
	header.SetWorkspace(ctx.GinContext, workspace)
}

func (ctx *Context[T]) OK(data interface{}) {
	if ctx.Response != nil {
		panic("Response already set")
	}
	traceID := ctx.GetTraceID()
	traces := ctx.GetTraces()
	var credit float64
	var duration uint
	for _, t := range traces {
		credit += t.Credit
		duration += t.Duration
	}
	resp := &types.Response{
		Success:    true,
		Data:       data,
		Duration:   duration,
		Credit:     credit,
		Pagination: ctx.Page,
		TraceID:    traceID,
		Traces:     traces,
	}
	ctx.Response = resp
}

func (ctx *Context[T]) Error(errCode string) {
	if ctx.Response != nil {
		panic("Response already set")
	}
	traceID := ctx.GetTraceID()
	traces := ctx.GetTraces()
	err := err_map.NewError(errCode)
	var credit float64
	var duration uint
	for _, t := range traces {
		credit += t.Credit
		duration += t.Duration
	}
	resp := &types.Response{
		Success:  false,
		TraceID:  traceID,
		Duration: duration,
		Credit:   credit,
		Traces:   traces,
		Error: &types.ErrorImpl{
			Code:    err.Code(),
			Message: err.Error(),
		},
	}
	ctx.Response = resp
}
