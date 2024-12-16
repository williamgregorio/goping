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
/*
  rtt = time_reply_recieved - time_request_sent
*/
// packet loss if no response on (recieve) withing _timeout_ hence=loss
/* packet loss log on tracking to diagnose network realibity*/
/*
intermittent network, focus on later after the concepts are ready
*/

// tracerouting - ttl(increment ttl field in ip (packet(s)))
// the idea is a router along a path decrements ttl, on 0, router sends us back and icmp time exceeded msg hence maps the path from response
// destination reach => ^stops when target host responds with (icmp echo/reply)

/* main */
func main() {
	fmt.Println(len(os.Args))

	if len(os.Args) < 2 {
		fmt.Println("Usage: pingo <hostname> or <ip>")
		os.Exit(7)
	}
}
