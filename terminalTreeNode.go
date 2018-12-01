package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
)

type terminalTreeNode struct {
	sc       *sshConfig
	nodeID   treePositionID //local node ID, same as sc.TreePositionID
	children []*terminalTreeNode
}

//return true if treeNode is ancestor of node or treeNode is same position as node
func (treeNode *terminalTreeNode) isAncestor(node *terminalTreeNode) bool {
	return strings.HasPrefix(node.nodeID.singleString, treeNode.nodeID.singleString)
}

func (treeNode *terminalTreeNode) isSamePosition(node *terminalTreeNode) bool {
	return treeNode.nodeID.singleString == node.nodeID.singleString
}

func (treeNode *terminalTreeNode) searchParent(node *terminalTreeNode) *terminalTreeNode {
	if node == nil {
		return nil
	}
	for _, child := range treeNode.children {
		if child.isAncestor(node) {
			return child.searchParent(node)
		}
	}
	return treeNode
}

func (treeNode *terminalTreeNode) searchNode(nodePosition string) *terminalTreeNode {
	if treeNode.nodeID.singleString == nodePosition {
		return treeNode
	}
	for _, child := range treeNode.children {
		if child.searchNode(nodePosition) != nil {
			return child
		}
	}
	return nil
}

func newTerminalTreeNode(sshConf *sshConfig) *terminalTreeNode {
	node := terminalTreeNode{sc: sshConf, nodeID: sshConf.TreePositionID, children: make([]*terminalTreeNode, 0)}
	return &node
}

func (treeNode *terminalTreeNode) String() string {
	if treeNode == nil {
		return "\n"
	}
	var sb strings.Builder
	sb.WriteString(treeNode.nodeID.singleString)
	sb.WriteString(": ")
	if treeNode.sc == nil {
		sb.WriteString("nil\n")
	} else {
		sb.WriteString(treeNode.sc.Title)
		sb.WriteString("\n")
	}
	return sb.String()
}

func (treeNode *terminalTreeNode) TreeString() string {
	var sb strings.Builder
	treeNode.treeStringHelper(&sb, 0)
	return sb.String()
}

//recusive function that iterates through the tree with inorder traversal
func (treeNode *terminalTreeNode) treeStringHelper(sb *strings.Builder, level int) {
	if level > 0 {
		for i := 0; i < level; i++ {
			sb.WriteString("  ")
		}
		sb.WriteString("|-")
	}
	sb.WriteString(treeNode.String())
	if treeNode.children != nil {
		for _, node := range treeNode.children {
			node.treeStringHelper(sb, level+1)
		}
	}
}

func (treeNode *terminalTreeNode) inOrderIteration(operation terminalTreeNodeOperation) error {
	err := operation.nodeOperation(treeNode)
	if err != nil {
		err := fmt.Errorf("Operation failed on node: %s", treeNode.String())
		return err
	}
	if treeNode.children != nil {
		for _, child := range treeNode.children {
			err := child.inOrderIteration(operation)
			if err != nil {
				return nil
			}
		}
	}
	return nil
}

func (treeNode *terminalTreeNode) toSlice() []*terminalTreeNode {
	result := make([]*terminalTreeNode, 0)

	result = append(result, treeNode)

	for i := range treeNode.children {
		child := treeNode.children[i]
		result = append(result, child.toSlice()...)
	}

	return result
}

func publicKeyFile(file string) ssh.AuthMethod {
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

//convert current node's configuration to a ssh client config
func (treeNode *terminalTreeNode) toSSHClientConfig() *ssh.ClientConfig {
	if treeNode.sc == nil {
		return nil
	}
	var result *ssh.ClientConfig
	switch treeNode.sc.AuthMethod {
	case Password:
		result = &ssh.ClientConfig{
			User: treeNode.sc.UserName,
			Auth: []ssh.AuthMethod{ssh.Password(treeNode.sc.AuthPhrase)},
		}
	case PrivateKey:
		result = &ssh.ClientConfig{
			User: treeNode.sc.UserName,
			Auth: []ssh.AuthMethod{publicKeyFile(treeNode.sc.AuthPhrase)},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
		}
	default:
		return nil
	}

	return result
}
