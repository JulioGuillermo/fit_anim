package fog

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type GoFog struct {
	Fog js.Value
}

func NewFog(color string, near, far float64) *GoFog {
	fog := mod.GetThree("Fog").New(color, near, far)
	return &GoFog{
		Fog: fog,
	}
}

func (f *GoFog) ToScene(scene *scene.GoScene) {
	scene.Scene.Set("fog", f.Fog)
}

func (l *GoFog) GetColor() string {
	return l.Fog.Get("color").Call("getHexString").String()
}

func (l *GoFog) SetColor(color string) {
	l.Fog.Get("color").Call("set", color)
}

func (l *GoFog) GetNear() float64 {
	return l.Fog.Get("near").Float()
}

func (l *GoFog) SetNear(n float64) {
	l.Fog.Set("near", n)
}

func (l *GoFog) GetFar() float64 {
	return l.Fog.Get("far").Float()
}

func (l *GoFog) SetFar(n float64) {
	l.Fog.Set("far", n)
}
