package usage

var cachedConsumption = make(map[string]float64)

func addConsumption(workspaceUUID string, consumption float64) {
	_, ok := cachedConsumption[workspaceUUID]
	if !ok {
		cachedConsumption[workspaceUUID] = consumption
	} else {
		cachedConsumption[workspaceUUID] += consumption
	}
}
