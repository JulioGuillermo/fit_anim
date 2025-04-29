package three

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/julioguillermo/fit_anim/three/camera"
	"github.com/julioguillermo/fit_anim/three/renderer"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type GoThree struct {
	Start time.Time
	Last  time.Time

	Scene    *scene.GoScene
	Camera   *camera.GoCamera
	Renderer *renderer.GoRenderer

	InitFunc   InitFunc
	RenderFunc RenderFunc

	Audio js.Value
}

func (p *GoThree) InitThree(jsScene, jsCamera, jsRenderer js.Value) {
	fmt.Println("Init GO Three")
	p.Scene = &scene.GoScene{
		Scene: jsScene,
	}
	p.Camera = &camera.GoCamera{
		Camera: jsCamera,
	}
	p.Renderer = &renderer.GoRenderer{
		Renderer: jsRenderer,
	}

	p.Start = time.Now()
	p.Last = p.Start
	p.Audio = js.Global().Get("document").Call("getElementById", "audio")

	if p.InitFunc != nil {
		p.InitFunc(p.Scene, p.Camera, p.Renderer, p.Start)
	}
}

func (p *GoThree) Render(scene, camera, renderer js.Value) {
	if p.RenderFunc != nil {
		audioTime := p.Audio.Get("currentTime").Float()
		p.RenderFunc(p.Scene, p.Camera, p.Renderer, p.Start, p.Last, time.Since(p.Last).Seconds(), audioTime)
		p.Last = time.Now()
	}
}
