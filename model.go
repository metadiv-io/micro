package micro

import (
	"github.com/gorilla/websocket"
	"github.com/metadiv-io/err_map"
)

type Service[T any] func(ctx *Context[T])

type Handler[T any] func() HandlerResponse[T]

type HandlerResponse[T any] struct {
	Service    Service[T]
	Pagination bool
	Sort       bool
}

// Websocket

type WSService[T any] func(ctx *Context[T], ws *websocket.Conn) err_map.Error

type WSHandler[T any] func() WSHandlerResponse[T]

type WSHandlerResponse[T any] struct {
	Service WSService[T]
}
