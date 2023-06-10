package system

import "github.com/metadiv-io/env"

var (
	AUTH_SERVICE_URL string
)

func init() {
	AUTH_SERVICE_URL = env.String("AUTH_SERVICE_URL", "")
	if AUTH_SERVICE_URL == "" {
		panic("AUTH_SERVICE_URL is required")
	}
}

var (
	USAGE_SERVICE_URL string
)

func init() {
	USAGE_SERVICE_URL = env.String("USAGE_SERVICE_URL", "")
	if USAGE_SERVICE_URL == "" {
		panic("USAGE_SERVICE_URL is required")
	}
}
