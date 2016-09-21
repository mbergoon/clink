package main

import "fmt"

func configure(c *ClinkConfig) {
	c.HandleFlags()
}

func main() {
	cconf := NewClinkConfig()

	configure(cconf)

	fmt.Println(*cconf)

}
