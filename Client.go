package tsdbmanager

import parser "github.com/hth0919/tsdbparser"





func(m *TSManager) NewTSManager(host string) *TSManager{
	m.Client = m.Client.NewTSClient(host)
	m.IMetric = IMetric{
		Metric:    "",
		Timestamp: 0,
		Value:     0,
		Tags: map[string]string{
			"Cluster" : "",
			"Node" : "",
			"Pod" : "",
		},
	}
	m.OMetric = parser.Metric{
		MetricName: "",
		Tags: map[string]string{},
		Dps: map[int64]float64{},
	}
	return m
}

func(m *TSManager)SetCNPName(Cluster string, Node string, Pod string, Mode IOType) {
	if Mode == input {
		m.IMetric.Tags["Cluster"] = Cluster
		m.IMetric.Tags["Node"] = Node
		m.IMetric.Tags["Pod"] = Pod
	} else if Mode == output {
		m.Client.Cluster = Cluster
		m.Client.Node = Node
		m.Client.Pod = Pod
	} else {
		panic("Invalid Mode")
	}
}

func(c *TSClient) NewTSClient(host string) *TSClient{
	client := &TSClient{
		Host:    host,
		Cluster: "",
		Node:    "",
		Pod:     "",
	}

	return client
}