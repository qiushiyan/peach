package main

import (
	"syscall/js"

	"github.com/qiushiyan/qlang/pkg/object"
	"github.com/qiushiyan/qlang/pkg/std"
	"github.com/qiushiyan/qlang/pkg/web"
)

func main() {
	c := make(chan struct{}, 0)

	std.Register()

	globalEnv := object.NewEnv()

	run := func(this js.Value, p []js.Value) interface{} {
		return eval(p[0].String(), globalEnv)
	}

	runScript := func(this js.Value, p []js.Value) interface{} {
		env := object.NewEnv()
		return eval(p[0].String(), env)
	}

	js.Global().Set("run", js.FuncOf(run))
	js.Global().Set("runScript", js.FuncOf(runScript))

	<-c
}

func eval(input string, env *object.Env) js.Value {
	result, error := web.Evaluate(input, env)
	obj := js.Global().Get("Object").New()
	if error != nil {
		obj.Set("error", js.ValueOf(error.Error()))
		obj.Set("result", js.Null())
	} else {
		obj.Set("error", js.Null())
		if result.Inspect() == "null" {
			// when the evaluation outcome is null
			// return {error: null, result: null} in js
			obj.Set("result", js.Null())
		} else {
			obj.Set("result", js.ValueOf(result.Inspect()))
		}
	}

	return obj
}
