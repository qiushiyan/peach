package object

import "fmt"

type Number struct {
	Value float64
}

func (n *Number) Inspect() string { return fmt.Sprintf("%v", n.Value) }
func (n *Number) Type() ObjectType {
	return NUMBER_OBJ
}
