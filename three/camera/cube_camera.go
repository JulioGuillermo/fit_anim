package camera

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
	"github.com/julioguillermo/fit_anim/three/render"
	"github.com/julioguillermo/fit_anim/three/renderer"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type GoCubeCamera struct {
	Camera js.Value
}

func CubeCamera(near, far float64, target render.RenderTarget) *GoCubeCamera {
	cam := mod.GetThree("CubeCamera").New(near, far, target.GetRender())
	return &GoCubeCamera{
		Camera: cam,
	}
}

func (c *GoCubeCamera) GetNear() float64 {
	return c.Camera.Get("near").Float()
}

func (c *GoCubeCamera) SetNear(n float64) {
	c.Camera.Set("near", n)
}

func (c *GoCubeCamera) GetFar() float64 {
	return c.Camera.Get("far").Float()
}

func (c *GoCubeCamera) SetFar(n float64) {
	c.Camera.Set("far", n)
}

func (c *GoCubeCamera) SetTarget(t render.RenderTarget) {
	c.Camera.Set("renderTarget", t.GetRender())
}

func (c *GoCubeCamera) Update(renderer *renderer.GoRenderer, scene *scene.GoScene) {
	c.Camera.Call("update", renderer.Renderer, scene.Scene)
}
