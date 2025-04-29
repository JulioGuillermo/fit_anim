package geometry

import "github.com/julioguillermo/fit_anim/three/mod"

type GoCylinderGeometry struct {
	GoGeometry
}

func CylinderGeometry(radTop, radBottom, height float64, segments int) *GoCylinderGeometry {
	geometry := mod.GetThree("CylinderGeometry").New(radTop, radBottom, height, segments)
	return &GoCylinderGeometry{
		GoGeometry: GoGeometry{
			Geometry: geometry,
		},
	}
}

func (c *GoCylinderGeometry) GetGeometry() *GoGeometry {
	return &c.GoGeometry
}

func (c *GoCylinderGeometry) GetRadTop() float64 {
	return c.Geometry.Get("radiusTop").Float()
}

func (c *GoCylinderGeometry) SetRadTop(r float64) {
	c.Geometry.Set("radiusTop", r)
}

func (c *GoCylinderGeometry) GetRadBottom() float64 {
	return c.Geometry.Get("radiusBottom").Float()
}

func (c *GoCylinderGeometry) SetRadBottom(r float64) {
	c.Geometry.Set("radiusBottom", r)
}

func (c *GoCylinderGeometry) GetHeight() float64 {
	return c.Geometry.Get("height").Float()
}

func (c *GoCylinderGeometry) SetHeight(r float64) {
	c.Geometry.Set("height", r)
}

func (c *GoCylinderGeometry) GetSegments() int {
	return c.Geometry.Get("radialSegments").Int()
}

func (c *GoCylinderGeometry) SetSegments(s int) {
	c.Geometry.Set("radialSegments", s)
}
