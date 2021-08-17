package gohack

import "fmt"


//TTYClient: ...
/*
Fields:
    Client:     [&Client] ...
*/
type TTYClient struct {
    Client *Client
}

// For the shell prompt...
type prompt struct {
    host    string
    uname   string
    pwd     string
    pstring string
}


//Start: ...
/*

*/
func (tc *TTYClient) Start(network string, host string, port int) error {
    return mainloop(*tc, network, host, port)
}


//Stop: ...
/*

*/
func (tc *TTYClient) Stop() error {
    return tc.Client.CloseSession()
}

func makePrompt(c Client) prompt {
    var host, uname, pwd string
    var hostErr, unameErr, pwdErr error

    var _prompt prompt

    host, _, hostErr = c.RunCommand("hostname")
    if hostErr != nil {
        return prompt{"", "", "", "$ "}
    }


    uname, _, unameErr = c.RunCommand("whoami")
    if unameErr != nil {
        return prompt{"", "", "", "$ "}
    }


    pwd, _, pwdErr = c.RunCommand("pwd")
    if pwdErr != nil {
        return prompt{"", "", "", "$ "}
    }
    _prompt.uname = uname
    _prompt.pwd   = pwd
    _prompt.host  = host

    _prompt.pstring = fmt.Sprintf("%s@%s | %s $ ", _prompt.uname, _prompt.host, _prompt.pwd)

    return _prompt
}


func mainloop(tc TTYClient, network string, host string, port int) error {
    var command string
    var serr, sout string
    var err error

    defer tc.Client.CloseSession()

    fmt.Println("To exit, type in \"gohack_exit\"")

    pmpt := makePrompt(*(tc.Client)).pstring
    for true {
        fmt.Print(pmpt)
        fmt.Scanf("%s", &command)

        if command == "gohack_exit" {
            fmt.Println("[*] Exiting tty shell...")
            break
        }
        tc.Client.StartSession(network, fmt.Sprintf("%s:%d", host, port))
        sout, serr, err = tc.Client.RunCommand(command)

        if err != nil {
            return err
        }

        if serr != "" {
            fmt.Println(serr)
        }

        fmt.Println(sout)
    }

    return err
}
