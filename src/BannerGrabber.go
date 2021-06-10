package gohack

import (
	"bytes"
	"errors"
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

	res, err := _blindGrabHTTP(bg.URL)

	if err != nil {
		fmt.Printf("%s[-] E: %s\n", ColorRed, err)
	} else {
		fmt.Printf("%s[*] Banner Grab Successful!%s\nBanner:\n%s%s", ColorYellow, ColorReset, ColorGreen, res)
	}

}

func _blindGrabHTTP(address string) (res string, err error) {
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
			return _curlGrab(cURL, address)
		} else if flag2 == nil {
			return _wgetGrab(wGET, address)
		}
	}
	return "", nil
}

func _curlGrab(cURL string, address string) (res string, err error) {
	fmt.Printf("%s[*] Using cURL to perform the grab.\n", ColorYellow)

	curlCommand := exec.Command(cURL, "-s", "-I", address)

	stdout, _ := curlCommand.StdoutPipe()
	stderr, _ := curlCommand.StderrPipe()

	curlCommand.Start()

	outBuf, errBuf := new(bytes.Buffer), new(bytes.Buffer)
	outBuf.ReadFrom(stdout)
	errBuf.ReadFrom(stderr)

	res = outBuf.String()
	errstr := errBuf.String()

	if errstr != "" {
		return "", errors.New(errstr)
	}

	return res, nil
}

func _wgetGrab(wGET string, address string) (res string, err error) {
	fmt.Printf("%s[*] Using wGET to perform the grab.\n", ColorYellow)

	wgetCommand := exec.Command(wGET, "-q", "-S", address)

	stdout, _ := wgetCommand.StdoutPipe()
	stderr, _ := wgetCommand.StderrPipe()

	wgetCommand.Start()

	outBuf, errBuf := new(bytes.Buffer), new(bytes.Buffer)
	outBuf.ReadFrom(stdout)
	errBuf.ReadFrom(stderr)

	res = outBuf.String()
	errstr := errBuf.String()

	if errstr != "" {
		return "", errors.New(errstr)
	}

	return res, nil
}
