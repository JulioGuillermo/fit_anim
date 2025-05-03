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

const (
	SphIntMax = 0.5
	SphIntMin = 0.05
	SphInc    = 0.5
	SphDec    = 0.3
	TorusInc  = 0.5
	TorusDec  = 0.3
)

type MainElement struct {
	Scene *scene.GoScene

	SphereMaterial *materials.GoPhysicalMaterial
	Sphere         *three.GoMesh

	TorusMaterial *materials.GoPhysicalMaterial
	Torus         *three.GoMesh
}

func CreateMainElement(scene *scene.GoScene) *MainElement {
	element := &MainElement{}
	element.Scene = scene
	element.Init()
	return element
}

func (p *MainElement) initMainSph() {
	geometry := geometry.IcosahedronGeometry(0.1, 10)
	material := materials.PhysicalMaterial("#555588")
	material.SetEmissive("#5555FF")
	material.SetEmissiveIntensity(0)

	mesh := three.Mesh(geometry, material)
	mesh.Move(point.Point{Z: -3, Y: 0.9})

	p.Scene.Add(mesh)
	p.Sphere = mesh
	p.SphereMaterial = material
}

func (p *MainElement) initMainSphTorus() {
	geometry := geometry.TorusKnotGeometry(0.2, 0.01, 100, 10, 3, 4)
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

func (p *MainElement) Init() {
	p.initMainSph()
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

	avgInt := frames.AvgIntencities() * 100
	maxInt := frames.MaxIntencity()
	p.SphereMaterial.SetEmissiveIntensity(avgInt)
	p.TorusMaterial.SetEmissiveIntensity(maxInt)
}
