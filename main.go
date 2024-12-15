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
/* raw sockets el priv, for interaction on network stack -> maybe*/
// log on measure of rtt on travel to (destination -> back to you)
// packet loss if no response on (recieve) withing _timeout_ hence=loss
/* packe loss log on tracking to diagnose network realibity*/
/*
intermittent network, focus on later after the concepts are ready
*/

/* main */
func main() {
	for i, arg := range os.Args {
		fmt.Printf("os.args[%d] -> %s\n", i, arg)
	}
}
