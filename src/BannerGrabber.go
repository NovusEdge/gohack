package gohack

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

//BannerGrabber : A banner grabber
/*

 */
type BannerGrabber struct {
	URL string
}

//Grab ...
func (bg *BannerGrabber) Grab(lower int, upper int, timeout time.Duration) {
	fmt.Printf("%s[*] Grabbing banner for: %s...\n", ColorYellow, bg.URL)

	var wg sync.WaitGroup

	for i := lower; i <= upper; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", bg.URL, j)
			res, err := _blindGrabHTTP(bg.URL)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%s[+] Port %d is Open.\n", ColorGreen, j)
		}(i)
	}
	wg.Wait()

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
}

func _curlGrab(cURL string, address string) (res string, err error) {
	fmt.Printf("%s[*] Using cURL to perform the grab.\n", ColorYellow)

	curlCommand := exec.Command(cURL, "-s", "-I", address)
	curlCommand.Run() //TODO: add a return statement
}

func _wgetGrab(wGET string, address string) (res string, err error) {
	fmt.Printf("%s[*] Using wGET to perform the grab.\n", ColorYellow)

	wgetCommand := exec.Command(wGET, "-q", "-S", address)
	wgetCommand.Run() // TODO: add a return statement
}
