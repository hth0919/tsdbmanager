package tsdbmanager

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func(m *TSManager)InsertMetric(MetricName string, Value interface{}) {
	query := "/api/put?details"
	rest := m.Client.Host + query
	m.IMetric.Metric = MetricName
	m.IMetric.Value = Value
	m.IMetric.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)

	inputdata, err := json.Marshal(m.IMetric)
	if err != nil {
		panic(err)
	}
	buff := bytes.NewBuffer(inputdata)
	resp, err := http.Post(rest,"application/json", buff)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}

func(m *TSManager)InsertMetricWithCNPName(Cluster string, Node string, Pod string, MetricName string, Value interface{}) {
	m.SetCNPName(Cluster,Node,Pod, input)
	m.InsertMetric(MetricName,Value)
}
