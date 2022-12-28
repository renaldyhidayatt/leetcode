package main

import (
	"creationapatterns/abstractfactory"
	"creationapatterns/builder"
	"creationapatterns/factory"
	"creationapatterns/objeckpool"
	"creationapatterns/prototype"
	"creationapatterns/singleton"
	"fmt"
)

func main() {
	abstractfactory.AbstractFactory()
	fmt.Println("Hello")
	builder.BuilderPatterns()
	fmt.Println("Hello")
	factory.Factory()
	fmt.Println("Hello")
	objeckpool.ObjeckPool()
	fmt.Println("Hello")
	prototype.ProtoType()
	fmt.Println("Hello")
	singleton.SingleTon()
}
