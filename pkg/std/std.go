package std

import "github.com/qiushiyan/qlang/pkg/object"

// The built-in functions / standard-library methods are stored here.
var Builtins = map[string]*object.Builtin{}

func RegisterBuiltin(name string, fun object.BuiltinFunction, paramsNum int) {
	Builtins[name] = &object.Builtin{Fn: fun, RequiredParametersNum: paramsNum}
}

func RegisterStd() {
	RegisterBuiltin("len", length, 1)
	RegisterBuiltin("print", print, 1)
	RegisterBuiltin("head", vectorHead, 1)
	RegisterBuiltin("tail", vectorTail, 1)
	RegisterBuiltin("append", vectorAppend, 2)
	RegisterBuiltin("keys", dictKeys, 1)
	RegisterBuiltin("values", dictValues, 1)
	RegisterBuiltin("random", random, 0)
	RegisterBuiltin("vector", vectorCreate, 1)
	RegisterBuiltin("to_vector", vectorConvert, 1)
}
