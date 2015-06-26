package detectpath

import (
	"regexp"
)

type MatchedPath struct {
	File  string
	Start int
}

func DetectPath(line string) *MatchedPath {
	// Extract file path from line
	ret := checkPath(line)
	if ret == nil {
		return nil
	}
	// Check if this result is an actual file
	return ret[0]
}

// Check whether line includes file path
// Priority: HomeDir > Basic > JustFile
func checkPath(line string) []*MatchedPath {
	pathes := make([]*MatchedPath, 0)
	if ret := isHomeDirPath(line); ret != nil {
		pathes = append(pathes, ret)
	}
	if ret := isBasicPath(line); ret != nil {
		pathes = append(pathes, ret)
	}
	if ret := isJustFilePath(line); ret != nil {
		pathes = append(pathes, ret)
	}
	return pathes
}

func isBasicPath(line string) *MatchedPath {
	re, err := regexp.Compile("(\\/?([a-z.A-Z0-9-_]+\\/)+[+@a-zA-Z0-9\\-_+.]+\\.[a-zA-Z0-9]+)[:-]?(\\w+)?")
	if err != nil {
		return nil
	}
	matched := re.FindStringSubmatch(line)
	if matched == nil {
		return nil
	}
	return &MatchedPath{matched[1], 0}
}

func isHomeDirPath(line string) *MatchedPath {
	re, err := regexp.Compile("(~\\/([a-z.A-Z0-9-_]+\\/)+[+@a-zA-Z0-9\\-_+.]+\\.[a-zA-Z0-9]+)[:-]?(\\w+)?")
	if err != nil {
		return nil
	}
	matched := re.FindStringSubmatch(line)
	if matched == nil {
		return nil
	}
	return &MatchedPath{matched[1], 0}
}

func isJustFilePath(line string) *MatchedPath {
	re, err := regexp.Compile("([+@a-zA-Z0-9\\-_+.]+\\.[a-zA-Z0-9]+)(\\s|$|:)+")
	if err != nil {
		return nil
	}
	matched := re.FindStringSubmatch(line)
	if matched == nil {
		return nil
	}
	return &MatchedPath{matched[1], 0}
}
