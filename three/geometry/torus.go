package geometry

import "github.com/julioguillermo/fit_anim/three/mod"

type GoTorusGeometry struct {
	GoGeometry
}

func TorusGeometry(radius, tubRad float64, radSegments, tubSegments int) *GoTorusGeometry {
	geometry := mod.GetThree("TorusGeometry").New(radius, tubRad, radSegments, tubSegments)
	return &GoTorusGeometry{
		GoGeometry: GoGeometry{
			Geometry: geometry,
		},
	}
}

func (t *GoTorusGeometry) GetGeometry() *GoGeometry {
	return &t.GoGeometry
}

func (t *GoTorusGeometry) GetRadius() float64 {
	return t.Geometry.Get("radius").Float()
}

func (t *GoTorusGeometry) SetRadius(r float64) {
	t.Geometry.Set("radius", r)
}

func (t *GoTorusGeometry) GetTubRadius() float64 {
	return t.Geometry.Get("tubeRadius").Float()
}

func (t *GoTorusGeometry) SetTubRadius(r float64) {
	t.Geometry.Set("tubeRadius", r)
}

func (t *GoTorusGeometry) GetRadSegments() int {
	return t.Geometry.Get("radialSegments").Int()
}

func (t *GoTorusGeometry) SetRadSegments(s int) {
	t.Geometry.Set("radialSegments", s)
}

func (t *GoTorusGeometry) GetTubSegments() int {
	return t.Geometry.Get("tubularSegments").Int()
}

func (t *GoTorusGeometry) SetTubSegments(s int) {
	t.Geometry.Set("tubularSegments", s)
}
