package light

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoAmbientLight struct {
	Light js.Value
}

func AmbientLight(color string, intencity float64) *GoAmbientLight {
	light := mod.GetThree("AmbientLight").New(color, intencity)

	return &GoAmbientLight{
		Light: light,
	}
}

func (l *GoAmbientLight) GetObject() js.Value {
	return l.Light
}

func (l *GoAmbientLight) GetColor() string {
	return l.Light.Get("color").Call("getHexString").String()
}

func (l *GoAmbientLight) SetColor(color string) {
	l.Light.Get("color").Call("set", color)
}

func (l *GoAmbientLight) GetIntencity() float64 {
	return l.Light.Get("intencity").Float()
}

func (l *GoAmbientLight) SetIntencity(i float64) {
	l.Light.Set("intencity", i)
}
