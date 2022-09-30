package object

import (
	"bytes"
	"fmt"
	"strings"
)

type DictPair struct {
	Key   Object // the origional key in DictLiteral, one of String, Integer, or Boolean
	Value Object
}

// {"a": 1} => &Dict{Pairs: {"a".Hash(): {"a": 1}}}
type Dict struct {
	Pairs map[DictKey]DictPair
}

func (d *Dict) Type() ObjectType { return DICT_OBJ }
func (d *Dict) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range d.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
