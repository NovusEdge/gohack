package gohack

import (
    helpers "gohack/lib/helpers"
    "time"
    "errors"
    "log"
)


//TTYClient: ...
/*
Fields:
    Client:     [&helpers.Client] ...
    Timeout:    [time.Duration] ...
*/
type TTYClient struct {
    Client     &helpers.Client
}


func (tc *TTYClient) Start() error {

}
