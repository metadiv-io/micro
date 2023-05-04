package micro

import (
	"github.com/metadiv-io/err_map"
)

type Pong struct {
	SystemUUID string         `json:"system_uuid"`
	SystemName string         `json:"system_name"`
	ApiMap     map[string]Api `json:"api_map"`
}

func PingHandler() HandlerResponse[struct{}] {
	return HandlerResponse[struct{}]{
		Service: func(ctx *Context[struct{}]) (interface{}, err_map.Error) {
			return &Pong{
				SystemUUID: SYSTEM_UUID,
				SystemName: SYSTEM_NAME,
				ApiMap:     API_MAP,
			}, nil
		},
	}
}
