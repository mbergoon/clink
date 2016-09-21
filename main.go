package main

import "fmt"

// configure initializes loggers and application arguments.
func configure(c *ClinkConfig) {
	c.HandleFlags()
	InitLoggers(c)
}

func appIntro() {
	//Display message with mode
	//Log beginning of applciation
	LogM(InfoLevel, "Hello from log")
}

func main() {

	cconf := NewClinkConfig()
	configure(cconf)

	for i := 0; i < 100000; i++ {
		appIntro()
	}

	fmt.Println(*cconf)

}
