package main

import (
	"flag"
	"fmt"
	"strings"
)

// Default Clink configuration values
const (
	DEFAULT_MODE        CmdMode  = 1
	DEFAULT_EXEC_MODE   ExecMode = 0
	DEFAULT_CONFIG_FILE string   = ""
	DEFAULT_HOST_URI    string   = ""
)

// CmdMode describes the current mode Clink is running in. Describing the actions
// which are available to take.
type CmdMode int

// CmdMode definition constants.
const (
	MNGE CmdMode = iota
	EXEC
	RPRT
	UNDF
)

// ExecMode describes the type of operation to run on the execution. Describes the
// type of scan to run.
type ExecMode int

// ExecMode definition constants.
const (
	ECHOMON ExecMode = iota
	PSCNMON
	UNDFMON
)

// ClinkConfig is the configuration for the Clink command. Detailing how the command will
// execute.
type ClinkConfig struct {
	cmdMode         CmdMode
	execMode        ExecMode
	configFile      string
	logLevel        string
	logMode         string
	logDestination  [4]string
	interval        int
	intervalFail    int
	intervalSuccess int
	hosts           []string
	internalHosts   []string
	timeout         int
	distributed     bool
	external        bool
	internalNodes   []string
	externalNodes   []string
	duration        int
}

// NewClinkConfig creates a new clink configuration. Note: the values take the zero value
// rather than the default values defined above.
func NewClinkConfig() *ClinkConfig {
	return &ClinkConfig{}
}

// commandUsage provides a custom usage prompt to overcome readablilty challenges
// presented by default flags implementation.
func commandUsage() {
	fmt.Printf("Usage: clink -[m|e|r[options]] -[echo|scan[options]] -[host|file]>\n")
	fmt.Printf("	Command Mode\n")
	fmt.Printf("		-m 			Management mode\n")
	fmt.Printf("		-e 			Execute mode\n")
	fmt.Printf("		-r 			Report mode\n")
	fmt.Printf("	Execute Mode\n")
	fmt.Printf("		-echo 			Execute icmp echo monitoring\n")
	fmt.Printf("		-pscan			Execute port scan monitoring\n")
	fmt.Printf("	Configuration Description\n")
	fmt.Printf("		-file 			Configuration file to describe execution of monitoring task\n")
	fmt.Printf("	Configuration Options\n")
	fmt.Printf("		-interval		Interval (ms) for rescanning each host\n")
	fmt.Printf("		-interval-fail		Fail retry interval (ms) (overrides -interval) \n")
	fmt.Printf("		-interval-success	Success retry interval (ms) (overrides -interval) \n")
	fmt.Printf("		-host 			Configure clink to launch scan on defined host(s)\n")
	fmt.Printf("		-internal-hosts		List (comma delimited) of hosts to only scan from internal if -distributed\n")
	fmt.Printf("		-timeout		Time (ms) to wait before closing connection and failing \n")
	fmt.Printf("		-distributed		Launch scan on defined external hosts concurrently\n")
	fmt.Printf("		-external		Sets current (local) monitor context to external\n")
	fmt.Printf("		-internal-nodes		List of internal scanning nodes\n")
	fmt.Printf("		-external-nodes		List of external scanning nodes\n")
	fmt.Printf("		-duration		Time (ms) to run monitor\n")
	fmt.Printf("	Logging Level\n")
	fmt.Printf("		-log-level		Set level of logging to include options: [TRACE|INFO|WARNING|ERROR] default: INFO\n")
	fmt.Printf("	Logging Destination\n")
	fmt.Printf("		-log-trace-dest		Sets where to write trace log options: [<filename>|DISCARD|STDOUT|STDERR]\n")
	fmt.Printf("		-log-info-dest		Sets where to write info log options: [<filename>|DISCARD|STDOUT|STDERR]\n")
	fmt.Printf("		-log-warning-dest	Sets where to write warning log options: [<filename>|DISCARD|STDOUT|STDERR]\n")
	fmt.Printf("		-log-error-dest		Sets where to write error log options: [<filename>|DISCARD|STDOUT|STDERR]\n")
}

