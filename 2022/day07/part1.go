package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Item interface {
	getName() string
	calculateSize() int
	print(level int)
}

type File struct {
	name string
	size int
}

type Directory struct {
	name   string
	items  []Item
	parent Item
}

type Result struct {
	directorySizes []int
}

func (file File) getName() string {
	return file.name
}

func (file File) calculateSize() int {
	return file.size
}

func (file File) print(level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("\033[0;32m%s (%d)\033[0m\n", file.name, file.size)
}

func (directory *Directory) getName() string {
	return directory.name
}

func (directory *Directory) calculateSize() int {
	size := 0
	for _, item := range directory.items {
		size += item.calculateSize()
	}
	return size
}

func (directory *Directory) print(level int) {
	for i := 0; i < level; i++ {
		fmt.Printf("  ")
	}
	fmt.Printf("\033[0;34m%s\033[0m\n", directory.name)
	for _, item := range directory.items {
		item.print(level + 1)
	}
}

func (directory *Directory) findDirectory(name string) *Directory {
	for _, item := range directory.items {
		if item.getName() == name {
			directory, ok := item.(*Directory)
			if ok {
				return directory
			}
		}
	}
	if directory.name == name {
		return directory
	}
	return nil
}

func (directory *Directory) addItem(item Item) {
	directory.items = append(directory.items, item)
}

func createFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}

func createDirectory(name string, parent Item) *Directory {
	return &Directory{
		name:   name,
		items:  make([]Item, 0),
		parent: parent,
	}
}

func main() {
	input, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	// create data structure
	rootDir := createDirectory("/", nil)
	currentDir := rootDir
	for _, line := range strings.Split(string(input), "\n") {
		if line[0] == '$' {
			// its a command
			if line[2] == 'c' {
				// its a CD command
				if line[5] == '.' {
					// going back to parent dir
					directory, ok := currentDir.parent.(*Directory)
					if ok {
						currentDir = directory
					}
				} else {
					// going into a dir
					dirName := line[5:]
					dir := currentDir.findDirectory(dirName)
					if dir == nil {
						panic("Cant find the dir")
					}
					currentDir = dir
				}
			}
		} else if line[0] == 'd' {
			// its dir output
			newDirName := line[4:]
			newDir := createDirectory(newDirName, currentDir)
			currentDir.addItem(newDir)
		} else {
			// its file output
			fileOutputParts := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(fileOutputParts[0])
			newFile := createFile(fileOutputParts[1], fileSize)
			currentDir.addItem(newFile)
		}
	}

	result := &Result{
		directorySizes: make([]int, 0),
	}
	findAnswer(rootDir, result)

	sum := 0
	for _, directorySize := range result.directorySizes {
		sum += directorySize
	}

	fmt.Printf("Result is: %d\n", sum)
}

func findAnswer(dir *Directory, result *Result) {
	dirSize := dir.calculateSize()
	if dirSize <= 100000 {
		result.directorySizes = append(result.directorySizes, dirSize)
	}
	for _, item := range dir.items {
		directory, ok := item.(*Directory)
		if ok {
			findAnswer(directory, result)
		}
	}
}
