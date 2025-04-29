package geometry

import (
	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoBoxGeometry struct {
	GoGeometry
}

func BoxGeometry(x, y, z float64) *GoBoxGeometry {
	geometry := mod.GetThree("BoxGeometry").New(x, y, z)
	return &GoBoxGeometry{
		GoGeometry: GoGeometry{
			Geometry: geometry,
		},
	}
}

func (b *GoBoxGeometry) GetGeometry() *GoGeometry {
	return &b.GoGeometry
}

func (b *GoBoxGeometry) GetX() float64 {
	return b.Geometry.Get("width").Float()
}

func (b *GoBoxGeometry) SetX(x float64) {
	b.Geometry.Set("width", x)
}

func (b *GoBoxGeometry) GetY() float64 {
	return b.Geometry.Get("height").Float()
}

func (b *GoBoxGeometry) SetY(y float64) {
	b.Geometry.Set("height", y)
}

func (b *GoBoxGeometry) GetZ() float64 {
	return b.Geometry.Get("depth").Float()
}

func (b *GoBoxGeometry) SetZ(z float64) {
	b.Geometry.Set("depth", z)
}
