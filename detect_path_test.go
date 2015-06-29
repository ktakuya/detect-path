package detectpath

import (
	"testing"
)

func TestDetectPath(t *testing.T) {
	testCases := []struct {
		Input  string
		Output string
	}{
		{
			"./feer/fwef.go",
			"./feer/fwef.go",
		},
		{
			"./feer-23/fwef-23.go",
			"./feer-23/fwef-23.go",
		},
		{
			"./feer-23/fwef-23.go:fjw",
			"./feer-23/fwef-23.go",
		},
		{
			"~/aaa/aaa.go",
			"~/aaa/aaa.go",
		},
		{
			"aaa.go",
			"aaa.go",
		},
		{
			"fjwoeifj ./aaa/bbb.sh fjweoi",
			"./aaa/bbb.sh",
		},
		{
			"fjwioajf ~/aaa/bbb.sh fwjeof",
			"~/aaa/bbb.sh",
		},
		{
			" /foo/bar/README.md                      |  9 ++++-",
			"/foo/bar/README.md",
		},
	}
	for _, v := range testCases {
		ret := CheckPath(v.Input)
		if ret == nil {
			t.Fatalf("\nExpected: %#v\n\nOutput: nil", v.Output)
		}
		if v.Output != ret[0].File {
			t.Fatalf("\nExpected: %#v\n\nOutput: %#v", v.Output, ret[0].File)
		}
	}
}
