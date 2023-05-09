package micro

import (
	"github.com/gorilla/websocket"
	"github.com/metadiv-io/err_map"
	"github.com/metadiv-io/sql"
)

// Response, Trace and Error

type Response struct {
	Success    bool            `json:"success"`
	TraceID    string          `json:"trace_id"`
	Duration   uint            `json:"duration"`
	Credit     float64         `json:"credit"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Error      *Error          `json:"error,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
	Traces     []Trace         `json:"traces,omitempty"`
}

type Trace struct {
	Success    bool    `json:"success"`
	SystemUUID string  `json:"system_uuid"`
	SystemName string  `json:"system_name"`
	TraceID    string  `json:"trace_id"`
	Time       string  `json:"time"`
	Duration   uint    `json:"duration"`
	Credit     float64 `json:"credit"`
	Error      *Error  `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// REST

type Service[T any] func(ctx *Context[T]) (interface{}, err_map.Error)

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
