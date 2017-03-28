package main

import (
	"fmt"
)

const (
	Error = -10000000
)

type Element int64

type LinkNode struct {
	Prior *LinkNode
	Next  *LinkNode
	Data  Element
}
type LinkNoder interface {
	add(head *LinkNode, new *LinkNode)
	GetLength(head *LinkNode)
	Search(head *LinkNode, data Element)
	Traverse(head *LinkNode)
	Delete(head *LinkNode, index int)
	Insert(head *LinkNode, index int, data Element)
	TraversePrior(node *LinkNode)
}

func add(head *LinkNode, data Element) {
	point := head
	for point.Next != nil {
		point = point.Next //移位
	}
	var node LinkNode //新节点
	node.Prior = point
	point.Next = &node
	node.Data = data
}

func Traverse(head *LinkNode) {
	point := head
	for point.Next != nil {
		fmt.Println(point.Next.Data)
		point = point.Next
	}
	fmt.Println("traverse OK")
}

func main() {
	var nodeArray []Element
	var head LinkNode = LinkNode{Data: 0, Prior: nil, Next: nil}
	for i := 0; i < 10; i++ {
		nodeArray = append(nodeArray, Element(i+1+i*100))
		add(&head, nodeArray[i])
	}
	Traverse(&head)
	GetLength(&head)
	a := Insert(&head, 2, 666)
	fmt.Println(a)
	Traverse(&head)
	fmt.Println("遍历前序")
	TraversePrior(&a)
	fmt.Println("遍历前序完成")
	Delete(&head, 2)
	Traverse(&head)
}

func GetLength(head *LinkNode) int {
	point := head
	var length int
	for point.Next != nil {
		point = point.Next
		length++
	}
	return length
}

func Insert(head *LinkNode, index int, data Element) LinkNode {
	var node LinkNode
	point := head
	if index < 0 || index > GetLength(head) {
		fmt.Println(Error)
		return node
	} else {
		for i := 0; i < index; i++ {
			point = point.Next
		}
		point.Prior.Next = &node
		node.Prior = point.Prior
		node.Next = point
		node.Data = data
		point.Prior = &node
		return node
	}
}
func Delete(head *LinkNode, index int) LinkNode {
	var node LinkNode
	point := head
	if index < 0 || index > GetLength(head) {
		fmt.Println(Error)
		return node
	} else {
		for i := 0; i < index; i++ {
			point = point.Next
		}
		point.Next.Prior = point.Prior
		point.Prior.Next = point.Next
		return node
	}
}

//遍历前序
func TraversePrior(node *LinkNode) {
	point := node.Prior
	for point.Prior != nil {
		fmt.Println(point.Data)
		point = point.Prior
	}
}
func Search(head *LinkNode, data Element) {
	point := head
	index := 0
	for point.Nest != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			break
		} else {
			index++
			point = point.Nest
			if index > GetLength(head)-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
}
