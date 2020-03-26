package tsdbmanager

import parser "github.com/hth0919/tsdbparser"

type TSManager struct {
	Client *TSClient
	OMetric parser.Metric
	IMetric IMetric
}

type IMetric struct {
	Metric string `json:"metric"`
	Timestamp int64 `json:"timestamp"`
	Value float64 `json:"value"`
	Tags map[string]string `json:"tags"`
}

type TSClient struct {
	Host string
	Cluster string
	Node string
	Pod string
}
type IOType int
const (
	input  IOType = 0
	output IOType = 1
)

