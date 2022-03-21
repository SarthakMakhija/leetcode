package file_system

import (
	"strings"
)

type FileSystem struct {
	root map[string]*fileSystemNode
}

type fileSystemNode struct {
	nextNodeByPathPart map[string]*fileSystemNode
	value              int
}

func Constructor() FileSystem {
	fileSystem := FileSystem{root: make(map[string]*fileSystemNode)}
	fileSystem.root["/"] = &fileSystemNode{nextNodeByPathPart: make(map[string]*fileSystemNode), value: -1}
	return fileSystem
}

func (fileSystem *FileSystem) CreatePath(path string, value int) bool {
	if len(path) == 0 || path == "/" {
		return false
	}

	pathParts := strings.Split(path, "/")[1:]
	head, pathPartsLength := fileSystem.root["/"], len(pathParts)
	for index := 0; index < len(pathParts)-1; index++ {
		pathPart := pathParts[index]
		next, ok := head.nextNodeByPathPart[pathPart]
		if !ok {
			return false
		}
		head = next
	}
	if _, ok := head.nextNodeByPathPart[pathParts[pathPartsLength-1]]; ok {
		return false
	}
	head.nextNodeByPathPart[pathParts[pathPartsLength-1]] =
		&fileSystemNode{
			nextNodeByPathPart: make(map[string]*fileSystemNode),
			value:              -1,
		}
	head = head.nextNodeByPathPart[pathParts[pathPartsLength-1]]
	head.value = value
	return true
}

func (fileSystem *FileSystem) Get(path string) int {
	if len(path) == 0 || path == "/" {
		return -1
	}

	pathParts := strings.Split(path, "/")[1:]
	head := fileSystem.root["/"]
	for _, pathPart := range pathParts {
		next, ok := head.nextNodeByPathPart[pathPart]
		if !ok {
			return -1
		}
		head = next
	}
	return head.value
}
