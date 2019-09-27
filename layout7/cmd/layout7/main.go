package main

import (
	"github.com/SpacedMonkeyTCT/comparelayout/layout7/internal/application"
	"github.com/SpacedMonkeyTCT/comparelayout/layout7/pkg/component"
	"github.com/SpacedMonkeyTCT/comparelayout/layout7/pkg/component2"
)

func main() {
	c1 := component.New()
	c2 := component2.New()
	m := application.NewModule(c1, c2)
	m.Action()
}
