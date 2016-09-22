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
	LogM(InfoLevel, "Hello from info log")
	LogM(TraceLevel, "Hello from trace log")
	LogM(WarningLevel, "Hello from warn log")
	LogM(ErrorLevel, "Hello from error log")
}

func main() {

	cconf := NewClinkConfig()
	configure(cconf)

	for i := 0; i < 100; i++ {
		appIntro()
	}

	fmt.Println(*cconf)

}
