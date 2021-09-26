package gohack

import (
	"time"
)

const (
	WindowsLocate string = `%tmp`
	UnixLoacte    string = "/usr/sbin"
)

type LogicBomb struct {
	BombID     string
	Deadline   *time.Time
	TargetFile string
}

func (lb *LogicBomb) Implant() {
	
}

func (lb *LogicBomb) Disarm() {

}

func Locate() {

}

func Remove(ID string) int {
	return 0
}

func RemoveAll() {

}
