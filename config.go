/*
sshConfig is used to unmarshal json-based configuration files.
for example:
{
        "treePosition": "1.2.3",
        "title": "host1",
        "hostName": "host1.mycompany.com", //could also be ip address "1.2.3.4"
		"disabled": true
        "userName": "ec2_user",
        "authMethod": 1,
        "authPhrase": "example.pem" //depends on the "authMethod", it could be the key file or password
}
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

//AuthenticationMethod enum for authentication
type AuthenticationMethod int

const (
	//Password uses password for authentication
	Password AuthenticationMethod = 0
	//PrivateKey uses private key for authentication
	PrivateKey AuthenticationMethod = 1
)

type sshConfig struct {
	TreePosition   string               `json:"treePosition"` //denotes the position in the terminal tree; uses '.' as the seperator; for example, 1.2.3 means the node is in the third layer of a tree where 1 is the root, 1.2 is its parent, and 1.2.3 is the node itself
	Title          string               `json:"title"`        //the window title
	HostName       string               `json:"hostName"`     //could be either the ip address or the dns name
	Disabled       bool                 `json:"disabled"`     //indicates whether this node is disabled or not; The node will not be included in the tree; disable some nodes to avoid creating a new file
	UserName       string               `json:"userName"`     //the username
	AuthMethod     AuthenticationMethod `json:"authMethod"`   //file key or password
	AuthPhrase     string               `json:"authPhrase"`   //depends on authMethod, it is either password or file path of the private key
	TreePositionID treePositionID
}

type sshConfigs []sshConfig

func parseConfigFile(fn string) (sshConfigs, error) {
	var terminals []sshConfig
	configFile, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	contents, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(contents, &terminals)
	if err != nil {
		return nil, nil
	}
	for i := 0; i < len(terminals); i++ {
		terminals[i].TreePositionID = makeFromSingleString(terminals[i].TreePosition)
	}
	return terminals, err
}

func findTree(forest []*terminalTreeNode, sc *sshConfig) *terminalTreeNode {
	if sc == nil {
		return nil
	}
	for _, root := range forest {
		if strings.HasPrefix(sc.TreePosition, root.sc.TreePosition) {
			return root
		}
	}
	//there is no root for this node, and a new tree is created
	tree := newTerminalTreeNode(sc)
	return tree
}

//builds a forest based on the sshConfig slice.
func buildForest(scs sshConfigs) ([]*terminalTreeNode, error) {
	forest := make([]*terminalTreeNode, 0, 1) // Usually, there should only be one tree in forest
	//sorts the configurations so that a root is created first
	sort.Sort(scs)
	for i := range scs {
		var tree *terminalTreeNode
		node := newTerminalTreeNode(&scs[i])
		for _, root := range forest {
			if root.isAncestor(node) {
				tree = root
			}
		}
		if tree == nil {
			//node is a new root, and is added to forest
			tree = node
			forest = append(forest, tree)
			continue
		}
		//finds the root of existing tree, and adds node to the tree
		parentNode := tree.searchParent(node)
		if parentNode.isSamePosition(node) {
			return nil, fmt.Errorf("Node with tree position (%s) already exist", node.sc.TreePosition)
		}
		parentNode.children = append(parentNode.children, node)
	}
	return forest, nil
}

//loadConfigFile loads configuration file into a forest. returns slices of tree roots
//fn: the file name of the configuration file, which should in a json format
func loadConfigFile(fn string) ([]*terminalTreeNode, error) {

	terminals, err := parseConfigFile(fn)
	if err != nil {
		return nil, err
	}
	return buildForest(terminals)
}

//sort.Interface

func (scs sshConfigs) Len() int {
	return len(scs)
}

func (scs sshConfigs) Less(i, j int) bool {
	return len(scs[i].TreePositionID.splitedSring) < len(scs[j].TreePositionID.splitedSring)
}

func (scs sshConfigs) Swap(i, j int) {
	scs[i], scs[j] = scs[j], scs[i]
}
