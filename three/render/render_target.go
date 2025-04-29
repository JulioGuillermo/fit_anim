package render

import "syscall/js"

type RenderTarget interface {
	GetRender() js.Value
	GetTexture() js.Value
}
