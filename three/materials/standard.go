package materials

import (
	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoStandardMaterial struct {
	GoBasicMaterial
}

func StandardMaterial(color string) *GoStandardMaterial {
	mat := mod.GetThree("MeshStandardMaterial").New(map[string]any{
		"color": color,
	})
	return &GoStandardMaterial{
		GoBasicMaterial: GoBasicMaterial{
			Material: mat,
		},
	}
}

func (m *GoStandardMaterial) GetEmissive() string {
	return m.Material.Get("emissive").Call("getHexString").String()
}

func (m *GoStandardMaterial) SetEmissive(color string) {
	m.Material.Get("emissive").Call("set", color)
}

func (m *GoStandardMaterial) GetEmissiveIntensity() float64 {
	return m.Material.Get("emissiveIntensity").Float()
}

func (m *GoStandardMaterial) SetEmissiveIntensity(i float64) {
	m.Material.Set("emissiveIntensity", i)
}

func (m *GoStandardMaterial) GetRoughness() float64 {
	return m.Material.Get("roughness").Float()
}

func (m *GoStandardMaterial) SetRoughness(s float64) {
	m.Material.Set("roughness", max(0, min(1, s)))
}

func (m *GoStandardMaterial) GetMetalness() float64 {
	return m.Material.Get("metalness").Float()
}

func (m *GoStandardMaterial) SetMetalness(s float64) {
	m.Material.Set("metalness", max(0, min(1, s)))
}
