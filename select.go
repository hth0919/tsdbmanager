package tsdbmanager

import (
	parser "github.com/hth0919/tsdbparser"
	"net/http"

)

func(m *TSManager)SelectMetric(MetricName string) float64{
	query :="/api/query?"
	start := "start=1h-ago"
	metric := "m=sum:"+MetricName
	cluster := "Cluster="+m.Client.Cluster
	node := "Node="+m.Client.Node
	pod := "Pod="+m.Client.Pod
	rest := m.Client.Host + query + start + "&" + metric + "{" + cluster + "," + node + "," + pod + "}"

	resp, err := http.Get(rest)
	if err != nil {
		panic(err)
	}
	m.OMetric = *parser.JsonUnmarshaller(resp.Body)
	defer resp.Body.Close()

	return parser.GetLast(&m.OMetric)
}

func(m *TSManager)SelectMetricWithCNPName(Cluster string, Node string, Pod string,MetricName string) float64{
	m.SetCNPName(Cluster,Node,Pod,output)
	return m.SelectMetric(MetricName)
}
