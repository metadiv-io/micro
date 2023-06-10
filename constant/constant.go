package constant

import "github.com/metadiv-io/err_map"

const (
	MICRO_HEADER_TRACE_ID  = "Micro-TraceID"
	MICRO_HEADER_TRACES    = "Micro-Traces"
	MICRO_HEADER_WORKSPACE = "Micro-Workspace"
)

const (
	ERR_CODE_UNAUTHORIZED          = "b97cf20d-42b6-470e-9e08-b4bb852c3811"
	ERR_CODE_FORBIDDEN             = "7792176d-0196-4a57-a959-93062c2b9b41"
	ERR_CODE_INTERNAL_SERVER_ERROR = "b6a82bc6-5884-41e1-8b6f-1a013b7da835"
	ERR_CODE_WORKSPACE_NOT_FOUND   = "94793665-c9da-48a6-84bb-3dd2fd771419"
)

func init() {
	err_map.Register(ERR_CODE_UNAUTHORIZED, "Unauthorized")
	err_map.Register(ERR_CODE_FORBIDDEN, "Forbidden")
	err_map.Register(ERR_CODE_INTERNAL_SERVER_ERROR, "Internal Server Error")
	err_map.Register(ERR_CODE_WORKSPACE_NOT_FOUND, "Workspace Not Found")
}

const (
	ERR_CODE_API_UUID_NOT_FOUND = "bff74b7c-09e3-4219-a2c2-1727e248c1f7"
	ERR_CODE_NOT_ENOUGH_CREDIT  = "7c04a968-6827-4332-bdfd-3642b8bda40e"
)

func init() {
	err_map.Register(ERR_CODE_API_UUID_NOT_FOUND, "Api UUID Not Found")
	err_map.Register(ERR_CODE_NOT_ENOUGH_CREDIT, "Not Enough Credit")
}
