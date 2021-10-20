package main

import (
	"flag"

	gobg "github.com/NovusEdge/go-bg"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "scanme.nmap.org", "URL for the target to scan.")

	flag.Parse()

	bg := gobg.BannerGrabber{URL: url}
	bg.Grab()
}
