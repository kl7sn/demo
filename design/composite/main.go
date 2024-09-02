package main

func main() {
	file1 := NewFile("file1")
	file2 := NewFile("file2")
	file3 := NewFile("file3")

	folder1 := NewFolder("folder1")
	folder1.add(file1)

	folder2 := NewFolder("folder2")
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("hello") // should return file1
}
