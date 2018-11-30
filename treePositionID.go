/*
The tree postion ID is a string which can be split into a string array with '.'
The tree position ID is used to indicate the position of the tree node.
For example, a tree node with position ID "1.2.3" means it is a child of node "1.2" and
the node "1.2" is a child of node "1". Please note that the parent node may not exist.
In above example, if node 1.2 not exist, node "1.2.3" will be put under node "1"
*/
package main

import "strings"

const treePositionSeperator string = "."

type treePositionID struct {
	singleString string
	splitedSring []string
}

func makeFromSingleString(str string) treePositionID {
	if str == "" {
		return treePositionID{singleString: "", splitedSring: []string{}}
	}
	id := treePositionID{singleString: str}
	id.splitedSring = strings.Split(str, treePositionSeperator)
	return id
}

func makeFromSplitedString(strs []string) treePositionID {
	if len(strs) == 0 {
		return treePositionID{singleString: "", splitedSring: []string{}}
	}
	id := treePositionID{splitedSring: strs}
	var sb strings.Builder
	i := 0
	for ; i < len(strs)-1; i++ {
		sb.WriteString(strs[i])
		sb.WriteString(treePositionSeperator)
	}
	sb.WriteString(strs[i])
	id.singleString = sb.String()[:len(sb.String())-len(treePositionSeperator)]
	return id
}
