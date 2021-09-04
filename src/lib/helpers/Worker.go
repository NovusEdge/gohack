package gohack

type Worker struct {
    Data chan struct{}
    Quit chan bool
    isStopped bool
}
