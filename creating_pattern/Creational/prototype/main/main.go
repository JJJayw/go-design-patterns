package main

import (
	"fmt"
	"go-design-patterns/creating_pattern/Creational/prototype/lib"
)

func main() {
	file1 := &lib.File{Name: "File1"}
	file2 := &lib.File{Name: "File2"}
	file3 := &lib.File{Name: "File3"}

	folder1 := &lib.Folder{
		Children: []lib.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &lib.Folder{
		Children: []lib.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
