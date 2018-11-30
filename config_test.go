/*
this is not a full unit test, I use it more for debug routing for now
TODO: add more unit tests: positive, negative
*/
package main

import (
	"fmt"
	"testing"
)

func TestParseConfigFile(t *testing.T) {
	configNodes, err := parseConfigFile("example.conf.json")
	if err != nil {
		t.Fail()
		return
	}
	for _, sc := range configNodes {
		fmt.Printf("\n%+v\n", sc)
	}
}

func TestLoadConfigFile(t *testing.T) {
	forest, err := loadConfigFile("example.conf.json")
	if err != nil {
		t.Fail()
		return
	}
	for _, tree := range forest {
		treeString := tree.TreeString()
		fmt.Println(treeString)
	}

	nodes := forest[0].toSlice()
	for i := range nodes {
		fmt.Printf("\n%+v\n", nodes[i].String())
	}

	node := forest[0].searchNode("1.1.1")
	if node == nil {
		t.Fail()
		return
	}
	fmt.Printf("\n%+v\n", node.String())

	nodes = node.toSlice()
	for i := range nodes {
		fmt.Printf("\n%+v\n", nodes[i].String())
	}
}
