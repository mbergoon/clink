package main

import (
	"flag"
	"fmt"
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

// Default Clink configuration values
const (
	DEFAULT_MODE        CmdMode  = 1
	DEFAULT_EXEC_MODE   ExecMode = 0
	DEFAULT_CONFIG_FILE string   = "clink.conf"
	DEFAULT_HOST_URI    string   = ""
)

// ClinkConfig is the configuration for the Clink command. Detailing how the command will
// execute.
type ClinkConfig struct {
	cmdMode        CmdMode
	execMode       ExecMode
	configFile     string
	hostUri        string
	logLevel       string
	logMode        string
	logDestination [4]string
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
	fmt.Printf("	Input Description\n")
	fmt.Printf("		-file 			Configuration file to describe execution of monitoring task\n")
	fmt.Printf("		-host 			Configure clink to launch scan on defined host\n")
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

	//File or Host
	flag.StringVar(&(c.configFile), "file", "clink.conf", "configuration file describing scan")
	flag.StringVar(&(c.hostUri), "host", "127.0.0.1", "host to scan")

	//Log Level
	var logLevel = flag.String("log-level", "INFO", "minimum log level to write")

	//Log Destination
	var logTraceDest = flag.String("log-trace-dest", "STDOUT", "stream or file to write on")
	var logInfoDest = flag.String("log-info-dest", "STDOUT", "stream or file to write on")
	var logWarningDest = flag.String("log-warning-dest", "STDERR", "stream or file to write on")
	var logErrorDest = flag.String("log-error-dest", "STDERR", "stream or file to write on")

	flag.Parse()

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

	//Build log config
	c.logLevel = *logLevel

	var logDests = [4]string{"DISCARD", "STDOUT", "STDOUT", "STDERR"}
	logDests[0] = *logTraceDest
	logDests[1] = *logInfoDest
	logDests[2] = *logWarningDest
	logDests[3] = *logErrorDest
	c.logDestination = logDests
}

// Process to follow...
func (c *ClinkConfig) Process() {
	return
}
