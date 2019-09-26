package application

import (
	"fmt"

	"github.com/SpacedMonkeyTCT/comparelayout/layout4/pkg"
)

// Framework ...
type Framework struct {
	c1 pkg.Componenter
	c2 pkg.Componenter2
}

// NewModule ...
func NewModule(c1 pkg.Componenter, c2 pkg.Componenter2) Framework {
	return Framework{c1: c1, c2: c2}
}

// Action ...
func (f Framework) Action() {
	fmt.Println("(Framework) Action()")
	f.c1.Method()
	f.c2.Method2()
}
