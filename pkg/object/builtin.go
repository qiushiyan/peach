package object

type BuiltinFunction func(args ...Object) Object
type Builtin struct {
	ParametersNum int
	Fn            BuiltinFunction
}

func (b *Builtin) Inspect() string  { return "builtin" }
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
