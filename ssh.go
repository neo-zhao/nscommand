package main

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func runSSHCommand(node *terminalTreeNode, command *string) (string, error) {
	clientConfig := node.toSSHClientConfig()
	if clientConfig == nil {
		return "", fmt.Errorf("Can not get ssh client config from node:%s", node.sc.TreePosition)
	}
	connection, err := ssh.Dial("tcp", node.sc.HostName+":"+node.sc.Port, clientConfig)
	if err != nil {
		return "", fmt.Errorf("Can not connect to host for node:%s Error:%s", node.sc.TreePosition, err)
	}
	defer connection.Close()
	session, err := connection.NewSession()
	if err != nil {
		return "", fmt.Errorf("Session failed for node:%s Error:%s", node.sc.TreePosition, err)
	}
	defer session.Close()

	var outBuffer, errorBuffer bytes.Buffer
	session.Stdout = &outBuffer
	session.Stderr = &errorBuffer

	err = session.Run(*command)
	if err != nil {
		return "", fmt.Errorf("Command failed for node:%s  Error:%s", node.sc.TreePosition, err)
	}

	return outBuffer.String(), nil
}
