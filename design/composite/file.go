package main

import (
	"fmt"
)

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("search keyword %s in file %s\n", keyword, f.getName())
}

func (f *File) getName() string {
	return f.name
}

func NewFile(name string) *File {
	return &File{name: name}
}
