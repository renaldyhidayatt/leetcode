package trie

import "fmt"

func ExampleNode() {

	node := NewNode()

	node.Insert("nikola")
	node.Insert("tesla")

	fmt.Println(node.Size())
	fmt.Println(node.Capacity())

	fmt.Println(node.Find("thomas"))
	fmt.Println(node.Find("edison"))
	fmt.Println(node.Find("nikola"))
	node.Remove("tesla")
	fmt.Println(node.Find("tesla"))

	fmt.Println(node.Size())
	fmt.Println(node.Capacity())

}
