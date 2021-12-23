package simplify_path

type StringStack struct {
	elements []string
}

func newStringStack() *StringStack {
	return &StringStack{elements: []string{}}
}

func (stack *StringStack) push(element string) {
	stack.elements = append(stack.elements, element)
}

func (stack *StringStack) pop() string {
	element := stack.get()
	if len(stack.elements) == 0 {
		return element
	}
	stack.elements = stack.elements[0 : len(stack.elements)-1]
	return element
}

func (stack *StringStack) get() string {
	length := len(stack.elements)
	if length == 0 {
		return ""
	}
	element := stack.elements[length-1]
	return element
}

func (stack *StringStack) all() []string {
	return stack.elements
}

func simplifyPath(path string) string {
	if len(path) == 0 {
		return ""
	}
	stack := newStringStack()
	for index := 0; index < len(path); {
		currentPathPart := string(path[index])
		if currentPathPart != "/" {
			pathPart := readPathFrom(index, path)
			if pathPart == ".." {
				stack.pop()
			} else if pathPart != "." {
				stack.push(pathPart)
			}
			index = index + len(pathPart)
		} else {
			index = index + 1
		}
	}

	parts := stack.all()
	canonicalPath := "/"
	for index, part := range parts {
		canonicalPath = canonicalPath + part
		if index != len(parts)-1 {
			canonicalPath = canonicalPath + "/"
		}
	}
	return canonicalPath
}

func readPathFrom(index int, path string) string {
	var parts []uint8
	for readIndex := index; readIndex < len(path); readIndex++ {
		pathPart := path[readIndex]
		if string(pathPart) == "/" {
			break
		}
		parts = append(parts, pathPart)
	}
	return string(parts)
}
