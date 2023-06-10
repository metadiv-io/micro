package call

import (
	"github.com/metadiv-io/micro/types"
	"github.com/metadiv-io/sql"
)

type Response[T any] struct {
	Success    bool             `json:"success"`
	TraceID    string           `json:"trace_id"`
	Duration   uint             `json:"duration"`
	Credit     float64          `json:"credit"`
	Pagination *sql.Pagination  `json:"pagination,omitempty"`
	Error      *types.ErrorImpl `json:"error,omitempty"`
	Data       *T               `json:"data,omitempty"`
	Traces     []types.Trace    `json:"traces,omitempty"`
}
