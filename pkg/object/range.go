package object

import "fmt"

type Range struct {
	Start int
	End   int
}

func (r *Range) Type() ObjectType { return RANGE_OBJ }
func (r *Range) Inspect() string  { return fmt.Sprintf("%d:%d", r.Start, r.End) }
