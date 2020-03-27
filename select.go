package tsdbmanager

import (
	parser "github.com/hth0919/tsdbparser"
	"io/ioutil"
	"net/http"

)

func(m *TSManager)GetMetric(MetricName string) interface{}{
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
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		str := string(respBody)
		println(str)
	}
	m.OMetric = *parser.JsonUnmarshaller(respBody)
	defer resp.Body.Close()

	return parser.GetLast(&m.OMetric)
}

func(m *TSManager)GetNodeZone() interface{} {
	query :="/api/query?"
	start := "start=1h-ago"
	metric := "m=sum:"+"default"
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

	return m.OMetric.Tags["Zone"]
}

func(m *TSManager)GetNodeRegion() interface{} {
	query :="/api/query?"
	start := "start=1h-ago"
	metric := "m=sum:"+"default"
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

	return m.OMetric.Tags["Region"]
}

func(m *TSManager)GetMetricWithCNPName(Cluster string, Node string, Pod string,MetricName string) interface{}{
	m.SetCNPName(Cluster,Node,Pod,output)
	return m.GetMetric(MetricName)
}
func(m *TSManager)GetNodeRegionWithCNPName(Cluster string, Node string, Pod string) interface{}{
	m.SetCNPName(Cluster,Node,Pod,output)
	return m.GetNodeRegion()
}
func(m *TSManager)GetNodeZoneWithCNPName(Cluster string, Node string, Pod string) interface{}{
	m.SetCNPName(Cluster,Node,Pod,output)
	return m.GetNodeZone()
}
