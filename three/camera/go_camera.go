package camera

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/point"
)

type GoCamera struct {
	Camera js.Value
}

func (c *GoCamera) GetObject() js.Value {
	return c.Camera
}

func (c *GoCamera) GetPosition() point.Point {
	pos := c.Camera.Get("position")
	return point.Point{
		X: pos.Get("x").Float(),
		Y: pos.Get("y").Float(),
		Z: pos.Get("z").Float(),
	}
}

func (c *GoCamera) GetRotation() point.Point {
	rot := c.Camera.Get("rotation")
	return point.Point{
		X: rot.Get("x").Float(),
		Y: rot.Get("y").Float(),
		Z: rot.Get("z").Float(),
	}
}

func (c *GoCamera) GetScale() point.Point {
	scl := c.Camera.Get("scale")
	return point.Point{
		X: scl.Get("x").Float(),
		Y: scl.Get("y").Float(),
		Z: scl.Get("z").Float(),
	}
}

func (c *GoCamera) SetPosition(pos point.Point) {
	p := c.Camera.Get("position")
	p.Set("x", pos.X)
	p.Set("y", pos.Y)
	p.Set("z", pos.Z)
}

func (c *GoCamera) SetRotation(rot point.Point) {
	r := c.Camera.Get("rotation")
	r.Set("x", rot.X)
	r.Set("y", rot.Y)
	r.Set("z", rot.Z)
}

func (c *GoCamera) SetScale(scl point.Point) {
	s := c.Camera.Get("scale")
	s.Set("x", scl.X)
	s.Set("y", scl.Y)
	s.Set("z", scl.Z)
}

func (c *GoCamera) Move(to point.Point) {
	pos := c.GetPosition()
	pos.X += to.X
	pos.Y += to.Y
	pos.Z += to.Z
	c.SetPosition(pos)
}

func (c *GoCamera) Rotate(to point.Point) {
	rot := c.GetRotation()
	rot.X += to.X
	rot.Y += to.Y
	rot.Z += to.Z
	c.SetRotation(rot)
}

func (c *GoCamera) Scale(to point.Point) {
	scl := c.GetScale()
	scl.X += to.X
	scl.Y += to.Y
	scl.Z += to.Z
	c.SetScale(scl)
}
