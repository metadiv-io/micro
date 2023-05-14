package usage

import "github.com/metadiv-io/env"

var (
	USAGE_SERVICE_URL string
)

func init() {
	USAGE_SERVICE_URL = env.String("USAGE_SERVICE_URL", "")
	if USAGE_SERVICE_URL == "" {
		panic("USAGE_SERVICE_URL is required")
	}
}
