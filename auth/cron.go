package auth

import (
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/micro/call"
	"github.com/metadiv-io/micro/system"
	"github.com/metadiv-io/micro/types"
)

type RegisterRequest struct {
	SystemUUID string               `json:"system_uuid"`
	SystemName string               `json:"system_name"`
	ApiMap     map[string]types.Api `json:"api_map"`
}

type RegisterResponse struct {
	ApiMap       map[string]types.Api `json:"api_map"`
	JwtPublicPem string               `json:"jwt_public_pem"`
}

func RegisterCron() {
	resp, err := call.POST[RegisterResponse](nil, system.AUTH_SERVICE_URL+"/register", &RegisterRequest{
		SystemUUID: system.SYSTEM_UUID,
		SystemName: system.SYSTEM_NAME,
		ApiMap:     types.API_MAP,
	}, map[string]string{})
	if err != nil {
		logger.Error("register cron", err.Error())
		return
	}
	if resp == nil {
		logger.Error("register cron", "response is nil")
		return
	}
	if !resp.Success && resp.Error != nil {
		logger.Error("register cron", resp.Error.Message)
		return
	}
	types.API_MAP = resp.Data.ApiMap
	system.JWT_PUBLIC_PEM = resp.Data.JwtPublicPem
}

// func SetupRegisterCron(e *engine.Engine, initExec bool) {
// 	if initExec {
// 		registerCron()
// 	}
// 	micro.Cron(e, "@every 1m", registerCron)
// }
