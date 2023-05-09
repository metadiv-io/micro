package call

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/micro"
)

func get[T any](ctx *gin.Context, url string, params map[string]string, headers map[string]string) (*Response[T], error) {
	var traceID, workspace string
	var traces []micro.Trace
	if ctx != nil {
		traceID = micro.GetTraceID(ctx)
		traces = micro.GetTraces(ctx)
		workspace = micro.GetWorkspace(ctx)
	}
	if headers == nil {
		headers = make(map[string]string)
	}
	headers[micro.MICRO_HEADER_TRACE_ID] = traceID
	tracesStr, _ := json.Marshal(traces)
	headers[micro.MICRO_HEADER_TRACES] = string(tracesStr)
	headers[micro.MICRO_HEADER_WORKSPACE] = workspace
	url += "?"
	for k, v := range params {
		url += k + "=" + v + "&"
	}
	url = url[:len(url)-1]

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var response Response[T]
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		micro.SetTraceID(ctx, traceID)
		micro.SetTraces(ctx, traces)
		micro.SetWorkspace(ctx, workspace)
	}

	return &response, nil
}

func nonGet[T any](ctx *gin.Context, url string, method string, body interface{}, headers map[string]string) (*Response[T], error) {
	var traceID, workspace string
	var traces []micro.Trace
	if ctx != nil {
		traceID = micro.GetTraceID(ctx)
		traces = micro.GetTraces(ctx)
		workspace = micro.GetWorkspace(ctx)
	}
	if headers == nil {
		headers = make(map[string]string)
	}
	headers[micro.MICRO_HEADER_TRACE_ID] = traceID
	tracesStr, _ := json.Marshal(traces)
	headers[micro.MICRO_HEADER_TRACES] = string(tracesStr)
	headers[micro.MICRO_HEADER_WORKSPACE] = workspace
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var response Response[T]
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		micro.SetTraceID(ctx, traceID)
		micro.SetTraces(ctx, traces)
		micro.SetWorkspace(ctx, workspace)
	}

	return &response, nil
}
