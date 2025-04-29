package light

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
	"github.com/julioguillermo/fit_anim/three/point"
)

type GoPointLight struct {
	GoAmbientLight
}

func PointLight(color string, intencity float64) *GoPointLight {
	light := mod.GetThree("PointLight").New(color, intencity)
	return &GoPointLight{
		GoAmbientLight: GoAmbientLight{
			Light: light,
		},
	}
}

func (l *GoPointLight) GetObject() js.Value {
	return l.Light
}

func (l *GoPointLight) GetPosition() point.Point {
	pos := l.Light.Get("position")
	return point.Point{
		X: pos.Get("x").Float(),
		Y: pos.Get("y").Float(),
		Z: pos.Get("z").Float(),
	}
}

func (l *GoPointLight) SetPosition(pos point.Point) {
	p := l.Light.Get("position")
	p.Set("x", pos.X)
	p.Set("y", pos.Y)
	p.Set("z", pos.Z)
}

func (l *GoPointLight) Move(to point.Point) {
	pos := l.GetPosition()
	pos.X += to.X
	pos.Y += to.Y
	pos.Z += to.Z
	l.SetPosition(pos)
}
