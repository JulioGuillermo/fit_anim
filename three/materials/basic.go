package materials

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoBasicMaterial struct {
	Material js.Value
}

func BasicMaterial(color string) *GoBasicMaterial {
	mat := mod.GetThree("MeshBasicMaterial").New(map[string]any{
		"color": color,
	})
	return &GoBasicMaterial{
		Material: mat,
	}
}

func (m *GoBasicMaterial) GetMaterial() js.Value {
	return m.Material
}

func (m *GoBasicMaterial) GetColor() string {
	return m.Material.Get("color").Call("getHexString").String()
}

func (m *GoBasicMaterial) SetColor(color string) {
	m.Material.Get("color").Call("set", color)
}
