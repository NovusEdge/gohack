package gohack

import "time"

var URL StrFlag = StrFlag{}
var Protocol StrFlag = StrFlag{}
var Timeout DurationFlag = DurationFlag{}
var Start IntFlag = IntFlag{}
var End IntFlag = IntFlag{}
var Port IntFlag = IntFlag{}

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
    return map[string]string{}
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


// var URL, Protocol string
// var Timeout time.Duration
// var Start, End, Port int
