package elements

import (
	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/three"
	"github.com/julioguillermo/fit_anim/three/geometry"
	"github.com/julioguillermo/fit_anim/three/materials"
	"github.com/julioguillermo/fit_anim/three/scene"
)

const MinFreq = 20
const MaxFreq = 22000

type GoFreqBoxs struct {
	Boxs []*GoFreqBox

	Base *three.GoMesh
	Mat  *materials.GoPhysicalMaterial
}

func FreqBoxs(freqs int, color, emissive string, scene *scene.GoScene) *GoFreqBoxs {
	freqDif := (MaxFreq - MinFreq) / freqs

	boxs := make([]*GoFreqBox, freqs)
	size := 0.2
	space := size + 0.1
	width := space * float64(freqs)

	geo := geometry.BoxGeometry(width, size, size)
	mat := materials.PhysicalMaterial(color)
	base := three.Mesh(geo, mat)
	scene.Add(base)

	for i := range freqs {
		freq := MinFreq + i*freqDif
		offset := (float64(i) - float64(freqs)/2 + 0.5)
		box := FreqBox(freq, freq+freqDif, offset*space, size, color, emissive)
		scene.Add(box.Box)
		base.Add(box.Box)
		boxs[i] = box
	}

	return &GoFreqBoxs{
		Boxs: boxs,
		Base: base,
		Mat:  mat,
	}
}

func (p *GoFreqBoxs) Update(frames *models.Frames) {
	for _, b := range p.Boxs {
		b.Update(frames)
	}
}
