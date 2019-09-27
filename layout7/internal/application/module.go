package application

import (
	"fmt"

	"github.com/SpacedMonkeyTCT/comparelayout/layout7/pkg/component"
	"github.com/SpacedMonkeyTCT/comparelayout/layout7/pkg/component2"
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
