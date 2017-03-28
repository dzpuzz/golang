package main

import (
	"fmt"
)

//单向链表实现
const (
	ERROR = -1000000001
)

type Element int64

type LinkNode struct {
	Data Element   //数据域
	Next *LinkNode //指针域，指向下一个节点
}

type LinkNoder interface {
	Add(head *LinkNode, new *LinkNode)
}

func Add(head *LinkNode, data Element) {
	point := head
	for point.Next != nil {
		point = point.Next
	}
	var node LinkNode  //新节点
	point.Next = &node //赋值
	node.Data = data
}

func main() {
	var head LinkNode = LinkNode{Data: 0, Next: nil}
	head.Data = 0
	var nodeArray []Element
	for i := 0; i < 10; i++ {
		nodeArray = append(nodeArray, Element(i+1+i*100))
		Add(&head, nodeArray[i])
	}
	Traverse(&head)
	Search(&head, 203)
	de := Delete(&head, 1)
	fmt.Printf("delete:%v\n\r", de)
	fmt.Println("====================")
	Traverse(&head)
}
func Traverse(head *LinkNode) {
	point := head
	for point.Next != nil {
		fmt.Println(point.Next.Data)
		point = point.Next
	}
}
func Search(head *LinkNode, data Element) {
	point := head
	index := 0
	for point.Next != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			break
		} else {
			index++
			point = point.Next
			if index >= GetLength(head) {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
}
func GetLength(head *LinkNode) int {
	point := head
	var length int
	for point.Next != nil {
		length++
		point = point.Next
	}
	return length
}

func Insert(head *LinkNode, index int, data Element) {
	if index < 0 || index > GetLength(head) {
		fmt.Printf("invalid index")
	} else {
		point := head
		for i := 0; i <= index-1; i++ {
			point = point.Next //移位至插入点
		}
		var node LinkNode //新节点
		node.Data = data
		node.Next = point.Next
		point.Next = &node
	}
}
func Delete(head *LinkNode, index int) Element {
	if index < 0 || index > GetLength(head) {
		fmt.Printf("invalid index")
		return ERROR
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Next //移位
		}
		data := point.Next.Data
		point.Next = point.Next.Next
		return data
	}
}
