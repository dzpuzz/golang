package main

import (
	"fmt"
)

var array = []int{1, 23, 343, 2, 32, 43}

func main() {
	//定义标准数组

	fmt.Println("原数组:", array)
	makeHeap()
	fmt.Println("建堆后:", array)
	heapSort()
	fmt.Println("排序后:", array)
}

// 建立堆
func makeHeap() {
	n := len(array)
	for c := n/2 - 1; c >= 0; c-- {
		down(c, n)
	}
}

//构建子叶
func down(c, n int) {
	for {
		no1 := 2*c + 1
		if no1 > n {
			break
		}
		no := no1
		if no2 := no1 + 1; no2 < n && !less(no, no2) {
			no = no2
		}
		//构建小顶堆
		if !less(no, c) {
			break
		}
		swap(no, c)
		c = no
	}

}
func less(a, b int) bool {
	return array[a] < array[b]
}

func swap(a, b int) {
	array[a], array[b] = array[b], array[a]
}
func heapSort() {
	//升序 Less(heap[a] > heap[b]) //最大堆
	//降序 Less(heap[a] < heap[b]) //最小堆
	for i := len(array) - 1; i > 0; i-- {
		//移除顶部元素到数组末尾,然后剩下的重建堆,依次循环
		swap(0, i)
		down(0, i)
	}
}
