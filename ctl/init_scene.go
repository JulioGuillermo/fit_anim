package ctl

import (
	"github.com/julioguillermo/fit_anim/ctl/elements"
	"github.com/julioguillermo/fit_anim/three"
	"github.com/julioguillermo/fit_anim/three/geometry"
	"github.com/julioguillermo/fit_anim/three/materials"
	"github.com/julioguillermo/fit_anim/three/point"
)

func (c *Controller) initFloor() {
	geometry := geometry.PlaneGeometry(100, 100)

	material := materials.PhysicalMaterial("#111222")

	mesh := three.Mesh(geometry, material)
	mesh.Rotate(point.Point{X: -89}.Rad())
	mesh.Move(point.Point{Y: -5})

	c.Scene.Add(mesh)
}

func (c *Controller) initScene() {
	c.initFloor()

	c.MainElement = elements.CreateMainElement(c.Scene)

	c.FreqBoxs0 = elements.FreqBoxs(10, "#FF0000", "#FF0088", c.Scene)
	c.FreqBoxs0.Base.Move(point.Point{Z: -5, X: -5})

	c.FreqBoxs1 = elements.FreqBoxs(10, "#FF0000", "#FF0088", c.Scene)
	c.FreqBoxs1.Base.Move(point.Point{Z: -5, X: 5})
	c.FreqBoxs1.Base.Rotate(point.Point{Y: 180}.Rad())

	c.FreqSphs0 = elements.FreqSphs(
		-5,
		-3,
		0,
		"#FF0000",
		"#FF0088",
		c.Scene,
	)
	c.FreqSphs1 = elements.FreqSphs(
		5,
		-3,
		0,
		"#FF0000",
		"#FF0088",
		c.Scene,
	)
}
