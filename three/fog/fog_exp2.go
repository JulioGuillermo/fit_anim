package fog

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type GoFogExp2 struct {
	Fog js.Value
}

func NewFogExp2(color string, density float64) *GoFog {
	fog := mod.GetThree("Fog").New(color, density)
	return &GoFog{
		Fog: fog,
	}
}

func (f *GoFogExp2) ToScene(scene *scene.GoScene) {
	scene.Scene.Set("fog", f.Fog)
}

func (l *GoFogExp2) GetColor() string {
	return l.Fog.Get("color").Call("getHexString").String()
}

func (l *GoFogExp2) SetColor(color string) {
	l.Fog.Get("color").Call("set", color)
}

func (l *GoFogExp2) GetDensity() float64 {
	return l.Fog.Get("density").Float()
}

func (l *GoFogExp2) SetDensity(n float64) {
	l.Fog.Set("density", n)
}
