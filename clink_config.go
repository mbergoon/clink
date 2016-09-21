package main

import (
	"flag"
	"fmt"
)

type CmdMode int

const (
	MNGE CmdMode = iota
	EXEC
	RPRT
	UNDF
)

type ExecMode int

const (
	ECHOMON ExecMode = iota
	PSCNMON
	UNDFMON
)

type ClinkConfig struct {
	cmdMode    CmdMode
	execMode   ExecMode
	configFile string
	hostUri    string
}

// Default Clink configuration values
const (
	DEFAULT_MODE      CmdMode  = 1
	DEFAULT_EXEC_MODE ExecMode = 0
)

func NewClinkConfig() *ClinkConfig {
	return &ClinkConfig{}
}

func commandUsage() {
	fmt.Printf("Usage: clink -[m|e|r[options]] -[echo|scan[options]] -[host|file]>\n")
	flag.PrintDefaults()
}

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

}

func (c *ClinkConfig) Process() {
	return
}
