package gohack

import (
	"fmt"
	"os/exec"
)

//BannerGrabber : A banner grabber
/*

 */
type BannerGrabber struct {
	URL     string
	Command string
}

//Grab ...
func (bg *BannerGrabber) Grab() {

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

	curlCommand := exec.Command(cURL, []string{"-s", "-I", address})
	wgetCommand := exec.Command(wGET, []string{"-q", "-S", address})

	if flag1 != nil && flag2 != nil {
		fmt.Printf("%s[*] Couldn't find either of cURL or wget, please install them to use the banner grabber...%s\n", ColorYellow, ColorReset)
	} else {
		if flag1 == nil {
			curlCommand.Run()
		} else if flag2 == nil {
			wgetCommand.Run()
		}
	}
}

func _curlGrab(address string) {
	cURL, flag1 := exec.LookPath("curl")

	if flag1 != nil {
		fmt.Printf("%s[-] cURL not found in PATH ...%s\n", ColorRed, ColorReset)

	} else {
		curlCommand := exec.Command(cURL, []string{"-s", "-I", address})
		curlCommand.Run()
	}

}
