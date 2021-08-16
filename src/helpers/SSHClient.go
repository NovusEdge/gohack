package gohack

import (
    ssh "golang.org/x/crypto/ssh"
    "errors"
    "log"
    "io"
    "io/ioutil"
)


//Client:
/*
Config [&ssh.ClientConfig]: ...

Client [&ssh.Client]: ...

Session: [&ssh.Session]: ...

AuthByCreds [bool]: ...

AuthByKey [bool]: ...
*/
type Client struct {
    Config *ssh.ClientConfig
    Client *ssh.Client
    Session *ssh.Session
    AuthByCreds bool
    AuthByKey bool
}


//MakeClient: ...
/*
params:
    config      [&ssh.ClientConfig]
    client      [&ssh.Client]
    authByCreds [bool]
    authByKey   [bool]
*/
func MakeClient(config *ssh.ClientConfig, client *ssh.Client, authByCreds bool, authByKey bool) (*Client, error){
    // Use XOR logic to branch the flow-control
    if !(authByCreds != authByKey) {
        log.Fatal("Invalid options for client-authentication")
        return nil, errors.New("E: Please choose either one of the authorization methods.")
    }

    c := Client{
        Config: config,
        Client: client,
        Session: nil,
        AuthByKey: authByKey,
        AuthByCreds: authByCreds,
    }

    return &c, nil
}


//SetConfig: ...
/*
params:
    customConfig: [&ssh.ClientConfig]
*/
func (c *Client) SetConfig(config *ssh.ClientConfig) {
    c.Config = config
}


//RunCommand: ...
/*
params:
    command: [string]
*/
func (c *Client) RunCommand(command string) (string, string, error) {
    if c.Session == nil {
        return "", "", errors.New("E: Couldn't run the command since no session was started.")
    }

    // For stdout
    sessStdOut, stdoutErr := c.Session.StdoutPipe()
    if stdoutErr != nil {
        return "", "", stdoutErr
    }

    // For stderr
    sessStderr, stderrErr := c.Session.StderrPipe()
    if stderrErr != nil {
        return "", "", stderrErr
    }

    // Executing Command
    err := c.Session.Run(command)
    if err != nil {
        return "", "", err
    }

    // Reading the errors for the command:
    errOut, ioerrErr := io.ReadAll(sessStderr)
    if ioerrErr != nil {
        return "", "", ioerrErr
    }

    // Reading the output for the command:
    output, ioerrOut := io.ReadAll(sessStdOut)
    if ioerrOut != nil {
        return "", "", ioerrOut
    }

    return string(output), string(errOut), nil
}


//StartSession: ...
/*
params:
    network: [string]
    host:    [string]
*/
func (c *Client) StartSession(network string, host string) {
    client, err := _auth(c.Config, network, host)

    if err != nil {
        log.Fatal(err)
    } else {
        c.Client = client
        session, err := client.NewSession()

        if err != nil {
    		c.Client.Close()
    		log.Fatal(err)
    	} else {
            c.Session = session
        }
    }
}

//CloseSession
func (c *Client) CloseSession() error {
    if c.Session != nil {
        c.Session.Close()
        c.Session = nil
        return nil

    } else {
        return errors.New("E: Close Session error 10101") //TODO: Add a proper error message for this one
    }
}


func _auth(sshConfig *ssh.ClientConfig, network string, host string) (*ssh.Client, error) {

    sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

    client, err := ssh.Dial(network, host, sshConfig)
	if err != nil {
		return nil, err
	}

    return client, nil
}


func _publicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}
