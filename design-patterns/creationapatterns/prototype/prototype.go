package prototype

import (
	"creationapatterns/prototype/interfaces"
	"creationapatterns/prototype/models"
	"fmt"
)

func ProtoType() {
	file1 := &models.File{Name: "File1"}
	file2 := &models.File{Name: "File2"}
	file3 := &models.File{Name: "File3"}
	folder1 := &models.Folder{
		Childrens: []interfaces.Inode{file1},
		Name:      "Folder1",
	}
	folder2 := &models.Folder{
		Childrens: []interfaces.Inode{folder1, file2, file3},
		Name:      "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")
	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
