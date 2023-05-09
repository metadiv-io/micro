package micro

import "github.com/metadiv-io/env"

const (
	tag_uri  = "uri"
	tag_json = "json"
	tag_form = "form"
)

// These are the headers for micro services.
const (
	MICRO_HEADER_TRACE_ID  = "Micro-TraceID"
	MICRO_HEADER_TRACES    = "Micro-Traces"
	MICRO_HEADER_WORKSPACE = "Micro-Workspace"
)

// These are the base info for the micro service.
var (
	SYSTEM_UUID string
	SYSTEM_NAME string
)

func init() {
	SYSTEM_UUID = env.String("SYSTEM_UUID", "")
	if SYSTEM_UUID == "" {
		panic("SYSTEM_UUID is required")
	}
	SYSTEM_NAME = env.String("SYSTEM_NAME", "")
	if SYSTEM_NAME == "" {
		panic("SYSTEM_NAME is required")
	}
}

// This map store the api uuid and info.
var API_MAP = make(map[string]Api)

type Api struct {
	Tag    string  `json:"tag"`
	UUID   string  `json:"uuid"`
	Credit float64 `json:"credit"`
}
