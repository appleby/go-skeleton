package goskeleton

import "fmt"

// Foo has a Name.
type Foo struct {
	Name string
}

// Greet says Hello, Foo.Name.
func (foo Foo) Greet() string {
	return fmt.Sprintf("Hello, %s", foo.Name)
}
