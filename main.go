package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/mbergoon/clink/icmpecho"
)

// configure initializes loggers and application arguments.
func configure(c *ClinkConfig) MonitorConfig {
	c.HandleFlags()
	InitLoggers(c)

	appInit()

	LogM(TraceLevel, "Built ClinkConfig with - "+fmt.Sprint(c))

	if c.configFile != "" {
		LogM(InfoLevel, "Attempting to read specified clink MonitorConfig")

		data, err := ioutil.ReadFile(c.configFile)
		if err != nil {
			LogM(ErrorLevel, "Failed to open/read the file - "+c.configFile)
			os.Exit(-1)
		}
		m, err := ParseMonitorConfig(bytes.NewReader(data))
		if err != nil {
			LogM(ErrorLevel, "Failed to parse configuration file")
			os.Exit(-1)
		}
		LogM(InfoLevel, "Read specified clink MonitorConfig successfully")
		LogM(TraceLevel, "Parsed MonitorConfiguration with - "+fmt.Sprint(m))
		return m
	} else {
		LogM(InfoLevel, "Clink MonitorConfig not specified. Attempting to build config from flags.")
		//Build config from options
		m, err := c.MonitorFromClinkConfig()
		if err != nil {
			LogM(ErrorLevel, "Failed to build monitor from CLI flags.")
			os.Exit(-1)
		}
		LogM(InfoLevel, "Built config from flags successfully")
		LogM(TraceLevel, "Built MonitorConfig from flags with - "+fmt.Sprint(m))
		return m
	}
}

// appInit handle boiler plate initialization. Called after flags are parse, and loggers
// are initialized but before configuration is constructed.
func appInit() {
	LogM(InfoLevel, "Initializing clink - building monitor based on settings.")
}

// main entry point for clink
func main() {
	cconf := NewClinkConfig()
	m := configure(cconf)

	switch {
	case cconf.cmdMode == MNGE:
		LogM(TraceLevel, "Command mode manage (MNGE) set")

	case cconf.cmdMode == RPRT:
		LogM(TraceLevel, "Command mode report (RPRT) set")

	case cconf.cmdMode == EXEC:
		LogM(TraceLevel, "Command mode execute (EXEC) set")
		for {
			for _, probe := range m.Probes {
				LogM(InfoLevel, probe.Name+":"+probe.Host+" "+fmt.Sprint(icmpecho.Echo(probe.Host, 3000)))
			}
			time.Sleep(time.Duration(m.Interval) * time.Millisecond)
		}
	case cconf.cmdMode == UNDF:
		LogM(ErrorLevel, "Command mode not specfied or unrecognized - specify [-m (MNGE)|-r (RPRT)|-e (EXEC)]")
		os.Exit(-1)
	default:
		LogM(ErrorLevel, "Command mode not specfied or incorrect - specify [MNGE|RPRT|EXEC]")
		os.Exit(-1)
	}
}
