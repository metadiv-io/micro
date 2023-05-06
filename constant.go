package micro

import "github.com/metadiv-io/env"

const (
	tag_uri  = "uri"
	tag_json = "json"
	tag_form = "form"
)

const (
	MICRO_HEADER_TRACE_ID = "Micro-TraceID"
	MICRO_HEADER_TRACES   = "Micro-Traces"
)

var (
	USER_TOKEN_PUBLIC_PEM   string
	SYSTEM_TOKEN_PUBLIC_PEM string
)

var (
	SYSTEM_UUID       string
	SYSTEM_NAME       string
	SYSTEM_TOKEN      string
	AUTH_SERVICE_URL  string
	USAGE_SERVICE_URL string
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
	SYSTEM_TOKEN = env.String("SYSTEM_TOKEN", "")
	if SYSTEM_TOKEN == "" {
		panic("SYSTEM_TOKEN is required")
	}
	AUTH_SERVICE_URL = env.String("AUTH_SERVICE_URL", "")
	if AUTH_SERVICE_URL == "" {
		panic("AUTH_SERVICE_URL is required")
	}
	USAGE_SERVICE_URL = env.String("USAGE_SERVICE_URL", "")
	if USAGE_SERVICE_URL == "" {
		panic("USAGE_SERVICE_URL is required")
	}
}
