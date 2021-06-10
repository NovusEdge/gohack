package main

import (
	"flag"
	p "gohack/src"
	"time"
)

func main() {
	var lower, upper int
	var url, protocol string
	var timeout time.Duration

	flag.IntVar(&lower, "start", 1, "The port to start scanning.")

	flag.IntVar(&upper, "end", 1024, "The port to end scanning.")

	flag.StringVar(&url, "url", "scanme.nmap.org", "URL for the target to scan.")

	flag.DurationVar(&timeout, "timeout", 500, "Timeout for each port scan")

	flag.StringVar(&protocol, "protocol", "tcp", "Protocol to scan the port on. (tcp/udp)")

	flag.Parse()

	ps := p.PortScanner{
		URL:      url,
		Protocol: protocol,
	}

	ps.Scan(lower, upper, timeout*time.Millisecond)
}
