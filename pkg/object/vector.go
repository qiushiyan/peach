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
	Values() []Object
	// Index(object Object) Object
}

// vector with mixed types
type Vector struct {
	Elements []Object
}

func (v *Vector) Type() ObjectType { return VECTOR_OBJ }
func (v *Vector) Inspect() string  { return inspectVector(v, "") }
func (v *Vector) Values() []Object { return v.Elements }

// vectors with object.Number
type NumericVector struct {
	Elements []Object
}

func (nv *NumericVector) Type() ObjectType { return NUMERIC_VECTOR_OBJ }
func (nv *NumericVector) Inspect() string  { return inspectVector(nv, "Numeric") }
func (nv *NumericVector) Values() []Object { return nv.Elements }

// vectors with object.String
type CharacterVector struct {
	Elements []Object
}

func (cv *CharacterVector) Type() ObjectType { return NUMERIC_VECTOR_OBJ }
func (cv *CharacterVector) Inspect() string  { return inspectVector(cv, "Characetr") }
func (cv *CharacterVector) Values() []Object { return cv.Elements }

// vectors with object.Boolean
type LogicalVector struct {
	Elements []Object
}

func (lv *LogicalVector) Type() ObjectType { return NUMERIC_VECTOR_OBJ }
func (lv *LogicalVector) Inspect() string  { return inspectVector(lv, "Logical") }
func (lv *LogicalVector) Values() []Object { return lv.Elements }

// helper function for inspecting
func inspectVector(v IVector, dtype string) string {
	values := v.Values()
	if len(values) == 0 {
		return "[]"
	}

	var out bytes.Buffer
	elements := []string{}
	for i, el := range values {
		if i <= max_display {
			elements = append(elements, el.Inspect())
		}
	}

	out.WriteString(fmt.Sprintf(dtype+"Vector with %d elements\n", len(values)))
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	if len(values) >= max_display {
		out.WriteString(", ...")
	}

	out.WriteString("]")
	return out.String()
}
