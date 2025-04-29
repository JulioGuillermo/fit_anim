package materials

import (
	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoPhongMaterial struct {
	GoBasicMaterial
}

func PhongMaterial(color string) *GoPhongMaterial {
	mat := mod.GetThree("MeshPhongMaterial").New(map[string]any{
		"color": color,
	})
	return &GoPhongMaterial{
		GoBasicMaterial: GoBasicMaterial{
			Material: mat,
		},
	}
}

func (m *GoPhongMaterial) GetEmissive() string {
	return m.Material.Get("emissive").Call("getHexString").String()
}

func (m *GoPhongMaterial) SetEmissive(color string) {
	m.Material.Get("emissive").Call("set", color)
}

func (m *GoPhongMaterial) GetEmissiveIntensity() float64 {
	return m.Material.Get("emissiveIntensity").Float()
}

func (m *GoPhongMaterial) SetEmissiveIntensity(i float64) {
	m.Material.Set("emissiveIntensity", i)
}

func (m *GoPhongMaterial) GetShininess() float64 {
	return m.Material.Get("shininess").Float()
}

func (m *GoPhongMaterial) SetShininess(s float64) {
	m.Material.Set("shininess", s)
}
