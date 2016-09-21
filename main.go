package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// configure initializes loggers and application arguments.
func configure(c *ClinkConfig) {
	InitLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	c.HandleFlags()
}

func appIntro() {
	//Display message with mode
	//Log beginning of applciation
}

func main() {

	cconf := NewClinkConfig()
	configure(cconf)

	appIntro()

	fmt.Println(*cconf)

}
