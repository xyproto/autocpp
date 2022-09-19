package autocpp

import (
	"fmt"
	"testing"
)

func TestNewLocalSystem(t *testing.T) {
	_, err := NewLocalSystem(true)
	if err != nil {
		t.Fail()
	}
}

func TestLocSysIncludes(t *testing.T) {
	locsys, err := NewLocalSystem(true)
	if err != nil {
		t.Fail()
	}
	fmt.Println(locsys.IncludeFiles())
}
