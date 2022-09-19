package autocpp

import (
	"fmt"
	"testing"
)

// an alternative to an "init" function, for initializing a new LocalSystem
var locsys = func() *LocalSystem {
	locsys, err := NewLocalSystem(true)
	if err != nil {
		panic(err)
	}
	return locsys
}()

func TestLocSysIncludes(t *testing.T) {
	fmt.Println(locsys.IncludeFiles())
}
