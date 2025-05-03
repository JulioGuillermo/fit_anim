package ctl

import (
	"github.com/julioguillermo/fit_anim/three/point"
	"github.com/julioguillermo/fit_anim/three/texture"
)

func (c *Controller) initCamera() {
	c.Camera.Move(point.Point{Z: 2.4, Y: 0.7})

	// bg := texture.LoadTexture("/img/bg.avif")
	bg := texture.LoadTexture("/fit_anim/img/bg_planet_big.png")
	c.Scene.SetBG(bg)
}
