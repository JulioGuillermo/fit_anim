package ctl

import (
	"github.com/julioguillermo/fit_anim/three/point"
	"github.com/julioguillermo/fit_anim/three/texture"
)

func (c *Controller) initCamera() {
	c.Camera.Move(point.Point{Z: 10})

	bg := texture.LoadTexture("/img/bg.avif")
	c.Scene.SetBG(bg)
}
