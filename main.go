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
	cmdPrint := flag.Bool("p", false, "print the forest in the configuration file")
	cmdRun := flag.Bool("r", false, "run the command on remote hosts")
	flag.Parse()

	forest, err := loadConfigFile(*configFile)
	errHandle(err)

	lineBreak := "================================================================================================"
	if *cmdPrint {
		for _, tree := range forest {
			fmt.Println(lineBreak)
			fmt.Println(tree.TreeString())
			fmt.Println(lineBreak)
		}
		return
	}

	if *cmdRun {
		var remoteHosts []*terminalTreeNode
		for _, tree := range forest {
			if rootNode := tree.searchNode(*rootNodeID); rootNode != nil {
				remoteHosts = rootNode.toSlice()
			}
		}

		for _, rh := range remoteHosts {
			output, err := runSSHCommand(rh, cmd)
			fmt.Println(lineBreak)
			fmt.Println("==", rh.sc.TreePosition, rh.sc.Title)
			fmt.Println(lineBreak)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(output)
			}
		}
	}

}
