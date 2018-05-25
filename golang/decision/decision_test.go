package decision

import (
	"testing"
)

type action int

const (
	normal action = 1
	thin   action = 2
	fat    action = 3

	tall  action = 10
	short action = 11
)

func TestNewNode(t *testing.T) {
	// [height, weight, ... ]
	// [180, 100, ... ]
	/*
				O    -   -   -   -> height>170 ? left:right
			   / \
			  /   \
			 /	   \
			o	    o  -   -  -  -> weight>100 ? left:right
		   / \	   / \
		  /   \   /   \
		 1     2  3    1

	*/
	n := NewNode(isTall)
	n.SetChild(int(tall), NewNode(func(src []interface{}) int {
		h := src[1].(int)
		if h > 100 {
			return int(normal)
		}
		return int(thin)
	}))
	n.SetChild(int(short), NewNode(func(src []interface{}) int {
		h := src[1].(int)
		if h > 100 {
			return int(fat)
		}
		return int(normal)
	}))
	tree := NewTree(2, n)
	a := tree.Judge(180, 100)
	t.Log(180, 100, action(a))
	a = tree.Judge(180, 300)
	t.Log(180, 300, action(a))
	a = tree.Judge(160, 300)
	t.Log(160, 300, action(a))
	a = tree.Judge(160, 50)
	t.Log(160, 50, action(a))
}

func isTall(src []interface{}) int {
	// return src[0] > src[1]
	h := src[0].(int)
	if h > 170 {
		return int(tall)
	}
	return int(short)
}

func (a action) String() string {
	switch a {
	case normal:
		return "normal"
	case thin:
		return "slim"
	case fat:
		return "fat"
	default:
		return "undefined"
	}
}
