package object

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

const (
	max_display = 50
)

// interface for all vectors
type IVector interface {
	Type() ObjectType
	Inspect() string
	Print() string
	ElementType() ObjectType
	Length() int
	Head(n int) Object
	Tail(n int) Object
	Values() []Object
	Slice(start int, end int) Object
	Append(objects ...Object) Object
	Set(start int, val Object) Object
	Replace(start int, end int, replacement IVector) Object
	Infix(func(Object, Object) Object, IVector) Object
}

// the inner element for all vectors, provides implementation for methods in IVector (doesn't have to be of IVector itself)
type BaseVector struct {
	Elements []Object
}

func (bv *BaseVector) Type() ObjectType        { return VECTOR_OBJ }
func (bv *BaseVector) ElementType() ObjectType { return NULL_OBJ }
func (bv *BaseVector) Inspect() string {
	// check for empty vectors created by std vector(num)
	vectorLength := len(bv.Elements)
	var out bytes.Buffer
	elements := []string{}
	for i, el := range bv.Elements {
		if el == nil {
			return "empty vector with " + fmt.Sprintf("%d", i) + " filled elements"
		}
		if i <= max_display {
			elements = append(elements, el.Inspect())
		}
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	if vectorLength >= max_display {
		out.WriteString(", ...")
	}

	out.WriteString("]")
	return out.String()
}
func (bv *BaseVector) Print(dtype string) string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf(dtype+"Vector with %d elements\n", len(bv.Elements)))
	out.WriteString(bv.Inspect())
	return out.String()
}

func (bv *BaseVector) Length() int {
	return len(bv.Elements)
}

func (bv *BaseVector) Values(n int) []Object { return bv.Elements }
func (bv *BaseVector) Head(n int) []Object {
	if n > len(bv.Elements) {
		n = len(bv.Elements)
	}
	return bv.Elements[:n]
}

func (bv *BaseVector) Tail(n int) []Object {
	if n > len(bv.Elements) {
		n = len(bv.Elements)
	}
	return bv.Elements[len(bv.Elements)-n:]
}

// append objects to vector, flattening if necessary
func (bv *BaseVector) Append(vals ...Object) Object {
	els := make([]Object, len(bv.Elements))
	copy(els, bv.Elements)

	for _, object := range vals {
		switch obj := object.(type) {
		case IVector:
			els = append(els, obj.Values()...)
		default:
			els = append(els, obj)
		}
	}
	return NewVector(els)
}
func (bv *BaseVector) Slice(start int, end int) Object {
	return NewVector(bv.Elements[start:end])
}

func (bv *BaseVector) Set(index int, val Object, t ObjectType) Object {
	if t != NULL_OBJ && val.Type() != t {
		return NewError("type mismatch for vector, expected %s, got %s", t, val.Type())
	}
	bv.Elements[index] = val
	return NULL
}
func (bv *BaseVector) Replace(start int, end int, replacement IVector, t ObjectType) Object {
	for i := start; i < end; i++ {
		if err := bv.Set(i, replacement.Values()[i-start], t); IsError(err) {
			return err
		}
	}
	return NULL
}

func (bv *BaseVector) Infix(f func(Object, Object) Object, other IVector) Object {
	var otherVal Object
	otherLength := other.Length()
	if len(bv.Elements) != otherLength {
		if otherLength != 1 {
			return NewError("Incompatible vector lengths, left=%d and right=%d", len(bv.Elements), otherLength)
		}
		otherVal = other.Values()[0]
	}

	els := []Object{}
	for i, el := range bv.Elements {
		var val Object
		if otherLength == 1 {
			// recyle length one vector
			val = f(el, otherVal)
		} else {
			val = f(el, other.Values()[i])
		}
		if IsError(val) {
			return val
		}
		els = append(els, val)
	}

	return NewVector(els)
}

// vector with mixed types
type Vector struct {
	BaseVector
}

func (v *Vector) Type() ObjectType         { return VECTOR_OBJ }
func (bv *Vector) ElementType() ObjectType { return NULL_OBJ }
func (v *Vector) Inspect() string          { return v.BaseVector.Inspect() }
func (v *Vector) Print() string            { return v.BaseVector.Print("") }

func (v *Vector) Length() int { return v.BaseVector.Length() }
func (v *Vector) Head(n int) Object {
	return &Vector{BaseVector{v.BaseVector.Head(n)}}
}
func (v *Vector) Tail(n int) Object {
	return &Vector{BaseVector{v.BaseVector.Tail(n)}}
}
func (v *Vector) Values() []Object { return v.BaseVector.Elements }
func (v *Vector) Append(objects ...Object) Object {
	return v.BaseVector.Append(objects...)
}
func (v *Vector) Slice(start int, end int) Object {
	return v.BaseVector.Slice(start, end)
}
func (v *Vector) Set(index int, val Object) Object {
	return v.BaseVector.Set(index, val, v.ElementType())
}
func (v *Vector) Replace(start int, end int, replacement IVector) Object {
	return v.BaseVector.Replace(start, end, replacement, v.ElementType())
}
func (v *Vector) Infix(f func(Object, Object) Object, other IVector) Object {
	return v.BaseVector.Infix(f, other)
}

// vectors with object.Number
type NumericVector struct {
	BaseVector
}

func (nv *NumericVector) Type() ObjectType        { return VECTOR_OBJ }
func (bv *NumericVector) ElementType() ObjectType { return NUMBER_OBJ }
func (nv *NumericVector) Inspect() string         { return nv.BaseVector.Inspect() }
func (nv *NumericVector) Print() string           { return nv.BaseVector.Print("Numeric") }

