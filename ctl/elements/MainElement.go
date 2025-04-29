package elements

import (
	"time"

	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/three"
	"github.com/julioguillermo/fit_anim/three/camera"
	"github.com/julioguillermo/fit_anim/three/geometry"
	"github.com/julioguillermo/fit_anim/three/materials"
	"github.com/julioguillermo/fit_anim/three/point"
	"github.com/julioguillermo/fit_anim/three/renderer"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type MainElement struct {
	Scene *scene.GoScene

	SphereMaterial *materials.GoPhysicalMaterial
	Sphere         *three.GoMesh

	TorusMaterial *materials.GoPhysicalMaterial
	Torus         *three.GoMesh

	SupportMaterial *materials.GoPhysicalMaterial
	Support         *three.GoMesh
}

func CreateMainElement(scene *scene.GoScene) *MainElement {
	element := &MainElement{}
	element.Scene = scene
	element.Init()
	return element
}

func (p *MainElement) initMainSph() {
	geometry := geometry.IcosahedronGeometry(0.5, 10)
	material := materials.PhysicalMaterial("#555588")
	material.SetEmissive("#5555FF")
	material.SetEmissiveIntensity(0)

	mesh := three.Mesh(geometry, material)
	mesh.Move(point.Point{Y: 5})

	p.Scene.Add(mesh)
	p.Sphere = mesh
	p.SphereMaterial = material
}

func (p *MainElement) initMainSphTorus() {
	geometry := geometry.TorusKnotGeometry(2, 0.1, 100, 10, 3, 4)
	material := materials.PhysicalMaterial("#555588")
	material.SetEmissive("#FF00FF")
	material.SetEmissiveIntensity(0)

	mesh := three.Mesh(geometry, material)
	mesh.Rotate(point.Point{X: 90}.Rad())

	p.Scene.Add(mesh)
	p.Sphere.Add(mesh)
	p.TorusMaterial = material
	p.Torus = mesh
}

func (p *MainElement) initMainSphSupport() {
	geometry := geometry.CylinderGeometry(0.1, 0.3, 10, 8)
	material := materials.PhysicalMaterial("#555588")
	material.SetEmissive("#00FF00")
	material.SetEmissiveIntensity(0)

	mesh := three.Mesh(geometry, material)

	p.Scene.Add(mesh)
	p.SupportMaterial = material
	p.Support = mesh
}

func (p *MainElement) Init() {
	p.initMainSph()
	p.initMainSphSupport()
	p.initMainSphTorus()
}

func (p *MainElement) RenderScene(
	scene *scene.GoScene,
	camera *camera.GoCamera,
	renderer *renderer.GoRenderer,
	start time.Time,
	last time.Time,
	delta float64,
	frames *models.Frames,
) {
	p.Sphere.Rotate(point.Point{Y: delta})

	l := frames.FreqEq(240).Int()
	sphE := p.SphereMaterial.GetEmissiveIntensity()
	if l > 0.01 {
		p.SphereMaterial.SetEmissiveIntensity(sphE*0.5 + 0.5)
	} else {
		p.SphereMaterial.SetEmissiveIntensity(sphE * 0.9)
	}

	h := frames.FreqGt(10000).Int()
	p.TorusMaterial.SetEmissiveIntensity(h * h)

	m := frames.FreqGt(500).FreqLt(5000).Int()
	if m > 0.6 {
		p.SupportMaterial.SetEmissiveIntensity(1)
	} else {
		p.SupportMaterial.SetEmissiveIntensity(0)
	}
}
