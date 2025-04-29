package geometry

import (
	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoOctahedronGeometry struct {
	GoGeometry
}

func OctahedronGeometry(radious float64, details int) *GoOctahedronGeometry {
	geo := mod.GetThree("OctahedronGeometry").New(radious, details)
	return &GoOctahedronGeometry{
		GoGeometry: GoGeometry{
			Geometry: geo,
		},
	}
}

func (o *GoOctahedronGeometry) GetGeometry() *GoGeometry {
	return &o.GoGeometry
}

func (o *GoOctahedronGeometry) GetRadius() float64 {
	return o.Geometry.Get("radius").Float()
}

func (o *GoOctahedronGeometry) SetRadius(r float64) {
	o.Geometry.Set("radius", r)
}

func (o *GoOctahedronGeometry) GetDetails() int {
	return o.Geometry.Get("detail").Int()
}

func (o *GoOctahedronGeometry) SetDetails(d int) {
	o.Geometry.Set("detail", d)
}
