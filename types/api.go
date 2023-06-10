package types

type Api struct {
	Tag    string  `json:"tag"`
	UUID   string  `json:"uuid"`
	Credit float64 `json:"credit"`
}

var API_MAP = make(map[string]Api)
