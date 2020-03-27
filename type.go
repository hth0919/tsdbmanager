package tsdbmanager

import parser "github.com/hth0919/tsdbparser"

const (
	input  IOType = 0
	output IOType = 1
)

type TSManager struct {
	Client *TSClient
	OMetric parser.Metric
	IMetric IMetric
}

type IMetric struct {
	Metric string `json:"metric"`
	Timestamp int64 `json:"timestamp"`
	Value interface{} `json:"value"`
	Tags map[string]string `json:"tags"`
}

type TSClient struct {
	Host string
	Cluster string
	Node string
	Pod string
}
type IOType int

var iotype = [...]int{
	0,1,
}
func (m IOType) String() int {
	return iotype[m%2]
}
