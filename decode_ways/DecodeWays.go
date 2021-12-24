package decode_ways

import (
	"strconv"
	"strings"
)

func numDecodings(str string) int {
	var numDecodingsInner func(fromIndex, toIndex int, currentIndex int) int

	length := len(str)
	numDecodingsCache := make(map[int]int)
	numDecodingsInner = func(fromIndex, toIndex int, currentIndex int) int {
		var part string
		var intPart int
		if toIndex <= length {
			part = str[fromIndex:toIndex]
			intPart, _ = strconv.Atoi(part)
		}
		ways, ok := numDecodingsCache[currentIndex]
		switch {
		case part == "0" || strings.Index(part, "0") == 0:
			return 0
		case intPart > 26 && currentIndex != 0:
			return 0
		case currentIndex > length:
			return 0
		case ok:
			return ways
		case currentIndex == length:
			return 1
		}
		nextIndex := currentIndex + 1
		numDecodingsCache[currentIndex] =
			numDecodingsInner(currentIndex, nextIndex, nextIndex) + numDecodingsInner(currentIndex, nextIndex+1, nextIndex+1)
		return numDecodingsCache[currentIndex]
	}
	return numDecodingsInner(0, length, 0)
}
