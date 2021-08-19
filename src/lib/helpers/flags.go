package gohack

import (
	"time"
	"fmt"
)

var URL StrFlag = StrFlag{}
var Protocol StrFlag = StrFlag{}
var Timeout DurationFlag = DurationFlag{}
var Start IntFlag = IntFlag{}
var End IntFlag = IntFlag{}
var Port IntFlag = IntFlag{}

var AssignedArray []interface{}

func BindAll() {
	URL.Bind("url")
	End.Bind("end")
	Start.Bind("start")
	Port.Bind("port")
	Protocol.Bind("protocol")
	Timeout.Bind("timeout")
}

func ReleaseAll() {
	URL.Release()
	End.Release()
	Start.Release()
	Port.Release()
	Protocol.Release()
	Timeout.Release()
}

func MakeArgMap() map[string]string {
	if checkForNone() {
		return map[string]string{"-h": ""}
	}
	res := make(map[string]string)

	if URL.Assigned { res[URL.Name] = URL.Value }
	if Protocol.Assigned { res[Protocol.Name] = Protocol.Value }
	if Start.Assigned { res[Start.Name] = fmt.Sprintf("%d", Start.Value) }
	if End.Assigned { res[End.Name] = fmt.Sprintf("%d", End.Value) }
	if Port.Assigned { res[Port.Name] = fmt.Sprintf("%d", Port.Value) }
	if Timeout.Assigned { res[Timeout.Name] = fmt.Sprintf("%d", Timeout.Value) }

	return res
}


func checkForNone() bool {
	return !(URL.Assigned || End.Assigned || Start.Assigned || Port.Assigned || Protocol.Assigned || Timeout.Assigned)
}

func (mp *IntFlag) Bind(name string) {
	mp.Name = name
	mp.Assigned = true
}

func (mp *IntFlag) Release() {
	mp.Name = ""
	mp.Assigned = false
}

func (mp *BoolFlag) Bind(name string) {
	mp.Name = name
	mp.Assigned = true
}

func (mp *BoolFlag) Release() {
	mp.Name = ""
	mp.Assigned = false
}

func (mp *StrFlag) Bind(name string) {
	mp.Name = name
	mp.Assigned = true
}

func (mp *StrFlag) Release() {
	mp.Name = ""
	mp.Assigned = false
}

func (mp *DurationFlag) Bind(name string) {
	mp.Name = name
	mp.Assigned = true
}

func (mp *DurationFlag) Release() {
	mp.Name = ""
	mp.Assigned = false
}

type BoolFlag struct {
	Name     string
	Value    bool
	Assigned bool
}

type DurationFlag struct {
	Name     string
	Value    time.Duration
	Assigned bool
}

type IntFlag struct {
	Name     string
	Value    int
	Assigned bool
}

type StrFlag struct {
	Name     string
	Value    string
	Assigned bool
}

// var URL, Protocol string
// var Timeout time.Duration
// var Start, End, Port int
