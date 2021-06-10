package main

import (
	"flag"
	src "gohack/src"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "scanme.nmap.org", "URL for the target to scan.")

	flag.Parse()

	bg := src.BannerGrabber{URL: url}
	bg.Grab()
}
