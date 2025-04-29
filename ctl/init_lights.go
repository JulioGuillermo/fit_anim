package ctl

import (
	"github.com/julioguillermo/fit_anim/three/light"
	"github.com/julioguillermo/fit_anim/three/point"
)

func (c *Controller) initLights() {
	dirLight := light.DirectionalLight("#FFFFFF", 1)
	dirLight.Move(point.Point{Z: 2, Y: 2, X: -2})
	dirLight.ToScene(c.Scene)

	pointLight := light.PointLight("#FFFFFF", 5)
	pointLight.Move(point.Point{Z: 2, Y: -2, X: -2})
	c.Scene.Add(pointLight)
}
