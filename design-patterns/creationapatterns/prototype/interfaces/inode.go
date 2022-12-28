package interfaces

type Inode interface {
	Print(string)
	Clone() Inode
}
