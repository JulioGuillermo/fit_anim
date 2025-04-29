package three

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/geometry"
	"github.com/julioguillermo/fit_anim/three/materials"
	"github.com/julioguillermo/fit_anim/three/mod"
	"github.com/julioguillermo/fit_anim/three/point"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type GoMesh struct {
	Mesh js.Value
}

func Mesh(g geometry.Geometry, mat materials.GoMaterial) *GoMesh {
	mesh := mod.GetThree("Mesh").New(g.GetGeometry().Geometry, mat.GetMaterial())
	return &GoMesh{
		Mesh: mesh,
	}
}

func (m *GoMesh) GetObject() js.Value {
	return m.Mesh
}

func (m *GoMesh) Add(obj scene.SceneObject) {
	m.AddJS(obj.GetObject())
}

func (m *GoMesh) AddJS(obj js.Value) {
	m.Mesh.Call("add", obj)
}

func (m *GoMesh) GetPosition() point.Point {
	pos := m.Mesh.Get("position")
	return point.Point{
		X: pos.Get("x").Float(),
		Y: pos.Get("y").Float(),
		Z: pos.Get("z").Float(),
	}
}

func (m *GoMesh) GetRotation() point.Point {
	rot := m.Mesh.Get("rotation")
	return point.Point{
		X: rot.Get("x").Float(),
		Y: rot.Get("y").Float(),
		Z: rot.Get("z").Float(),
	}
}

func (m *GoMesh) GetScale() point.Point {
	scl := m.Mesh.Get("scale")
	return point.Point{
		X: scl.Get("x").Float(),
		Y: scl.Get("y").Float(),
		Z: scl.Get("z").Float(),
	}
}

func (m *GoMesh) SetPosition(pos point.Point) {
	p := m.Mesh.Get("position")
	p.Set("x", pos.X)
	p.Set("y", pos.Y)
	p.Set("z", pos.Z)
}

func (m *GoMesh) SetRotation(rot point.Point) {
	r := m.Mesh.Get("rotation")
	r.Set("x", rot.X)
	r.Set("y", rot.Y)
	r.Set("z", rot.Z)
}

func (m *GoMesh) SetScale(scl point.Point) {
	s := m.Mesh.Get("scale")
	s.Set("x", scl.X)
	s.Set("y", scl.Y)
	s.Set("z", scl.Z)
}

func (m *GoMesh) Move(to point.Point) {
	pos := m.GetPosition()
	pos.X += to.X
	pos.Y += to.Y
	pos.Z += to.Z
	m.SetPosition(pos)
}

func (m *GoMesh) Rotate(to point.Point) {
	rot := m.GetRotation()
	rot.X += to.X
	rot.Y += to.Y
	rot.Z += to.Z
	m.SetRotation(rot)
}

func (m *GoMesh) Scale(to point.Point) {
	scl := m.GetScale()
	scl.X += to.X
	scl.Y += to.Y
	scl.Z += to.Z
	m.SetScale(scl)
}
