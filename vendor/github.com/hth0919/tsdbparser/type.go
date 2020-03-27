package tsdbParser

type Metric struct {
	MetricName 	string `json:"metric"`
	Tags   		map[string]string `json:"tags"`
	Dps 		map[string]interface{} `json:"dps"`
}
