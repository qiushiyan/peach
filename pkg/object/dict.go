package object

import (
	"bytes"
	"fmt"
	"strings"
)

type DictPair struct {
	Key   Object // the original key in DictLiteral, one of String, Integer, or Boolean
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
func (d *Dict) Keys() []Object {
	keys := []Object{}
	for _, pair := range d.Pairs {
		keys = append(keys, pair.Key)
	}
	return keys
}
func (d *Dict) Set(key Object, value Object) {
	d.Pairs[key.(Hashable).Hash()] = DictPair{key, value}
}

func (d *Dict) Get(key Object) (Object, bool) {
	pair, ok := d.Pairs[key.(Hashable).Hash()]
	if !ok {
		return nil, false
	}
	return pair.Value, true
}
