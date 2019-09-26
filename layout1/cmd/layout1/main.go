package main

import (
	"github.com/SpacedMonkeyTCT/comparelayout/layout1/internal/application"
	"github.com/SpacedMonkeyTCT/comparelayout/layout1/pkg/component"
)

func main() {
	c := component.New()
	m := application.NewModule(c)
	m.Action()
}
