package main

import (
	"flag"
	"fmt"
	"os"
)

func errHandle(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	//command line parameters
	configFile := flag.String("conf", "config.json", "name of the configuration")
	rootNodeID := flag.String("root", "1", "root node positionID")
	cmd := flag.String("cmd", "ls", "command send to remote hosts")
	flag.Parse()

	forest, err := loadConfigFile(*configFile)
	errHandle(err)
	var remoteHosts []*terminalTreeNode
	for _, tree := range forest {
		if rootNode := tree.searchNode(*rootNodeID); rootNode != nil {
			remoteHosts = rootNode.toSlice()
		}
	}

	for _, rh := range remoteHosts {
		runSSHCommand(rh, cmd)
	}
}
