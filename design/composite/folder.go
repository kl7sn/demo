package main

import (
	"fmt"
)

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("Searching in folder %s for keyword %s\n", f.name, keyword)
	for _, component := range f.components {
		component.search(keyword)
	}
}

func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}
