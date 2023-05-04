package call

import (
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/sql"
)

type Response[T any] struct {
	Success    bool            `json:"success"`
	TraceID    string          `json:"trace_id"`
	Duration   uint            `json:"duration"`
	Credit     float64         `json:"credit"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Error      *micro.Error    `json:"error,omitempty"`
	Data       *T              `json:"data,omitempty"`
	Traces     []micro.Trace   `json:"traces,omitempty"`
}
