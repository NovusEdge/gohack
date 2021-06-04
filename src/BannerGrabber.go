package gohack

import (
	"fmt"
	"os/exec"
)

//BannerGrabber : A banner grabber
/*

 */
type BannerGrabber struct {
	URL string
}

//Grab ...
func (bg *BannerGrabber) Grab() {
	fmt.Printf("%s[*] Grabbing banner for: %s...\n", ColorYellow, bg.URL)

	_blindGrabHTTP(bg.URL)
}

func _blindGrabHTTP(address string) {
	cURL, flag1 := exec.LookPath("curl")
	wGET, flag2 := exec.LookPath("wget")

	if flag1 != nil {
		fmt.Printf("%s[-] cURL not found in PATH ...%s\n", ColorRed, ColorReset)

	}

	if flag2 != nil {
		fmt.Printf("%s[-] wget not found in PATH ...%s\n", ColorRed, ColorReset)

	}

	if flag1 != nil && flag2 != nil {
		fmt.Printf("%s[*] Couldn't find either of cURL or wget, please install them to use the banner grabber...%s\n", ColorYellow, ColorReset)
	} else {
		if flag1 == nil {
			_curlGrab(cURL, address)
		} else if flag2 == nil {
			_wgetGrab(wGET, address)
		}
	}
}

func _curlGrab(cURL string, address string) {
	fmt.Printf("%s[*] Using cURL to perform the grab.\n", ColorYellow)

	curlCommand := exec.Command(cURL, "-s", "-I", address)
	curlCommand.Run()
}

func _wgetGrab(wGET string, address string) {
	fmt.Printf("%s[*] Using wGET to perform the grab.\n", ColorYellow)

	wgetCommand := exec.Command(wGET, "-q", "-S", address)
	wgetCommand.Run()
}
