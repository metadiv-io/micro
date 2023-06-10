package types

type Trace struct {
	Success    bool       `json:"success"`
	SystemUUID string     `json:"system_uuid"`
	SystemName string     `json:"system_name"`
	TraceID    string     `json:"trace_id"`
	Time       string     `json:"time"`
	Duration   uint       `json:"duration"`
	Credit     float64    `json:"credit"`
	Error      *ErrorImpl `json:"error,omitempty"`
}
