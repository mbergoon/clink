package main

import (
	"encoding/json"
	"errors"
	"io"
)

/*{
	"Id": 100,
	"Name": "Monitor01",
	"Probes": [
		{"Name": "PName", "Host": "PHost", "ScanFromInternalOnly": true, "Internal": true},
		{"Name": "PName1", "Host": "PHost1", "ScanFromInternalOnly": true, "Internal": true}
	],
	"Timeout": 5000,
	"Stype": "ECHO",
	"ScanFrequency": 10,
	"Distributed": true,
	"DistributedNodes": [
		{"Name": "US01", "Host": "clink-01.co.us", "Internal": false},
		{"Name": "US02", "Host": "clink-02.co.us", "Internal": false},
		{"Name": "UK01", "Host": "clink-01.co.uk", "Internal": false}
	],
	"Internal": true,
	"Interval": 100,
	"Duration": 99
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
