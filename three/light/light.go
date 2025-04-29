package light

import "syscall/js"

type Light interface {
	GetObject() js.Value
	GetColor() string
	SetColor(color string)
	GetIntencity() float64
	SetIntencity(i float64)
}
