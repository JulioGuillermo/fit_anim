package light

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
	"github.com/julioguillermo/fit_anim/three/point"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type GoDirectionalLight struct {
	GoAmbientLight
}

func DirectionalLight(color string, intencity float64) *GoDirectionalLight {
	light := mod.GetThree("DirectionalLight").New(color, intencity)
	return &GoDirectionalLight{
		GoAmbientLight: GoAmbientLight{
			Light: light,
		},
	}
}

func (l *GoDirectionalLight) GetObject() js.Value {
	return l.Light
}

func (l *GoDirectionalLight) GetTarget() js.Value {
	return l.Light.Get("target")
}

func (l *GoDirectionalLight) ToScene(scene *scene.GoScene) {
	scene.AddJS(l.GetObject())
	scene.AddJS(l.GetTarget())
}

func (l *GoDirectionalLight) GetPosition() point.Point {
	pos := l.Light.Get("position")
	return point.Point{
		X: pos.Get("x").Float(),
		Y: pos.Get("y").Float(),
		Z: pos.Get("z").Float(),
	}
}

func (l *GoDirectionalLight) SetPosition(pos point.Point) {
	p := l.Light.Get("position")
	p.Set("x", pos.X)
	p.Set("y", pos.Y)
	p.Set("z", pos.Z)
}

func (l *GoDirectionalLight) GetTargetPosition() point.Point {
	pos := l.GetTarget().Get("position")
	return point.Point{
		X: pos.Get("x").Float(),
		Y: pos.Get("y").Float(),
		Z: pos.Get("z").Float(),
	}
}

func (l *GoDirectionalLight) SetTargetPosition(pos point.Point) {
	p := l.GetTarget().Get("position")
	p.Set("x", pos.X)
	p.Set("y", pos.Y)
	p.Set("z", pos.Z)
}

func (l *GoDirectionalLight) Move(to point.Point) {
	pos := l.GetPosition()
	pos.X += to.X
	pos.Y += to.Y
	pos.Z += to.Z
	l.SetPosition(pos)
}

func (l *GoDirectionalLight) MoveTarget(to point.Point) {
	pos := l.GetTargetPosition()
	pos.X += to.X
	pos.Y += to.Y
	pos.Z += to.Z
	l.SetTargetPosition(pos)
}
