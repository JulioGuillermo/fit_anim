package elements

import (
	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type GoFreqSphs struct {
	Sphs []*GoFreqSph
}

func FreqSphs(x, y, z float64, color, emissive string, scene *scene.GoScene) *GoFreqSphs {
	minFreq := 200
	maxFreq := 3000
	incs := 20
	sphs := []*GoFreqSph{}
	size := 0.2
	space := size*2 + 0.1
	freqs := (maxFreq - minFreq) / incs

	for i := range freqs {
		freq := i * incs
		sph := FreqSph(
			freq,
			x,
			y+float64(freqs-i)*space/2,
			z-float64(freqs-i)*space,
			size,
			color,
			emissive,
		)
		scene.Add(sph.Sph)
		sphs = append(sphs, sph)
	}

	return &GoFreqSphs{
		Sphs: sphs,
	}
}

func (p *GoFreqSphs) Update(frames *models.Frames) {
	for _, b := range p.Sphs {
		b.Update(frames)
	}
}
