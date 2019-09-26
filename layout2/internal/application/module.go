package application

import (
	"fmt"

	component "github.com/SpacedMonkeyTCT/comparelayout/layout2/pkg/interfaces"
)

// Framework ...
type Framework struct {
	c component.Componenter
}

// NewModule ...
func NewModule(c component.Componenter) Framework {
	return Framework{c: c}
}

// Action ...
func (f Framework) Action() {
	fmt.Println("(Framework) Action()")
	f.c.Method()
}
