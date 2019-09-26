package component2

import "fmt"

// Framework ...
type Framework2 struct{}

// New ...
func New() Framework2 {
	return Framework2{}
}

// Method ...
func (Framework2) Method2() {
	fmt.Println("(Framework2) Method2()")
}
