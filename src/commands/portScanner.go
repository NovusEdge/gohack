package main

import (
	"flag"
	"time"

	gops "github.com/NovusEdge/go-ps"
)

func main() {
	var lower, upper int
	var domain, protocol string
	var timeout time.Duration

	flag.IntVar(&lower, "start", 1, "The port to start scanning.")

	flag.IntVar(&upper, "end", 1024, "The port to end scanning.")

	flag.StringVar(&domain, "domain", "scanme.nmap.org", "Domain for the target to scan.")

	flag.DurationVar(&timeout, "timeout", 500, "Timeout for each port scan")

	flag.StringVar(&protocol, "protocol", "tcp", "Protocol to scan the port on. (tcp/udp)")

	flag.Parse()

	ps := gops.PortScanner{
		Domain:   domain,
		Protocol: protocol,
	}

	ps.Scan(lower, upper, timeout*time.Millisecond)
}
