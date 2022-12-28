package models

import (
	"creationapatterns/prototype/interfaces"
	"fmt"
)

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() interfaces.Inode {
	return &File{Name: f.Name + "_clone"}
}
