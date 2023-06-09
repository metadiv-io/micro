package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/ginmid"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/auth"
)

func UserGET[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UserCacheGET[T any](engine *micro.Engine, route string, cacheDuration time.Duration, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(
		ginmid.Cache(cacheDuration, newGinServiceHandler(engine, handler)), newMiddleware...)...)
}

func UserPOST[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.POST(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UserPUT[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.PUT(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UserDELETE[T any](engine *micro.Engine, route string, handler micro.Handler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.DELETE(route, joinMiddlewareAndService(newGinServiceHandler(engine, handler), newMiddleware...)...)
}

func UserWS[T any](engine *micro.Engine, route string, handler micro.WSHandler[T], middleware ...gin.HandlerFunc) {
	newMiddleware := append([]gin.HandlerFunc{auth.UserOnly}, middleware...)
	engine.GinEngine.GET(route, joinMiddlewareAndService(newGinWSServiceHandler(engine, handler), newMiddleware...)...)
}
