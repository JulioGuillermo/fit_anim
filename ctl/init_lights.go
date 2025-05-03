package ctl

import (
	"github.com/julioguillermo/fit_anim/three/light"
	"github.com/julioguillermo/fit_anim/three/point"
)

func (c *Controller) initLights() {
	dirLight := light.DirectionalLight("#FFFFFF", 0.2)
	dirLight.Move(point.Point{Z: 1, Y: 2})
	dirLight.ToScene(c.Scene)
	c.TopLight = dirLight

	// pointLight := light.PointLight("#FFFFFF", 5)
	// pointLight.Move(point.Point{Z: 2, Y: 1})
	// c.Scene.Add(pointLight)
}