// HandleFlags initializes the potential settings and parses their value into the config
// object. Care should be taken to update `commandUsage` function to reflect any changes
// to this function.
func (c *ClinkConfig) HandleFlags() {
	//Set Default Usage
	flag.Usage = commandUsage

	//Command modes
	var m = flag.Bool("m", false, "command mode 'manage'")
	var e = flag.Bool("e", false, "command mode 'execute'")
	var r = flag.Bool("r", false, "command mode 'report'")

	//Execute modes
	var echo = flag.Bool("echo", false, "execute mode 'icmp echo monitor'")
	var pscan = flag.Bool("pscan", false, "execute mode 'port scan monitor'")

	//File
	flag.StringVar(&(c.configFile), "file", "", "configuration file describing scan")

	//Log Level
	var logLevel = flag.String("log-level", "INFO", "minimum log level to write")

	//Log Destination
	var logTraceDest = flag.String("log-trace-dest", "STDOUT", "stream or file to write on")
	var logInfoDest = flag.String("log-info-dest", "STDOUT", "stream or file to write on")
	var logWarningDest = flag.String("log-warning-dest", "STDERR", "stream or file to write on")
	var logErrorDest = flag.String("log-error-dest", "STDERR", "stream or file to write on")

	//Configuration Options
	var interval = flag.Int("interval", 1000, "")
	var intervalFail = flag.Int("interval-fail", 1000, "")
	var intervalSuccess = flag.Int("--success", 1000, "")
	var hosts = flag.String("hosts", "", "hosts to scan")

	var internalHosts = flag.String("internal-hosts", "", "")
	var timeout = flag.Int("timeout", 5000, "")
	var distributed = flag.Bool("distributed", false, "")
	var external = flag.Bool("external", false, "")
	var internalNodes = flag.String("internal-nodes", "", "")
	var externalNodes = flag.String("external-nodes", "", "")
	var duration = flag.Int("duration", 10000, "")

	flag.Parse()

	var tailHosts = flag.Args()

	//Build config
	switch {
	case *m:
		c.cmdMode = MNGE
	case *e:
		c.cmdMode = EXEC
	case *r:
		c.cmdMode = RPRT
	default:
		c.cmdMode = UNDF
	}

	switch {
	case *echo:
		c.execMode = ECHOMON
	case *pscan:
		c.execMode = PSCNMON
	default:
		c.execMode = UNDFMON
	}

	c.interval = *interval
	c.intervalFail = *intervalFail
	c.intervalSuccess = *intervalSuccess
	c.timeout = *timeout
	c.duration = *duration

	c.distributed = *distributed
	c.external = *external

	c.hosts = strings.Split(*hosts, ",")
	c.hosts = append(c.hosts, tailHosts...)
	c.internalHosts = strings.Split(*internalHosts, ",")
	c.internalNodes = strings.Split(*internalNodes, ",")
	c.externalNodes = strings.Split(*externalNodes, ",")

	//Build log config
	c.logLevel = *logLevel

	var logDests = [4]string{"DISCARD", "STDOUT", "STDOUT", "STDERR"}
	logDests[0] = *logTraceDest
	logDests[1] = *logInfoDest
	logDests[2] = *logWarningDest
	logDests[3] = *logErrorDest
	c.logDestination = logDests

}

// MonitorFromClinkConfig generates a basic monitor as described by ClinkConfig
// making use of defaults, and other assumptions to run from CLI without specifying
// a config file.
func (c *ClinkConfig) MonitorFromClinkConfig() (MonitorConfig, error) {
	m := &MonitorConfig{
		Id:               0,
		Name:             "CLI-CONFIG",
		Probes:           *new([]Probe),
		Timeout:          c.timeout,
		Stype:            "CLI",
		ScanFrequency:    c.interval,
		Distributed:      c.distributed,
		DistributedNodes: *new([]Node),
		Internal:         !c.external,
		Interval:         c.interval,
		IntervalFail:     c.intervalFail,
		IntervalSuccess:  c.intervalSuccess,
	}

	for _, n := range c.internalNodes {
		if n != "" {
			node := Node{
				Name:     n,
				Host:     n,
				Internal: true,
			}
			m.DistributedNodes = append(m.DistributedNodes, node)
		}
	}

	for _, n := range c.externalNodes {
		if n != "" {
			node := Node{
				Name:     n,
				Host:     n,
				Internal: true,
			}
			m.DistributedNodes = append(m.DistributedNodes, node)
		}
	}

	for _, h := range c.internalHosts {
		if h != "" {
			probe := Probe{
				Name:                 h,
				Host:                 h,
				Internal:             true,
				ScanFromInternalOnly: true,
			}
			m.Probes = append(m.Probes, probe)
		}
	}

	for _, h := range c.hosts {
		if h != "" {
			probe := Probe{
				Name:                 h,
				Host:                 h,
				Internal:             false,
				ScanFromInternalOnly: false,
			}
			m.Probes = append(m.Probes, probe)
		}
	}

	return *m, nil
}

// Process to follow...
func (c *ClinkConfig) Process() {
	return
}
