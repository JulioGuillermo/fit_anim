package geometry

import (
	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoPlaneGeometry struct {
	GoGeometry
}

func PlaneGeometry(x, y float64) *GoPlaneGeometry {
	geometry := mod.GetThree("PlaneGeometry").New(x, y)
	return &GoPlaneGeometry{
		GoGeometry: GoGeometry{
			Geometry: geometry,
		},
	}
}

func (b *GoPlaneGeometry) GetGeometry() *GoGeometry {
	return &b.GoGeometry
}

func (b *GoPlaneGeometry) GetX() float64 {
	return b.Geometry.Get("width").Float()
}

func (b *GoPlaneGeometry) SetX(x float64) {
	b.Geometry.Set("width", x)
}

func (b *GoPlaneGeometry) GetY() float64 {
	return b.Geometry.Get("height").Float()
}

func (b *GoPlaneGeometry) SetY(y float64) {
	b.Geometry.Set("height", y)
}
