package types

import "github.com/metadiv-io/sql"

type Response struct {
	Success    bool            `json:"success"`
	TraceID    string          `json:"trace_id"`
	Duration   uint            `json:"duration"`
	Credit     float64         `json:"credit"`
	Pagination *sql.Pagination `json:"pagination,omitempty"`
	Error      *ErrorImpl      `json:"error,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
	Traces     []Trace         `json:"traces,omitempty"`
}
