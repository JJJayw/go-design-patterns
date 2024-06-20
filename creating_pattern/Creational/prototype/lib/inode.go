package lib

type Inode interface {
	Print(string)
	Clone() Inode
}
