package three

import (
	"syscall/js"
	"time"

	"github.com/julioguillermo/fit_anim/three/camera"
	"github.com/julioguillermo/fit_anim/three/renderer"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type InitFunc func(*scene.GoScene, *camera.GoCamera, *renderer.GoRenderer, time.Time)
type RenderFunc func(*scene.GoScene, *camera.GoCamera, *renderer.GoRenderer, time.Time, time.Time, float64, float64)

func InitGoThree(initFunc InitFunc, renderFunc RenderFunc) *GoThree {
	goThree := &GoThree{
		InitFunc:   initFunc,
		RenderFunc: renderFunc,
	}

	js.Global().Set("goInitScene", js.FuncOf(func(this js.Value, args []js.Value) any {
		goThree.InitThree(args[0], args[1], args[2])
		return nil
	}))
	js.Global().Set("goRender", js.FuncOf(func(this js.Value, args []js.Value) any {
		goThree.Render(args[0], args[1], args[2])
		return nil
	}))

	return goThree
}
