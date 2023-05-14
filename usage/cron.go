package usage

import (
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/micro"
	"github.com/metadiv-io/micro/call"
)

type Consumption struct {
	WorkspaceUUID string  `json:"workspace_uuid"`
	Credit        float64 `json:"credit"`
}

type SendConsumptionRequest struct {
	Consumptions []Consumption `json:"consumptions"`
}

func sendConsumptionCron() {
	var consumptions []Consumption
	for workspaceUUID, consumption := range cachedConsumption {
		consumptions = append(consumptions, Consumption{
			WorkspaceUUID: workspaceUUID,
			Credit:        consumption,
		})
	}
	if len(consumptions) > 0 {
		resp, err := call.POST[SendConsumptionRequest](nil, USAGE_SERVICE_URL+"/usage", SendConsumptionRequest{
			Consumptions: consumptions,
		}, nil)
		if err != nil {
			logger.Error("send consumption cron:", err.Error())
			return
		}
		if !resp.Success {
			logger.Error("send consumption cron:", resp.Error.Message)
			return
		}
		cachedConsumption = make(map[string]float64)
	}
}

func SetupSendUsageCron(e *micro.Engine) {
	micro.Cron(e, "@every 5m", sendConsumptionCron)
}
