package object

import "fmt"

type Number struct {
	Value float64
}

func (n *Number) Inspect() string { return fmt.Sprintf("%v", n.Value) }
func (n *Number) Type() ObjectType {
	return NUMBER_OBJ
}
func (i *Number) Hash() DictKey {
	return DictKey{Type: i.Type(), Value: uint64(i.Value)}
}

// only for builtin functions
type Integer struct {
	Value int64
}

func (n *Integer) Inspect() string { return fmt.Sprintf("%d", n.Value) }
func (n *Integer) Type() ObjectType {
	return NUMBER_OBJ
}
func (i *Integer) Hash() DictKey {
	return DictKey{Type: i.Type(), Value: uint64(i.Value)}
}
