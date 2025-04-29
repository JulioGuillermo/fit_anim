package materials

import "syscall/js"

type GoMaterial interface {
	GetMaterial() js.Value
}
