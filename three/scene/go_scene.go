package scene

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/texture"
)

type SceneObject interface {
	GetObject() js.Value
}

type GoScene struct {
	Scene js.Value
}

func (p *GoScene) Add(o SceneObject) {
	p.AddJS(o.GetObject())
}

func (p *GoScene) AddJS(o js.Value) {
	p.Scene.Call("add", o)
}

func (p *GoScene) SetBG(texture *texture.GoTexture) {
	p.Scene.Set("background", texture.Texture)
}

func (p *GoScene) GetJSObjByName(name string) js.Value {
	return p.Scene.Call("getObjectByName", name, true)
}
