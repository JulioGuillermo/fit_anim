package geometry

import "github.com/julioguillermo/fit_anim/three/mod"

type GoTorusKnotGeometry struct {
	GoTorusGeometry
}

func TorusKnotGeometry(radius, tubRad float64, radSegments, tubSegments, p, q int) *GoTorusKnotGeometry {
	geometry := mod.GetThree("TorusKnotGeometry").New(radius, tubRad, radSegments, tubSegments, p, q)
	return &GoTorusKnotGeometry{
		GoTorusGeometry: GoTorusGeometry{
			GoGeometry: GoGeometry{
				Geometry: geometry,
			},
		},
	}
}

func (t *GoTorusKnotGeometry) GetP() int {
	return t.Geometry.Get("p").Int()
}

func (t *GoTorusKnotGeometry) SetP(p int) {
	t.Geometry.Set("p", p)
}

func (t *GoTorusKnotGeometry) GetQ() int {
	return t.Geometry.Get("q").Int()
}

func (t *GoTorusKnotGeometry) SetQ(q int) {
	t.Geometry.Set("q", q)
}
