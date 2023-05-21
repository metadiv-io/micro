package usage

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/err_map"
	"github.com/metadiv-io/micro"
)

func AbortWorkspaceNotFound(ctx *gin.Context) {
	traceID := micro.GetTraceID(ctx)
	traces := micro.GetTraces(ctx)
	err := err_map.NewError(ERR_CODE_WORKSPACE_NOT_FOUND)
	traces = append(traces, micro.Trace{
		Success:    false,
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SystemUUID: micro.SYSTEM_UUID,
		SystemName: micro.SYSTEM_NAME,
		TraceID:    traceID,
		Error: &micro.Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
	})
	var credit float64
	var duration uint
	for _, t := range traces {
		credit += t.Credit
		duration += t.Duration
	}
	ctx.AbortWithStatusJSON(403, micro.Response{
		Success: false,
		Error: &micro.Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
		TraceID:  traceID,
		Traces:   traces,
		Credit:   credit,
		Duration: duration,
	})
}

func AbortApiUUIDNotFound(ctx *gin.Context) {
	traceID := micro.GetTraceID(ctx)
	traces := micro.GetTraces(ctx)
	err := err_map.NewError(ERR_CODE_WORKSPACE_NOT_FOUND)
	traces = append(traces, micro.Trace{
		Success:    false,
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SystemUUID: micro.SYSTEM_UUID,
		SystemName: micro.SYSTEM_NAME,
		TraceID:    traceID,
		Error: &micro.Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
	})
	var credit float64
	var duration uint
	for _, t := range traces {
		credit += t.Credit
		duration += t.Duration
	}
	ctx.AbortWithStatusJSON(403, micro.Response{
		Success: false,
		Error: &micro.Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
		TraceID:  traceID,
		Traces:   traces,
		Credit:   credit,
		Duration: duration,
	})
}

func AbortNotEnoughCredit(ctx *gin.Context) {
	traceID := micro.GetTraceID(ctx)
	traces := micro.GetTraces(ctx)
	err := err_map.NewError(ERR_CODE_NOT_ENOUGH_CREDIT)
	traces = append(traces, micro.Trace{
		Success:    false,
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SystemUUID: micro.SYSTEM_UUID,
		SystemName: micro.SYSTEM_NAME,
		TraceID:    traceID,
		Error: &micro.Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
	})
	var credit float64
	var duration uint
	for _, t := range traces {
		credit += t.Credit
		duration += t.Duration
	}
	ctx.AbortWithStatusJSON(403, micro.Response{
		Success: false,
		Error: &micro.Error{
			Code:    err.Code(),
			Message: err.Error(),
		},
		TraceID:  traceID,
		Traces:   traces,
		Credit:   credit,
		Duration: duration,
	})
}
