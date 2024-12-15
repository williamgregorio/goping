package main

import (
	"fmt"
	"os"
)

// have sockets -> x/net
// have icmp help -> x/net
// will need some time support so whichever eq to time.h from sys

// packet (header, identity->int(track), payload)
// raw sockets -> interaction with icmp (send/recieve)!no=tcp/udp

/* main */
func main() {
	for i, arg := range os.Args {
		fmt.Printf("os.args[%d] -> %s\n", i, arg)
	}
}
