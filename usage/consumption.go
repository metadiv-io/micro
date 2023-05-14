package usage

var cachedConsumption = make(map[string]float64)

func addConsumption(subscriptionUUID string, consumption float64) {
	_, ok := cachedConsumption[subscriptionUUID]
	if !ok {
		cachedConsumption[subscriptionUUID] = consumption
	} else {
		cachedConsumption[subscriptionUUID] += consumption
	}
}
