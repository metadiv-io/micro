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
	}, map[string]string{
		"Authorization": "Bearer " + micro.SYSTEM_TOKEN,
	}, "", nil)
	if err != nil || resp == nil || !resp.Success {
		return false
	}
	return resp.Data.Allowed
}
