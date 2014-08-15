package main

import (
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	"log"
	"bytes"
)

func main() {
	const (
		User     = "vagrant"
		Password = "vagrant"
		Uri      = "192.168.33.10:22"
		Cmd      = "ip a"
	)

	config := &ssh.ClientConfig{
		User: User,
		Auth: []ssh.AuthMethod{
			ssh.Password(Password),
		},
	}
	conn, err := ssh.Dial("tcp", Uri, config)
	if err != nil {
		log.Fatalf("unable to connect: %s", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("unable to create session: %s", err)
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	err = session.Run(Cmd)
	if err != nil {
		log.Fatalf("execute command is failure: %s", err)
	}

	fmt.Println(stdoutBuf.String())
}

