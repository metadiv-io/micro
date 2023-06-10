package micro

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/micro/auth"
	"github.com/metadiv-io/micro/constant"
	"github.com/metadiv-io/micro/ginhelp"
	"github.com/metadiv-io/micro/header"
	"github.com/metadiv-io/micro/system"
	"github.com/metadiv-io/micro/types"
	"github.com/metadiv-io/micro/usage"
	"github.com/metadiv-io/sql"
)

func GET[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), middleware...)...)
}

func CacheGET[T any](engine *Engine, route string, cacheDuration time.Duration, handler Handler[T], middleware ...gin.HandlerFunc) {
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

func AdminGET[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func AdminCacheGET[T any](engine *Engine, route string, cacheDuration time.Duration, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(
		ginmid.Cache(cacheDuration, newGinServiceHandler(engine, handler)), newMiddleware...)...)
}

func AdminPOST[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.POST(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func AdminPUT[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.PUT(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func AdminDELETE[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.DELETE(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func AdminWS[T any](engine *Engine, route string, handler WSHandler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinWSServiceHandler(engine, handler), newMiddleware...)...)
}

func UsageGET[T any](engine *Engine, route string, uuid string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	types.API_MAP["GET:"+route] = types.Api{
		Tag:  "GET:" + route,
		UUID: uuid,
	}
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UsageCacheGET[T any](engine *Engine, route string, uuid string, cacheDuration time.Duration, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	types.API_MAP["GET:"+route] = types.Api{
		Tag:  "GET:" + route,
		UUID: uuid,
	}
	engine.GinEngine.GET(route, joinMiddlewareAndService(
		ginmid.Cache(cacheDuration, newGinServiceHandler(engine, handler)), newMiddleware...)...)
}

func UsagePOST[T any](engine *Engine, route string, uuid string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	types.API_MAP["POST:"+route] = types.Api{
		Tag:  "POST:" + route,
		UUID: uuid,
	}
	engine.GinEngine.POST(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UsagePUT[T any](engine *Engine, route string, uuid string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	types.API_MAP["PUT:"+route] = types.Api{
		Tag:  "PUT:" + route,
		UUID: uuid,
	}
	engine.GinEngine.PUT(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UsageDELETE[T any](engine *Engine, route string, uuid string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	types.API_MAP["DELETE:"+route] = types.Api{
		Tag:  "DELETE:" + route,
		UUID: uuid,
	}
	engine.GinEngine.DELETE(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UsageWS[T any](engine *Engine, route string, uuid string, handler WSHandler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	types.API_MAP["WS:"+route] = types.Api{
		Tag:  "WS:" + route,
		UUID: uuid,
	}
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinWSServiceHandler(engine, handler), newMiddleware...)...)
}

func UserGET[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UserCacheGET[T any](engine *Engine, route string, cacheDuration time.Duration, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(
		ginmid.Cache(cacheDuration, newGinServiceHandler(engine, handler)), newMiddleware...)...)
}

func UserPOST[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.POST(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UserPUT[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.PUT(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UserDELETE[T any](engine *Engine, route string, handler Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.DELETE(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UserWS[T any](engine *Engine, route string, handler WSHandler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinWSServiceHandler(engine, handler), newMiddleware...)...)
}

func newGinServiceHandler[T any](engine *Engine, handler Handler[T]) gin.HandlerFunc {
	handlerSetup := handler()
	return func(c *gin.Context) {
		now := time.Now()
		var credit float64
		api, ok := types.API_MAP[c.Request.Method+":"+c.FullPath()]
		if ok {
			credit = api.Credit
		}
		traceID := header.GetTraceID(c)
		header.SetTraceID(c, traceID)
		ctx := &Context[T]{
			GinContext: c,
			Request:    ginhelp.GinRequest[T](c),
			TraceID:    traceID,
		}
		if handlerSetup.Pagination {
			ctx.Page = ginhelp.GinRequest[sql.Pagination](c)
		}
		if handlerSetup.Sort {
			ctx.Sort = ginhelp.GinRequest[sql.Sort](c)
		}
		handlerSetup.Service(ctx)
		traces := header.GetTraces(c)
		overrideCredit := ctx.GinContext.GetFloat64("credit")
		if overrideCredit > 0 {
			credit = overrideCredit
		}

		resp, ok := ctx.Response.(types.Response)
		if !ok {
			panic("response is nil")
		}
		if !resp.Success {
			traces = append(traces, types.Trace{
				Success:    false,
				Time:       time.Now().Format("2006-01-02 15:04:05"),
				SystemUUID: system.SYSTEM_UUID,
				SystemName: system.SYSTEM_NAME,
				TraceID:    traceID,
				Duration:   uint(time.Since(now).Microseconds()),
				Credit:     credit,
				Error:      resp.Error,
			})
			if resp.Error.Code == constant.ERR_CODE_UNAUTHORIZED {
				ctx.GinContext.JSON(401, resp)
				return
			} else if resp.Error.Code == constant.ERR_CODE_FORBIDDEN {
				ctx.GinContext.JSON(403, resp)
				return
			} else if resp.Error.Code == constant.ERR_CODE_WORKSPACE_NOT_FOUND {
				ctx.GinContext.JSON(403, resp)
				return
			} else if resp.Error.Code == constant.ERR_CODE_INTERNAL_SERVER_ERROR {
				ctx.GinContext.JSON(500, resp)
				return
			}
			resp.Traces = traces
			ctx.GinContext.JSON(200, resp)
			return
		}
		traces = append(traces, types.Trace{
			Success:    true,
			Time:       time.Now().Format("2006-01-02 15:04:05"),
			TraceID:    traceID,
			Duration:   uint(time.Since(now).Microseconds()),
			Credit:     credit,
			SystemUUID: system.SYSTEM_UUID,
			SystemName: system.SYSTEM_NAME,
		})
		resp.Traces = traces
		ctx.GinContext.JSON(200, resp)
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
			Request:    ginhelp.GinRequest[T](c),
		}
		err1 := handlerSetup.Service(ctx, ws)
		if err != nil {
			if err1.Code() == constant.ERR_CODE_UNAUTHORIZED {
				ctx.GinContext.JSON(401, nil)
				return
			} else if err1.Code() == constant.ERR_CODE_FORBIDDEN {
				ctx.GinContext.JSON(403, nil)
				return
			} else if err1.Code() == constant.ERR_CODE_WORKSPACE_NOT_FOUND {
				ctx.GinContext.JSON(403, nil)
				return
			} else if err1.Code() == constant.ERR_CODE_INTERNAL_SERVER_ERROR {
				ctx.GinContext.JSON(500, nil)
				return
			}
			ctx.GinContext.JSON(200, nil)
			return
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
