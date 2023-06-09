package micro

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/sql"
	"github.com/robfig/cron"
)

type Engine struct {
	GinEngine  *gin.Engine
	CronWorker *cron.Cron
}

func NewEngine() *Engine {
	e := &Engine{
		GinEngine:  gin.Default(),
		CronWorker: cron.New(),
	}
	GET(e, "/ping", PingHandler, ginmid.RateLimited(time.Minute, 30))
	return e
}

func (e *Engine) Run(addr string) {
	e.CronWorker.Start()
	e.GinEngine.Run(addr)
}

func GET[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), middleware...)...)
}

func GETWithCache[T any](engine *Engine, route string, cacheDuration time.Duration, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.GET(route, joinMiddlewareAndService(
		ginmid.Cache(cacheDuration, newGinServiceHandler(engine, handler)), middleware...)...)
}

func POST[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.POST(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), middleware...)...)
}

func PUT[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.PUT(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), middleware...)...)
}

func DELETE[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.DELETE(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), middleware...)...)
}

func WS[T any](engine *Engine, route string, handler WSHandler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinWSServiceHandler(engine, handler), middleware...)...)
}

func Cron(engine *Engine, spec string, job func()) {
	engine.CronWorker.AddFunc(spec, job)
}

func newGinServiceHandler[T any](engine *Engine, handler Handler[T]) gin.HandlerFunc {
	handlerSetup := handler()
	return func(c *gin.Context) {
		now := time.Now()
		var credit float64
		api, ok := API_MAP[c.Request.Method+":"+c.FullPath()]
		if ok {
			credit = api.Credit
		}
		traceID := GetTraceID(c)
		SetTraceID(c, traceID)
		ctx := &Context[T]{
			GinContext: c,
			Request:    GinRequest[T](c),
			TraceID:    traceID,
		}
		if handlerSetup.Pagination {
			ctx.Page = GinRequest[sql.Pagination](c)
		}
		if handlerSetup.Sort {
			ctx.Sort = GinRequest[sql.Sort](c)
		}
		resp, err := handlerSetup.Service(ctx)
		traces := GetTraces(c)
		overrideCredit := ctx.GinContext.GetFloat64("credit")
		if overrideCredit > 0 {
			credit = overrideCredit
		}
		if err != nil {
			traces = append(traces, Trace{
				Success:    false,
				Time:       time.Now().Format("2006-01-02 15:04:05"),
				SystemUUID: SYSTEM_UUID,
				SystemName: SYSTEM_NAME,
				TraceID:    traceID,
				Duration:   uint(time.Since(now).Microseconds()),
				Credit:     credit,
				Error: &Error{
					Code:    err.Code(),
					Message: err.Error(),
				},
			})
			ctx.Error(err, traceID, traces)
			return
		}
		traces = append(traces, Trace{
			Success:    true,
			Time:       time.Now().Format("2006-01-02 15:04:05"),
			TraceID:    traceID,
			Duration:   uint(time.Since(now).Microseconds()),
			Credit:     credit,
			SystemUUID: SYSTEM_UUID,
			SystemName: SYSTEM_NAME,
		})
		ctx.OK(resp, traceID, traces, ctx.Page)
	}
}

var wsUpGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func newGinWSServiceHandler[T any](engine *Engine, handler WSHandler[T]) gin.HandlerFunc {
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
		ctx := &Context[T]{
			GinContext: c,
			Request:    GinRequest[T](c),
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
