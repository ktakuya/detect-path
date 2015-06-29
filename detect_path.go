package detectpath

import (
	"fmt"
	"os"
	"regexp"
)

type MatchedPath struct {
	File  string
	Start int
}

func DetectPath(line string) *MatchedPath {
	// Extract file path from line
	ret := CheckPath(line)
	if ret == nil {
		return nil
	}
	// Check if this result is an actual file
	fmt.Println(ret)
	for _, path := range ret {
		if Exists(path.File) {
			return path
		}
	}
	return nil
}

// Check whether line includes file path
// Priority: HomeDir > Basic > JustFile
func CheckPath(line string) []*MatchedPath {
	regs := []string{
		// Home dir
		"(~\\/([a-z.A-Z0-9-_]+\\/)+[+@a-zA-Z0-9\\-_+.]+\\.[a-zA-Z0-9]+)[:-]?(\\w+)?",
		// Normal
		"(\\/?([a-z.A-Z0-9-_]+\\/)+[+@a-zA-Z0-9\\-_+.]+\\.[a-zA-Z0-9]+)[:-]?(\\w+)?",
		// Just file
		"([+@a-zA-Z0-9\\-_+.]+\\.[a-zA-Z0-9]+)(\\s|$|:)+",
	}
	pathes := make([]*MatchedPath, 0)
	for _, v := range regs {
		if ret := comparePathWithRegexp(line, v); ret != nil {
			pathes = append(pathes, ret)
		}
	}
	return pathes
}

func comparePathWithRegexp(line, reg string) *MatchedPath {
	re, err := regexp.Compile(reg)
	if err != nil {
		return nil
	}
	matched := re.FindStringSubmatch(line)
	if matched == nil {
		return nil
	}
	return &MatchedPath{matched[1], 0}
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
