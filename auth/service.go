package auth

import (
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/call"
)

type UsageRequest struct {
	ApiUUID    string   `json:"api_uuid"`
	Workspaces []string `json:"workspaces"`
}

type UsageResponse struct {
	Allowed bool `json:"allowed"`
}

func isUsageAllowed(apiUUID string, workspaces []string) bool {
	resp, err := call.POST[UsageResponse](micro.USAGE_SERVICE_URL+"/micro/allowed", &UsageRequest{
		ApiUUID:    apiUUID,
		Workspaces: workspaces,
	}, map[string]string{}, "", nil)
	if err != nil || resp == nil || !resp.Success {
		return false
	}
	return resp.Data.Allowed
}

type SystemAllowedRequest struct {
	IP string `json:"ip"`
}

type SystemAllowedResponse struct {
	Allowed bool `json:"allowed"`
}

func isSystemAllowed(ip string) bool {
	resp, err := call.POST[SystemAllowedResponse](micro.AUTH_SERVICE_URL+"/micro/system/allowed", &SystemAllowedRequest{
		IP: ip,
	}, map[string]string{}, "", nil)
	if err != nil || resp == nil || !resp.Success {
		return false
	}
	return resp.Data.Allowed
}
