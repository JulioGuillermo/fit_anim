package geometry

import (
	"github.com/julioguillermo/fit_anim/three/mod"
)

type GoIcosahedronGeometry struct {
	GoGeometry
}

func IcosahedronGeometry(radious float64, details int) *GoIcosahedronGeometry {
	geo := mod.GetThree("IcosahedronGeometry").New(radious, details)
	return &GoIcosahedronGeometry{
		GoGeometry: GoGeometry{
			Geometry: geo,
		},
	}
}

func (i *GoIcosahedronGeometry) GetGeometry() *GoGeometry {
	return &i.GoGeometry
}

func (i *GoIcosahedronGeometry) GetRadius() float64 {
	return i.Geometry.Get("radius").Float()
}

func (i *GoIcosahedronGeometry) SetRadius(r float64) {
	i.Geometry.Set("radius", r)
}

func (i *GoIcosahedronGeometry) GetDetails() int {
	return i.Geometry.Get("detail").Int()
}

func (i *GoIcosahedronGeometry) SetDetails(d int) {
	i.Geometry.Set("detail", d)
}
