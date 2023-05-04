package auth

import (
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/call"
)

type AuthRequest struct {
	AuthMap map[string][]string `json:"auth_map"`
}

type AuthResponse struct {
	Allowed bool `json:"allowed"`
}

func isAllowed(authMap map[string][]string) bool {
	resp, err := call.POST[AuthResponse](micro.AUTH_SERVICE_URL+"/micro/allowed", &AuthRequest{
		AuthMap: authMap,
	}, map[string]string{
		"Authorization": "Bearer " + micro.SYSTEM_TOKEN,
	}, "", nil)
	if err != nil || resp == nil || !resp.Success {
		return false
	}
	return resp.Data.Allowed
}

type UsageRequest struct {
	ApiUUID string              `json:"api_uuid"`
	AuthMap map[string][]string `json:"auth_map"`
}

type UsageResponse struct {
	Allowed bool `json:"allowed"`
}

func isUsageAllowed(apiUUID string, authMap map[string][]string) bool {
	resp, err := call.POST[UsageResponse](micro.USAGE_SERVICE_URL+"/micro/allowed", &UsageRequest{
		ApiUUID: apiUUID,
		AuthMap: authMap,
	}, map[string]string{
		"Authorization": "Bearer " + micro.SYSTEM_TOKEN,
	}, "", nil)
	if err != nil || resp == nil || !resp.Success {
		return false
	}
	return resp.Data.Allowed
}
