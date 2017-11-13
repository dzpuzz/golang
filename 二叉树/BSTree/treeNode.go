package tree

import (
	"container/list"
)

//二叉排序树树
type BSTree struct {
	root *BSTNode
	size int
}

//二叉树节点
type BSTNode struct {
	data   interface{}
	lchild *BSTNode
	rchild *BSTNode
	parent *BSTNode
}

//新建一个节点
func NewBSTNode(e interface{}) *BSTNode {
	return &BSTNode{data: e}
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

// 插入节点
func (tree *BSTree) InsertNode(node *BSTNode) *BSTree {
	if tree.root == nil {
		tree.root = node
		return tree
	}
	thisNode := tree.root
	for {
		if thisNode.data.(int) < node.data.(int) {
			if thisNode.rchild == nil {
				thisNode.rchild = node
				node.parent = thisNode
				break
			} else {
				thisNode = thisNode.rchild
				continue
			}
		} else {
			if thisNode.lchild == nil {
				thisNode.lchild = node
				node.parent = thisNode
				break
			} else {
				thisNode = thisNode.lchild
				continue
			}
		}
	}
	return tree
}

// 中序遍历 morris
func (tree *BSTree) InOrderTraversal() *list.List {
	l := list.New()
	cur := tree.root
	for cur != nil {
		// 左节点为空直接输出
		if cur.lchild == nil {
			l.PushBack(cur.data)
			cur = cur.rchild
		} else {
			// find predecessor
			prev := cur.lchild
			for prev.rchild != nil && prev.rchild != cur {
				prev = prev.rchild
			}
			if prev.rchild == nil {
				prev.rchild = cur
				cur = cur.lchild
			} else {
				l.PushBack(cur.data)
				prev.rchild = nil
				cur = cur.rchild
			}
		}
	}
	return l
}

// 前序遍历 morris
func (tree *BSTree) PerOrderTraversal() *list.List {
	l := list.New()
	cur := tree.root
	for cur != nil {
		// 左节点为空直接输出
		if cur.lchild == nil {
			l.PushBack(cur.data)
			cur = cur.rchild
		} else {
			// find predecessor
			prev := cur.lchild
			for prev.rchild != nil && prev.rchild != cur {
				prev = prev.rchild
			}
			if prev.rchild == nil {
				l.PushBack(cur.data)
				prev.rchild = cur
				cur = cur.lchild
			} else {
				prev.rchild = nil
				cur = cur.rchild
			}
		}
	}
	return l
}

// 后续遍历 morris
func (tree *BSTree) PostOrderTraversal() *list.List {
	l := list.New()
	assistNode := NewBSTNode(0)
	assistNode.lchild = tree.root
	cur := assistNode
	for cur != nil {
		if cur.lchild == nil {
			cur = cur.rchild
		} else {
			prev := cur.lchild
			for prev.rchild != nil && prev.rchild != cur {
				prev = prev.rchild
			}
			if prev.rchild == nil {
				prev.rchild = cur
				cur = cur.lchild
			} else {
				prev.rchild = nil
				getNodeIntoList(cur.lchild, prev, l)
				cur = cur.rchild
			}
		}
	}
	return l
}
func getReverseNodes(from *BSTNode, to *BSTNode) {

	if from == to {
		return
	}
	x := from
	y := from.rchild
	for true {
		z := y.rchild
		y.rchild = x
		x = y
		y = z
		if x == to {
			break
		}
	}
}
func getNodeIntoList(from *BSTNode, to *BSTNode, l *list.List) {
	getReverseNodes(from, to)
	for true {
		l.PushBack(to.data)
		if to == from {
			break
		}
		to = to.rchild
	}
	getReverseNodes(to, from)
}
