package component

import "fmt"

// Framework ...
type Framework struct{}

// New ...
func New() Framework {
	return Framework{}
}

// Method ...
func (Framework) Method() {
	fmt.Println("(Framework) Method()")
}
