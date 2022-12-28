package main

import (
	"behaviouralpatterns/chainofrespon"
	"behaviouralpatterns/command"
	"behaviouralpatterns/iterator"
	"behaviouralpatterns/mediator"
	"behaviouralpatterns/memento"
	nullobjec "behaviouralpatterns/nullObjec"
	"fmt"
)

func main() {
	chainofrespon.ChainOfRespon()
	fmt.Println("Hello")
	command.Commant()
	fmt.Println("Hello")
	iterator.Iterator()
	fmt.Println("Hello")
	mediator.Mediator()
	fmt.Println("Hello")
	memento.Memento()
	fmt.Println("Hello")
	nullobjec.NullObject()

}
