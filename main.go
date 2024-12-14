package main

import (
	"fmt"
	"os"
)

/* main */

// have sockets -> x/net
// have icmp help -> x/net
// will need some time support so whichever eq to time.h from sys
// create the command argv -> os
func main() {
	for i, arg := range os.Args {
		fmt.Printf("os.args[%d] -> %s\n", i, arg)
	}
}
