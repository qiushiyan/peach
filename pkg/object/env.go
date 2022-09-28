package object

func NewEnv() *Env {
	s := make(map[string]Object)
	return &Env{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Env) *Env {
	env := NewEnv()
	env.outer = outer
	return env
}

type Env struct {
	store map[string]Object
	outer *Env
}

func (e *Env) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}
func (e *Env) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
