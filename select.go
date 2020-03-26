package tsdbmanager

import (
	parser "github.com/hth0919/tsdbparser"
	"io/ioutil"
	"net/http"

)

func(m *TSManager)SelectMetric(MetricName string) float64{
	query :="/api/query?"
	start := "1h-ago"
	metric := "m=sum:"+MetricName
	cluster := "Cluster="+m.Client.Cluster
	node := "Node="+m.Client.Node
	pod := "Pod="+m.Client.Pod
	rest := m.Client.Host + query + start + "&" + metric + "{" + cluster + "," + node + "," + pod + "}"

	resp, err := http.Get(rest)
	if err != nil {
		panic(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		str := string(respBody)
		println(str)
	}
	m.OMetric = *parser.JsonUnmarshaller(respBody)
	defer resp.Body.Close()

	return parser.GetLast(&m.OMetric)
}

func(m *TSManager)SelectMetricWithCNPName(Cluster string, Node string, Pod string,MetricName string) {
	m.SetCNPName(Cluster,Node,Pod,output)
	m.SelectMetric(MetricName)
}