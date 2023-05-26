package micro

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTraceID(c *gin.Context) string {
	traceID := c.GetHeader(MICRO_HEADER_TRACE_ID)
	if traceID == "" {
		return uuid.NewString()
	}
	return traceID
}

func SetTraceID(c *gin.Context, traceID string) {
	c.Request.Header.Set(MICRO_HEADER_TRACE_ID, traceID)
}

func GetTraces(c *gin.Context) []Trace {
	var traces []Trace
	traceHeader := c.GetHeader(MICRO_HEADER_TRACES)
	if traceHeader != "" {
		_ = json.Unmarshal([]byte(traceHeader), &traces)
	}
	if len(traces) == 0 {
		traces = make([]Trace, 0)
	}
	return traces
}

func SetTraces(c *gin.Context, traces []Trace) {
	b, _ := json.Marshal(traces)
	c.Request.Header.Set(MICRO_HEADER_TRACES, string(b))
}

func GetWorkspace(c *gin.Context) string {
	return c.GetHeader(MICRO_HEADER_WORKSPACE)
}

func SetWorkspace(c *gin.Context, workspace string) {
	c.Request.Header.Set(MICRO_HEADER_WORKSPACE, workspace)
}

func GetApiUUID(c *gin.Context) string {
	info, ok := API_MAP[c.Request.Method+":"+c.FullPath()]
	if ok {
		return info.UUID
	}
	return ""
}
