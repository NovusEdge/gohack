package gohack

type MappedFlag struct {
	Name     string
	Value    interface{}
	Assigned bool
}

var URL, Protocol, Timeout, Start, End, Port MappedFlag

func (mp *MappedFlag) Bind(name string) {
	mp.Name = name
	mp.Assigned = true
}

func (mp *MappedFlag) Release() {
	mp.Name = ""
	mp.Assigned = false
}

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

func MakeArgMap() {

}

// var URL, Protocol string
// var Timeout time.Duration
// var Start, End, Port int
