package tsdbParser

import (
	"encoding/json"
	"strconv"
)

func JsonUnmarshaller (jsonfile []byte) *Metric {

	c:= make([]Metric,0,0)
	err :=json.Unmarshal(jsonfile, &c)
	if err != nil {
		panic(err)
	}

	return &c[0]
}

func GetLast(input *Metric) interface{}  {
	var last int
	last = 0
	for k := range input.Dps {
		i,err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		if i>last {
			last = i
		}
	}

	l := strconv.Itoa(last)

	return input.Dps[l]
}