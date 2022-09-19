package autocpp

import (
	"fmt"
	"testing"

	"github.com/xyproto/env"
)

const exampleProjectDirectory = "~/cppprojects/fireworks"

func TestNewSources(t *testing.T) {
	_, err := NewSources(env.ExpandUser(exampleProjectDirectory), true)
	if err != nil {
		t.Fail()
	}
	//fmt.Println(src)
}

func TestIncludes(t *testing.T) {
	src, err := NewSources(env.ExpandUser(exampleProjectDirectory), true)
	if err != nil {
		t.Fail()
	}
	shortIncludes := src.ShortIncludes()
	fmt.Println(shortIncludes)
}

func TestPrintIncludeInfo(t *testing.T) {

	src, err := NewSources(env.ExpandUser(exampleProjectDirectory), true)
	if err != nil {
		t.Fail()
	}

	locsys, err := NewLocalSystem(true)
	if err != nil {
		t.Fail()
	}

	src.PrintIncludeInfo(locsys)
}
