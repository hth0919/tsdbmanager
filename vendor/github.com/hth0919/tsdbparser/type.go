package tsdbParser

type Metric struct {
	MetricName 	string `json:"metric"`
	Tags   		map[string]string `json:"tags"`
	Dps 		map[int64]float64 `json:"dps"`
}
