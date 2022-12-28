package models

import (
	"creationapatterns/prototype/interfaces"
	"fmt"
)

type Folder struct {
	Childrens []interfaces.Inode
	Name      string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Childrens {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() interfaces.Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildrens []interfaces.Inode
	for _, i := range f.Childrens {
		copy := i.Clone()
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.Childrens = tempChildrens
	return cloneFolder
}
