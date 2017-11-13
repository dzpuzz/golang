package main

import (
	"fmt"

	tree "./BSTree"
)

func main() {

	num := [6]int{1, 5, 4, 3, 2, 6}
	btsTree := tree.NewBSTree()
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < len(num); i++ {
		// num[i] = r.Intn(100)
		btsTree.InsertNode(tree.NewBSTNode(num[i]))
	}
	inOrderList := btsTree.PostOrderTraversal()
	for e := inOrderList.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}
