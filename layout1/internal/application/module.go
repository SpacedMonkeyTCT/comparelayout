package application

import (
	"fmt"

	component "github.com/SpacedMonkeyTCT/comparelayout/layout1/pkg/component/interface"
	component2 "github.com/SpacedMonkeyTCT/comparelayout/layout1/pkg/component2/interface"
)

// Framework ...
type Framework struct {
	c1 component.Componenter
	c2 component2.Componenter2
}

// NewModule ...
func NewModule(c1 component.Componenter, c2 component2.Componenter2) Framework {
	return Framework{c1: c1, c2: c2}
}

// Action ...
func (f Framework) Action() {
	fmt.Println("(Framework) Action()")
	f.c1.Method()
	f.c2.Method2()
}
