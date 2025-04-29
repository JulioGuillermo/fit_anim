package mod

import "syscall/js"

func GetThree(mod string) js.Value {
	return js.Global().Get("THREE").Get(mod)
}
