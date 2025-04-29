package elements

import (
	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/three"
	"github.com/julioguillermo/fit_anim/three/geometry"
	"github.com/julioguillermo/fit_anim/three/materials"
)

type GoFreqSph struct {
	Freq int

	Sph *three.GoMesh
	Mat *materials.GoPhysicalMaterial
}

func FreqSph(Freq int, x, y, z, size float64, color, emissive string) *GoFreqSph {
	geo := geometry.IcosahedronGeometry(size, 20)

	mat := materials.PhysicalMaterial(color)
	mat.SetEmissive(emissive)
	mat.SetEmissiveIntensity(0)

	sph := three.Mesh(geo, mat)
	pos := sph.GetPosition()
	pos.X = x
	pos.Y = y
	pos.Z = z
	sph.SetPosition(pos)

	return &GoFreqSph{
		Freq: Freq,
		Sph:  sph,
		Mat:  mat,
	}
}

func (p *GoFreqSph) Update(frames *models.Frames) {
	i := frames.FreqEq(p.Freq).Int()
	fi := p.Mat.GetEmissiveIntensity()
	if i > 0.03 {
		p.Mat.SetEmissiveIntensity(fi*0.5 + 0.5)
	} else if i < 0.01 {
		p.Mat.SetEmissiveIntensity(fi * 0.9)
	}
}
