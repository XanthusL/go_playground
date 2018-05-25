package decision

import (
	"fmt"
	"sync"
)
//http://docs.jboss.org/drools/release/6.0.0.Final/drools-docs/html/HybridReasoningChapter.html#ReteOO
type (
	// Node
	Node struct {
		//ID     int64
		//Name   string
		Children map[int]*Node // case(action) -> child
		Judge    func([]interface{}) (action int)
	}
	// Tree
	Tree struct {
		// Root
		Root      *Node
		dataCount int
		sync.RWMutex
	}
)

// NewTree
func NewTree(datCnt int, root *Node) *Tree {
	return &Tree{
		Root:      root,
		dataCount: datCnt,
	}
}

// NewNode
func NewNode(f func([]interface{}) int) *Node {
	return &Node{
		Judge:    f,
		Children: make(map[int]*Node),
	}
}

func (tree *Tree) Judge(dat ...interface{}) (action int) {
	if len(dat) < tree.dataCount {
		s := fmt.Sprintf("length of dat %d is not enough, %d needed", len(dat), tree.dataCount)
		panic(s)
	}

	tree.RLock()
	defer tree.RUnlock()

	n := tree.Root
	for {
		x := n.Judge(dat)
		child, ok := n.Children[x]
		if !ok {
			return x
		}
		n = child
	}

}

func (n *Node) SetChild(action int, node *Node) {
	n.Children[action] = node
}
