package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/auth"
	"github.com/metadiv-io/micro/usage"
)

func UsageGET[T any](engine *micro.Engine, route string, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	micro.API_MAP["GET:"+route] = micro.Api{
		Tag:  "GET:" + route,
		UUID: uuid,
	}
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UsageCacheGET[T any](engine *micro.Engine, route string, uuid string, cacheDuration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	micro.API_MAP["GET:"+route] = micro.Api{
		Tag:  "GET:" + route,
		UUID: uuid,
	}
	engine.GinEngine.GET(route, joinMiddlewareAndService(
		ginmid.Cache(cacheDuration, newGinServiceHandler(engine, handler)), newMiddleware...)...)
}

func UsagePOST[T any](engine *micro.Engine, route string, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	micro.API_MAP["POST:"+route] = micro.Api{
		Tag:  "POST:" + route,
		UUID: uuid,
	}
	engine.GinEngine.POST(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UsagePUT[T any](engine *micro.Engine, route string, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	micro.API_MAP["PUT:"+route] = micro.Api{
		Tag:  "PUT:" + route,
		UUID: uuid,
	}
	engine.GinEngine.PUT(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UsageDELETE[T any](engine *micro.Engine, route string, uuid string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	micro.API_MAP["DELETE:"+route] = micro.Api{
		Tag:  "DELETE:" + route,
		UUID: uuid,
	}
	engine.GinEngine.DELETE(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UsageWS[T any](engine *micro.Engine, route string, uuid string, handler micro.WSHandler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly, usage.UsageRequired}, middleware...)
	micro.API_MAP["WS:"+route] = micro.Api{
		Tag:  "WS:" + route,
		UUID: uuid,
	}
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinWSServiceHandler(engine, handler), newMiddleware...)...)
}
