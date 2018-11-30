package main

type terminalTreeNodeOperation interface {
	nodeOperation(treeNode *terminalTreeNode) error
}

type terminalTreeNodeMap struct {
	m map[string]*terminalTreeNode
}

func (tmap *terminalTreeNodeMap) nodeOperation(treeNode *terminalTreeNode) error {
	if tmap.m == nil {
		tmap.m = make(map[string]*terminalTreeNode)
	}
	tmap.m[treeNode.nodeID.singleString] = treeNode
	return nil
}
