package object

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	max_display = 50
)

// interface for all vectors
type IVector interface {
	Type() ObjectType
	Inspect() string
	Length() int
	Head(n int) Object
	Tail(n int) Object
}

type BaseVector struct {
	Elements []Object
}

func (bv *BaseVector) Inspect(dtype string) string {
	vectorLength := len(bv.Elements)
	var out bytes.Buffer
	elements := []string{}
	for i, el := range bv.Elements {
		if i <= max_display {
			elements = append(elements, el.Inspect())
		}
	}

	out.WriteString(fmt.Sprintf(dtype+"Vector with %d elements\n", vectorLength))
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	if vectorLength >= max_display {
		out.WriteString(", ...")
	}

	out.WriteString("]")
	return out.String()
}

func (bv *BaseVector) Length() int {
	return len(bv.Elements)
}

func (bv *BaseVector) Head(n int) []Object {
	if n > len(bv.Elements) {
		n = len(bv.Elements)
	}
	return bv.Elements[:n]
}

func (bv *BaseVector) Tail(n int) []Object {
	return bv.Elements[len(bv.Elements)-n:]
}

// vector with mixed types
type Vector struct {
	BaseVector
}

func (v *Vector) Type() ObjectType { return VECTOR_OBJ }
func (v *Vector) Inspect() string  { return v.BaseVector.Inspect("") }
func (v *Vector) Length() int      { return v.BaseVector.Length() }
func (v *Vector) Head(n int) Object {
	return &Vector{BaseVector{v.BaseVector.Head(n)}}
}
func (v *Vector) Tail(n int) Object {
	return &Vector{BaseVector{v.BaseVector.Tail(n)}}
}

// vectors with object.Number
type NumericVector struct {
	BaseVector
}

func (nv *NumericVector) Type() ObjectType { return NUMERIC_VECTOR_OBJ }
func (nv *NumericVector) Inspect() string  { return nv.BaseVector.Inspect("Numeric") }
func (nv *NumericVector) Length() int      { return nv.BaseVector.Length() }
func (nv *NumericVector) Head(n int) Object {
	return &NumericVector{BaseVector{nv.BaseVector.Head(n)}}
}
func (nv *NumericVector) Tail(n int) Object {
	return &NumericVector{BaseVector{nv.BaseVector.Tail(n)}}
}

// vectors with object.String
type CharacterVector struct {
	BaseVector
}

func (cv *CharacterVector) Type() ObjectType { return NUMERIC_VECTOR_OBJ }
func (cv *CharacterVector) Inspect() string  { return cv.BaseVector.Inspect("Character") }
func (cv *CharacterVector) Length() int      { return cv.BaseVector.Length() }
func (cv *CharacterVector) Head(n int) Object {
	return &CharacterVector{BaseVector{cv.BaseVector.Head(n)}}
}
func (cv *CharacterVector) Tail(n int) Object {
	return &CharacterVector{BaseVector{cv.BaseVector.Tail(n)}}
}

// vectors with object.Boolean
type LogicalVector struct {
	BaseVector
}

func (lv *LogicalVector) Type() ObjectType { return NUMERIC_VECTOR_OBJ }
func (lv *LogicalVector) Inspect() string  { return lv.BaseVector.Inspect("Logical") }
func (lv *LogicalVector) Length() int      { return lv.BaseVector.Length() }
func (lv *LogicalVector) Head(n int) Object {
	return &LogicalVector{BaseVector{lv.BaseVector.Head(n)}}
}
func (lv *LogicalVector) Tail(n int) Object {
	return &LogicalVector{BaseVector{lv.BaseVector.Tail(n)}}
}
