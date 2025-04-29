package texture

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
)

type Texture interface {
	GetTexture() js.Value
}

type GoTexture struct {
	Texture js.Value
}

func LoadTexture(files ...string) *GoTexture {
	if len(files) == 0 {
		return nil
	}

	var texture js.Value
	loader := mod.GetThree("TextureLoader").New()

	if len(files) == 1 {
		texture = loader.Call("load", files[0])
	} else {
		texture = loader.Call("load", files)
	}

	return &GoTexture{
		Texture: texture,
	}
}

func (t *GoTexture) GetTexture() js.Value {
	return t.Texture
}
