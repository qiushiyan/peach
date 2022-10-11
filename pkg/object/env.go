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
	// if not found in current env, recursively search outer env
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *Env) Set(name string, val Object) Object {
	// env.Set only set value in the current env
	// to give for and if blocks a chance to set outer variables using assign expressions see evalAssignExpression
	e.store[name] = val
	return val
}

func (e *Env) GetOuterEnv() *Env {
	return e.outer
}

func (e *Env) InStore(name string) bool {
	_, ok := e.store[name]
	return ok
}

func (e *Env) GetAllEnvs() []*Env {
	envs := []*Env{e}
	for e.outer != nil {
		envs = append(envs, e.outer)
		e = e.outer
	}
	return envs
}
