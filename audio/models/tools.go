package models

import "syscall/js"

func GetNumFunc(fun func(p float64) any) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 {
			return this
		}
		p := args[0].Float()
		return fun(p)
	})
}
