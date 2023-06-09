package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/auth"
)

func AdminGET[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func AdminCacheGET[T any](engine *micro.Engine, route string, cacheDuration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(
		ginmid.Cache(cacheDuration, newGinServiceHandler(engine, handler)), newMiddleware...)...)
}

func AdminPOST[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.POST(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func AdminPUT[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.PUT(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func AdminDELETE[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.DELETE(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func AdminWS[T any](engine *micro.Engine, route string, handler micro.WSHandler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.AdminOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinWSServiceHandler(engine, handler), newMiddleware...)...)
}
