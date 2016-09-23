package main

import (
	"encoding/json"
	"errors"
	"io"
)

/*{
	"Id": 100,
	"Name": "cameron",
	"Probes": [
		{"Name": "PName", "Host": "PHost", "ScanFromInternalOnly": true, "IsInternalResource": true},
		{"Name": "PName1", "Host": "PHost1", "ScanFromInternalOnly": true, "IsInternalResource": true}
	],
	"Duration": 99,
	"Stype": "ECHO",
	"Internal": false,
	"ScanFrequency": 10,
	"DistributedScan": false,
	"DistributedHosts": [
		"clink-01.com", "clink-02.com"
	]
}*/

type MonitorConfig struct {
	Id               int
	Name             string
	Probes           []Probe
	Timeout          int
	Stype            string
	ScanFrequency    int
	Distributed      bool
	DistributedNodes []Node
	Internal         bool
	Interval         int
	IntervalFail     int
	IntervalSuccess  int
	Duration         int
}

type Probe struct {
	Name                 string
	Host                 string
	ScanFromInternalOnly bool
	Internal             bool
}

type Node struct {
	Name     string
	Host     string
	Internal bool
}

func ParseMonitorConfig(jsonStream io.Reader) (MonitorConfig, error) {
	dec := json.NewDecoder(jsonStream)
	var m MonitorConfig
	for {
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			return m, errors.New("monitor-config: failed to parse configuration")
		}

	}
	return m, nil
}
