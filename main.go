package main

import "fmt"

// configure initializes loggers and application arguments.
func configure(c *ClinkConfig) {
	c.HandleFlags()
	InitLoggers(c)
}

func appIntro() {
	LogM(InfoLevel, "Hello from info log")
	LogM(TraceLevel, "Hello from trace log")
	LogM(WarningLevel, "Hello from warn log")
	LogM(ErrorLevel, "Hello from error log")
}

func main() {

	cconf := NewClinkConfig()
	configure(cconf)

	// fmt.Println(CLR_0 + "HELLO" + CLR_N + CLR_R + "HELLO" + CLR_N + CLR_G + "HELLO" + CLR_N + CLR_Y + "HELLO" + CLR_N + CLR_B + "HELLO" + CLR_N + CLR_M + "HELLO" + CLR_N + CLR_C + "HELLO" + CLR_N + CLR_W + "HELLO" + CLR_N)

	// fmt.Println("\x1b[31;1m hello \x1b[0m")

	for i := 0; i < 100; i++ {
		appIntro()
	}

	fmt.Println(*cconf)

}
