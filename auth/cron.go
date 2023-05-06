package auth

import (
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/call"
)

type PingRequest struct {
	SystemUUID string               `json:"system_uuid"`
	SystemName string               `json:"system_name"`
	ApiMap     map[string]micro.Api `json:"api_map"`
}

type PingResponse struct {
	ApiMap          map[string]micro.Api `json:"api_map"`
	SystemPublicPem string               `json:"system_public_pem"`
	UserPublicPem   string               `json:"user_public_pem"`
}

func pingAuthCron() {
	resp, err := call.POST[PingResponse](micro.AUTH_SERVICE_URL+"/micro/ping", &PingRequest{
		SystemUUID: micro.SYSTEM_UUID,
		SystemName: micro.SYSTEM_NAME,
		ApiMap:     micro.API_MAP,
	}, map[string]string{
		"Authorization": "Bearer " + micro.SYSTEM_TOKEN,
	}, "", nil)
	if err != nil || resp == nil || !resp.Success {
		return
	}
	micro.API_MAP = resp.Data.ApiMap
	SYSTEM_PUBLIC_PEM = resp.Data.SystemPublicPem
	USER_PUBLIC_PEM = resp.Data.UserPublicPem
}

func SetupPingAuthCron(e *micro.Engine) {
	micro.Cron(e, "@every 1m", pingAuthCron)
}

func pingUsageCron() {
	resp, err := call.POST[PingResponse](micro.USAGE_SERVICE_URL+"/micro/ping", &PingRequest{
		SystemUUID: micro.SYSTEM_UUID,
		SystemName: micro.SYSTEM_NAME,
		ApiMap:     micro.API_MAP,
	}, map[string]string{
		"Authorization": "Bearer " + micro.SYSTEM_TOKEN,
	}, "", nil)
	if err != nil || resp == nil || !resp.Success {
		return
	}
	micro.API_MAP = resp.Data.ApiMap
}

func SetupPingUsageCron(e *micro.Engine) {
	micro.Cron(e, "@every 1m", pingUsageCron)
}
