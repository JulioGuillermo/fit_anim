package render

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoCubeRenderTarget struct {
	Render js.Value
}

func CubeRenderTarget(size int) *GoCubeRenderTarget {
	render := mod.GetThree("WebGLCubeRenderTarget").New(size)
	return &GoCubeRenderTarget{
		Render: render,
	}
}

func (r *GoCubeRenderTarget) GetRender() js.Value {
	return r.Render
}

func (r *GoCubeRenderTarget) GetTexture() js.Value {
	return r.Render.Get("texture")
}
