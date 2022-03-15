package simplify_path

type pathStack struct {
	elements []string
}

func newPathStack() *pathStack {
	return &pathStack{elements: []string{}}
}

func (stack *pathStack) push(element string) {
	stack.elements = append(stack.elements, element)
}

func (stack *pathStack) pop() string {
	element := stack.get()
	if len(stack.elements) == 0 {
		return element
	}
	stack.elements = stack.elements[0 : len(stack.elements)-1]
	return element
}

func (stack *pathStack) get() string {
	length := len(stack.elements)
	if length == 0 {
		return ""
	}
	element := stack.elements[length-1]
	return element
}

func (stack *pathStack) all() []string {
	return stack.elements
}

func (stack *pathStack) isEmpty() bool {
	return len(stack.elements) == 0
}

func (stack *pathStack) path(separator string) string {
	path := ""
	for !stack.isEmpty() {
		pathPart := separator + stack.pop()
		path = pathPart + path
	}
	return path
}

func simplifyPath(path string) string {
	if len(path) == 0 {
		return ""
	}
	pathStack := newPathStack()
	for index := 0; index < len(path); {
		pathPart := ""
		for index < len(path) && path[index] != '/' {
			pathPart = pathPart + string(path[index])
			index++
		}
		switch {
		case pathPart == "..":
			pathStack.pop()
		case pathPart == "." || len(pathPart) == 0:
			//nothing
		default:
			pathStack.push(pathPart)
		}
		index++
	}
	canonicalPath := pathStack.path("/")
	if len(canonicalPath) == 0 {
		return "/"
	}
	return canonicalPath
}
