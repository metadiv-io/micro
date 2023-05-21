package usage

import (
	"github.com/metadiv-io/env"
	"github.com/metadiv-io/err_map"
)

var (
	USAGE_SERVICE_URL string
)

func init() {
	USAGE_SERVICE_URL = env.String("USAGE_SERVICE_URL", "")
	if USAGE_SERVICE_URL == "" {
		panic("USAGE_SERVICE_URL is required")
	}
}

const (
	ERR_CODE_WORKSPACE_NOT_FOUND = "381170ba-9ecb-4f4c-aa11-a6e777a00968"
	ERR_CODE_API_UUID_NOT_FOUND  = "bff74b7c-09e3-4219-a2c2-1727e248c1f7"
	ERR_CODE_NOT_ENOUGH_CREDIT   = "7c04a968-6827-4332-bdfd-3642b8bda40e"
)

func init() {
	err_map.Register(ERR_CODE_WORKSPACE_NOT_FOUND, "Workspace Not Found")
	err_map.Register(ERR_CODE_API_UUID_NOT_FOUND, "Api UUID Not Found")
	err_map.Register(ERR_CODE_NOT_ENOUGH_CREDIT, "Not Enough Credit")
}
