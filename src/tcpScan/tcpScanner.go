package gohack

import (
	"fmt"
	colors "gohack/src"
	"net"
	"sync"
	"time"
)

//ScanTCP ...
func ScanTCP() {
	var wg sync.WaitGroup

	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%s[+] Port %d is Open.\n", colors.ColorGreen, j)
		}(i)
	}
	wg.Wait()
}
