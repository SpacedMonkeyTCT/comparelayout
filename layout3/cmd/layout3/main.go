package main

import (
	"github.com/SpacedMonkeyTCT/comparelayout/layout3/internal/application"
	component "github.com/SpacedMonkeyTCT/comparelayout/layout3/pkg/component/implement"
	component2 "github.com/SpacedMonkeyTCT/comparelayout/layout3/pkg/component2/implement"
)

func main() {
	c1 := component.New()
	c2 := component2.New()
	m := application.NewModule(c1, c2)
	m.Action()
}
