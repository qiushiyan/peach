package object

import (
	"fmt"
	"hash/fnv"
)

type String struct {
	Value string
}

func (s *String) Inspect() string { return fmt.Sprintf("\"%s\"", s.Value) }
func (s *String) Type() ObjectType {
	return STRING_OBJ
}
func (s *String) Hash() DictKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return DictKey{Type: s.Type(), Value: h.Sum64()}
}
