package ginhelp

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/err_map"
	"github.com/metadiv-io/micro/constant"
	"github.com/metadiv-io/micro/header"
	"github.com/metadiv-io/micro/system"
	"github.com/metadiv-io/micro/types"
)

func AbortUnauthorized(ctx *gin.Context) {
	traceID := header.GetTraceID(ctx)
	traces := header.GetTraces(ctx)
	err := err_map.NewError(constant.ERR_CODE_UNAUTHORIZED)
	traces = append(traces, types.Trace{
		Success:    false,
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SystemUUID: system.SYSTEM_UUID,
		SystemName: system.SYSTEM_NAME,
		TraceID:    traceID,
		Error: &types.ErrorImpl{
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
	ctx.AbortWithStatusJSON(401, types.Response{
		Success: false,
		Error: &types.ErrorImpl{
			Code:    err.Code(),
			Message: err.Error(),
		},
		TraceID:  traceID,
		Traces:   traces,
		Credit:   credit,
		Duration: duration,
	})
}

func AbortForbidden(ctx *gin.Context) {
	traceID := header.GetTraceID(ctx)
	traces := header.GetTraces(ctx)
	err := err_map.NewError(constant.ERR_CODE_FORBIDDEN)
	traces = append(traces, types.Trace{
		Success:    false,
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SystemUUID: system.SYSTEM_UUID,
		SystemName: system.SYSTEM_NAME,
		TraceID:    traceID,
		Error: &types.ErrorImpl{
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
	ctx.AbortWithStatusJSON(403, types.Response{
		Success: false,
		Error: &types.ErrorImpl{
			Code:    err.Code(),
			Message: err.Error(),
		},
		TraceID:  traceID,
		Traces:   traces,
		Credit:   credit,
		Duration: duration,
	})
}

func AbortWorkspaceNotFound(ctx *gin.Context) {
	traceID := header.GetTraceID(ctx)
	traces := header.GetTraces(ctx)
	err := err_map.NewError(constant.ERR_CODE_WORKSPACE_NOT_FOUND)
	traces = append(traces, types.Trace{
		Success:    false,
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SystemUUID: system.SYSTEM_UUID,
		SystemName: system.SYSTEM_NAME,
		TraceID:    traceID,
		Error: &types.ErrorImpl{
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
	ctx.AbortWithStatusJSON(403, types.Response{
		Success: false,
		Error: &types.ErrorImpl{
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
	traceID := header.GetTraceID(ctx)
	traces := header.GetTraces(ctx)
	err := err_map.NewError(constant.ERR_CODE_WORKSPACE_NOT_FOUND)
	traces = append(traces, types.Trace{
		Success:    false,
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SystemUUID: system.SYSTEM_UUID,
		SystemName: system.SYSTEM_NAME,
		TraceID:    traceID,
		Error: &types.ErrorImpl{
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
	ctx.AbortWithStatusJSON(403, types.Response{
		Success: false,
		Error: &types.ErrorImpl{
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
	traceID := header.GetTraceID(ctx)
	traces := header.GetTraces(ctx)
	err := err_map.NewError(constant.ERR_CODE_NOT_ENOUGH_CREDIT)
	traces = append(traces, types.Trace{
		Success:    false,
		Time:       time.Now().Format("2006-01-02 15:04:05"),
		SystemUUID: system.SYSTEM_UUID,
		SystemName: system.SYSTEM_NAME,
		TraceID:    traceID,
		Error: &types.ErrorImpl{
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
	ctx.AbortWithStatusJSON(403, types.Response{
		Success: false,
		Error: &types.ErrorImpl{
			Code:    err.Code(),
			Message: err.Error(),
		},
		TraceID:  traceID,
		Traces:   traces,
		Credit:   credit,
		Duration: duration,
	})
}
