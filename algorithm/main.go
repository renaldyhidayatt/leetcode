package main

import (
	"algoritmAndDs/avltree"
	"algoritmAndDs/fibonancihead"
	"algoritmAndDs/queue"
	"algoritmAndDs/stackgo"
	"algoritmAndDs/treetraversal"
	"fmt"
)

func main() {
	stackgo.StackMain()
	fmt.Println("----------")
	queue.QueueMain()

	fibonancihead.FibonanciMain()
	treetraversal.TreetraversalMain()
	fmt.Println("----------")
	avltree.AvlTreeMain()
}
