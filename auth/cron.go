package auth

import (
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/call"
)

type RegisterRequest struct {
	SystemUUID string               `json:"system_uuid"`
	SystemName string               `json:"system_name"`
	ApiMap     map[string]micro.Api `json:"api_map"`
}

type RegisterResponse struct {
	ApiMap       map[string]micro.Api `json:"api_map"`
	JwtPublicPem string               `json:"jwt_public_pem"`
}

func registerCron() {
	resp, err := call.POST[RegisterResponse](nil, AUTH_SERVICE_URL+"/micro/ping", &RegisterRequest{
		SystemUUID: micro.SYSTEM_UUID,
		SystemName: micro.SYSTEM_NAME,
		ApiMap:     micro.API_MAP,
	}, map[string]string{})
	if err != nil || resp == nil || !resp.Success {
		return
	}
	micro.API_MAP = resp.Data.ApiMap
	JWT_PUBLIC_PEM = resp.Data.JwtPublicPem
}

func SetupRegisterCron(e *micro.Engine) {
	micro.Cron(e, "@every 1m", registerCron)
}