func (nv *NumericVector) New(elements []Object) IVector {
	return &NumericVector{BaseVector{elements}}
}
func (nv *NumericVector) Length() int { return nv.BaseVector.Length() }
func (nv *NumericVector) Head(n int) Object {
	return &NumericVector{BaseVector{nv.BaseVector.Head(n)}}
}
func (nv *NumericVector) Tail(n int) Object {
	return &NumericVector{BaseVector{nv.BaseVector.Tail(n)}}
}
func (nv *NumericVector) Values() []Object {
	return nv.BaseVector.Elements
}
func (nv *NumericVector) Append(objects ...Object) Object {
	return nv.BaseVector.Append(objects...)
}
func (nv *NumericVector) Slice(start int, end int) Object {
	return nv.BaseVector.Slice(start, end)
}
func (nv *NumericVector) Set(index int, val Object) Object {
	return nv.BaseVector.Set(index, val, nv.ElementType())
}
func (nv *NumericVector) Replace(start int, end int, replacement IVector) Object {
	return nv.BaseVector.Replace(start, end, replacement, nv.ElementType())
}
func (nv *NumericVector) Infix(f func(Object, Object) Object, other IVector) Object {
	return nv.BaseVector.Infix(f, other)
}

// vectors with object.String
type CharacterVector struct {
	BaseVector
}

func (cv *CharacterVector) Type() ObjectType        { return VECTOR_OBJ }
func (bv *CharacterVector) ElementType() ObjectType { return STRING_OBJ }
func (cv *CharacterVector) Inspect() string         { return cv.BaseVector.Inspect() }
func (cv *CharacterVector) Print() string           { return cv.BaseVector.Print("Character") }
func (cv *CharacterVector) New(elements []Object) IVector {
	return &CharacterVector{BaseVector{elements}}
}
func (cv *CharacterVector) Length() int { return cv.BaseVector.Length() }
func (cv *CharacterVector) Head(n int) Object {
	return &CharacterVector{BaseVector{cv.BaseVector.Head(n)}}
}
func (cv *CharacterVector) Tail(n int) Object {
	return &CharacterVector{BaseVector{cv.BaseVector.Tail(n)}}
}
func (cv *CharacterVector) Values() []Object {
	return cv.BaseVector.Elements
}
func (cv *CharacterVector) Append(objects ...Object) Object {
	return cv.BaseVector.Append(objects...)
}
func (cv *CharacterVector) Slice(start int, end int) Object {
	return cv.BaseVector.Slice(start, end)
}
func (cv *CharacterVector) Set(index int, val Object) Object {
	return cv.BaseVector.Set(index, val, cv.ElementType())
}
func (cv *CharacterVector) Replace(start int, end int, replacement IVector) Object {
	return cv.BaseVector.Replace(start, end, replacement, cv.ElementType())
}
func (cv *CharacterVector) Infix(f func(Object, Object) Object, other IVector) Object {
	return cv.BaseVector.Infix(f, other)
}

// vectors with object.Boolean
type LogicalVector struct {
	BaseVector
}

func (lv *LogicalVector) Type() ObjectType        { return VECTOR_OBJ }
func (bv *LogicalVector) ElementType() ObjectType { return BOOLEAN_OBJ }
func (lv *LogicalVector) Inspect() string         { return lv.BaseVector.Inspect() }
func (lv *LogicalVector) Print() string           { return lv.BaseVector.Print("Logical") }
func (lv *LogicalVector) New(elements []Object) IVector {
	return &LogicalVector{BaseVector{elements}}
}
func (lv *LogicalVector) Length() int { return lv.BaseVector.Length() }
func (lv *LogicalVector) Head(n int) Object {
	return &LogicalVector{BaseVector{lv.BaseVector.Head(n)}}
}
func (lv *LogicalVector) Tail(n int) Object {
	return &LogicalVector{BaseVector{lv.BaseVector.Tail(n)}}
}
func (lv *LogicalVector) Values() []Object {
	return lv.BaseVector.Elements
}
func (lv *LogicalVector) Append(objects ...Object) Object {
	return lv.BaseVector.Append(objects...)
}
func (lv *LogicalVector) Slice(start int, end int) Object {
	return lv.BaseVector.Slice(start, end)
}
func (lv *LogicalVector) Set(index int, val Object) Object {
	return lv.BaseVector.Set(index, val, lv.ElementType())
}
func (lv *LogicalVector) Replace(start int, end int, replacement IVector) Object {
	return lv.BaseVector.Replace(start, end, replacement, lv.ElementType())
}
func (lv *LogicalVector) Infix(f func(Object, Object) Object, other IVector) Object {
	return lv.BaseVector.Infix(f, other)
}

// create vector based on type of first element
func NewVector(elements []Object) Object {
	firstElement := elements[0]
	if len(elements) == 1 && IsError(firstElement) {
		return firstElement
	}
	switch firstElement.(type) {
	case *Number:
		if sameType(elements, reflect.TypeOf(firstElement)) {
			return &NumericVector{BaseVector: BaseVector{Elements: elements}}
		}
	case *String:
		if sameType(elements, reflect.TypeOf(firstElement)) {
			return &CharacterVector{BaseVector: BaseVector{Elements: elements}}
		}
	case *Boolean:
		if sameType(elements, reflect.TypeOf(firstElement)) {
			return &LogicalVector{BaseVector: BaseVector{Elements: elements}}
		}
	}
	return &Vector{BaseVector: BaseVector{Elements: elements}}
}

func sameType(values []Object, t reflect.Type) bool {
	if len(values) == 0 {
		return true
	}
	for _, v := range values[1:] {
		if reflect.TypeOf(v) != t {
			return false
		}
	}
	return true
}
