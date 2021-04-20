package gohack

import (
	"net"
	"time"
)

//BannerGrabber : A banner grabber
/*

 */
type BannerGrabber struct {
	URL      string
	Protocol string
}

//Grab ...
func (bg *BannerGrabber) Grab(protocol string, port int, timeout time.Duration) {
	d := net.Dialer{Timeout: timeout, LocalAddr: nil}

}
