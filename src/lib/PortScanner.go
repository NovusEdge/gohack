package gohack

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//PortScanner : A port scanner
/*

 */
type PortScanner struct {
	Domain   string
	Protocol string
}

//Scan ...
func (ps *PortScanner) Scan(lower int, upper int, timeout time.Duration) {
	var wg sync.WaitGroup

	for i := lower; i <= upper; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", ps.Domain, j)
			conn, err := net.DialTimeout(ps.Protocol, address, timeout)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%s[+] Port %d is Open.%s\n", ColorGreen, j, ColorReset)
		}(i)
	}
	wg.Wait()
}
