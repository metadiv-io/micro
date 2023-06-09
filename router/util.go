package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/sql"
)

func newGinServiceHandler[T any](engine *micro.Engine, handler micro.Handler[T]) gin.HandlerFunc {
	handlerSetup := handler()
	return func(c *gin.Context) {
		now := time.Now()
		var credit float64
		api, ok := micro.API_MAP[c.Request.Method+":"+c.FullPath()]
		if ok {
			credit = api.Credit
		}
		traceID := micro.GetTraceID(c)
		micro.SetTraceID(c, traceID)
		ctx := &micro.Context[T]{
			GinContext: c,
			Request:    micro.GinRequest[T](c),
			TraceID:    traceID,
		}
		if handlerSetup.Pagination {
			ctx.Page = micro.GinRequest[sql.Pagination](c)
		}
		if handlerSetup.Sort {
			ctx.Sort = micro.GinRequest[sql.Sort](c)
		}
		resp, err := handlerSetup.Service(ctx)
		traces := micro.GetTraces(c)
		overrideCredit := ctx.GinContext.GetFloat64("credit")
		if overrideCredit > 0 {
			credit = overrideCredit
		}
		if err != nil {
			traces = append(traces, micro.Trace{
				Success:    false,
				Time:       time.Now().Format("2006-01-02 15:04:05"),
				SystemUUID: micro.SYSTEM_UUID,
				SystemName: micro.SYSTEM_NAME,
				TraceID:    traceID,
				Duration:   uint(time.Since(now).Microseconds()),
				Credit:     credit,
				Error: &micro.Error{
					Code:    err.Code(),
					Message: err.Error(),
				},
			})
			ctx.Error(err, traceID, traces)
			return
		}
		traces = append(traces, micro.Trace{
			Success:    true,
			Time:       time.Now().Format("2006-01-02 15:04:05"),
			TraceID:    traceID,
			Duration:   uint(time.Since(now).Microseconds()),
			Credit:     credit,
			SystemUUID: micro.SYSTEM_UUID,
			SystemName: micro.SYSTEM_NAME,
		})
		ctx.OK(resp, traceID, traces, ctx.Page)
	}
}

var wsUpGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func newGinWSServiceHandler[T any](engine *micro.Engine, handler micro.WSHandler[T]) gin.HandlerFunc {
	handlerSetup := handler()
	return func(c *gin.Context) {
		ws, err := wsUpGrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		defer ws.Close()
		ctx := &micro.Context[T]{
			GinContext: c,
			Request:    micro.GinRequest[T](c),
		}
		err1 := handlerSetup.Service(ctx, ws)
		if err != nil {
			ctx.Error(err1, "", nil)
		}
	}
}

func joinMiddlewareAndService(service gin.HandlerFunc, middleware ...gin.HandlerFunc) []gin.HandlerFunc {
	var funcs = make([]gin.HandlerFunc, 0)
	if len(middleware) > 0 {
		funcs = append(funcs, middleware...)
	}
	funcs = append(funcs, service)
	return funcs
}
