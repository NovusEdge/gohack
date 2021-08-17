package main

import (
    h "gohack/lib/helpers"
    ssh "golang.org/x/crypto/ssh"
)


func main() {
    conf := ssh.ClientConfig{
        User: "bandit0",
        Auth: []ssh.AuthMethod{
            ssh.Password("bandit0"),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    c, err := h.MakeClient(&conf, nil, false, true)
    if err != nil {
        panic(err)
    }

    tc := h.TTYClient{
        Client: c,
    }

    tc.Start("tcp", "bandit.labs.overthewire.org", 2220)
}
