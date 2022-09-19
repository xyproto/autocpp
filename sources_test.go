package autocpp

import (
	"fmt"
	"testing"

	"github.com/xyproto/env"
)

const exampleProjectDirectory = "~/cppprojects/fireworks"

// an alternative to an "init" function, for initializing a new LocalSystem
var src = func() *Sources {
	src, err := NewSources(env.ExpandUser(exampleProjectDirectory), true)
	if err != nil {
		panic(err)
	}
	return src
}()

func TestIncludes(t *testing.T) {
	shortIncludes := src.ShortIncludes()
	fmt.Println(shortIncludes)
}

func TestPrintIncludeInfo(t *testing.T) {
	src.FindAndPrintIncludePaths(locsys)
}
