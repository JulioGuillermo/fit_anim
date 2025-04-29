package elements

import (
	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/three"
	"github.com/julioguillermo/fit_anim/three/geometry"
	"github.com/julioguillermo/fit_anim/three/materials"
	"github.com/julioguillermo/fit_anim/three/point"
)

type GoFreqBox struct {
	MinFreq int
	MaxFreq int

	Box *three.GoMesh
	Mat *materials.GoPhysicalMaterial

	Pos point.Point
}

func FreqBox(MinFreq, MaxFreq int, x, size float64, color, emissive string) *GoFreqBox {
	geo := geometry.BoxGeometry(size, size, size)

	mat := materials.PhysicalMaterial(color)
	mat.SetEmissive(emissive)
	mat.SetEmissiveIntensity(0)

	box := three.Mesh(geo, mat)
	pos := box.GetPosition()
	pos.X = x
	box.SetPosition(pos)

	return &GoFreqBox{
		MinFreq: MinFreq,
		MaxFreq: MaxFreq,
		Box:     box,
		Mat:     mat,
		Pos:     pos,
	}
}

func (p *GoFreqBox) Update(frames *models.Frames) {
	i := frames.FreqGt(p.MinFreq).FreqLt(p.MaxFreq).Int()
	pos := p.Pos
	pos.Y = i
	p.Box.SetPosition(pos)
	p.Mat.SetEmissiveIntensity(i)
}
