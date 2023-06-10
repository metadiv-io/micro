package micro

import (
	"github.com/metadiv-io/micro/system"
	"github.com/metadiv-io/micro/types"
)

type Pong struct {
	SystemUUID string               `json:"system_uuid"`
	SystemName string               `json:"system_name"`
	ApiMap     map[string]types.Api `json:"api_map"`
}

func PingHandler() HandlerResponse[struct{}] {
	return HandlerResponse[struct{}]{
		Service: func(ctx *Context[struct{}]) {
			ctx.OK(&Pong{
				SystemUUID: system.SYSTEM_UUID,
				SystemName: system.SYSTEM_NAME,
				ApiMap:     types.API_MAP,
			})
		},
	}
}
