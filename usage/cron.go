package usage

import (
	"github.com/metadiv-io/logger"
	"github.com/metadiv-io/micro/call"
	"github.com/metadiv-io/micro/system"
)

type Consumption struct {
	SubscriptionUUID string  `json:"subscription_uuid"`
	Credit           float64 `json:"credit"`
}

type SendConsumptionRequest struct {
	Consumptions []Consumption `json:"consumptions"`
}

func SendConsumptionCron() {
	var consumptions []Consumption
	for subscriptionUUID, consumption := range cachedConsumption {
		consumptions = append(consumptions, Consumption{
			SubscriptionUUID: subscriptionUUID,
			Credit:           consumption,
		})
	}
	if len(consumptions) > 0 {
		resp, err := call.POST[SendConsumptionRequest](nil, system.USAGE_SERVICE_URL+"/usage", SendConsumptionRequest{
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
