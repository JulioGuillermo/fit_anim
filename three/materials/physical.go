package materials

import (
	"github.com/julioguillermo/fit_anim/three/mod"
	"github.com/julioguillermo/fit_anim/three/texture"
)

type GoPhysicalMaterial struct {
	GoStandardMaterial
}

func PhysicalMaterial(color string) *GoPhysicalMaterial {
	mat := mod.GetThree("MeshPhysicalMaterial").New(map[string]any{
		"color": color,
	})
	return &GoPhysicalMaterial{
		GoStandardMaterial: GoStandardMaterial{
			GoBasicMaterial: GoBasicMaterial{
				Material: mat,
			},
		},
	}
}

func (m *GoPhysicalMaterial) GetClearcoat() float64 {
	return m.Material.Get("clearcoat").Float()
}

func (m *GoPhysicalMaterial) SetClearcoat(s float64) {
	m.Material.Set("clearcoat", max(0, min(1, s)))
}

func (m *GoPhysicalMaterial) GetClearCoatRoughness() float64 {
	return m.Material.Get("clearCoatRoughness").Float()
}

func (m *GoPhysicalMaterial) SetClearCoatRoughness(s float64) {
	m.Material.Set("clearCoatRoughness", max(0, min(1, s)))
}

func (m *GoPhysicalMaterial) GetReflectivity() float64 {
	return m.Material.Get("reflectivity").Float()
}

func (m *GoPhysicalMaterial) SetReflectivity(r float64) {
	m.Material.Set("reflectivity", r)
}

func (m *GoPhysicalMaterial) GetTransmission() float64 {
	return m.Material.Get("transmission").Float()
}

func (m *GoPhysicalMaterial) SetTransmission(t float64) {
	m.Material.Set("transmission", t)
}

func (m *GoPhysicalMaterial) GetThickness() float64 {
	return m.Material.Get("thickness").Float()
}

func (m *GoPhysicalMaterial) SetThickness(t float64) {
	m.Material.Set("thickness", t)
}

func (m *GoPhysicalMaterial) SetEnvMap(t texture.Texture) {
	m.Material.Set("envMap", t.GetTexture())
}

func (m *GoPhysicalMaterial) SetNormalMap(t texture.Texture) {
	m.Material.Set("normalMap", t.GetTexture())
}
