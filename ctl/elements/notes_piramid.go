package elements

import (
	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/audio/notes"
	"github.com/julioguillermo/fit_anim/three"
	"github.com/julioguillermo/fit_anim/three/geometry"
	"github.com/julioguillermo/fit_anim/three/materials"
	"github.com/julioguillermo/fit_anim/three/point"
	"github.com/julioguillermo/fit_anim/three/scene"
)

const (
	Scale    = 0.07
	IntMax   = 0.07
	IntMin   = 0.04
	NextBias = 0.7
	Inc      = 0.5
	Dec      = 0.1
)

type GoNote struct {
	Note   string
	Octave int

	Mesh *three.GoMesh
	Mat  *materials.GoStandardMaterial
}

type GoNoteSharp struct {
	Note   string
	Octave int

	MeshFront *three.GoMesh
	MeshBack  *three.GoMesh
	MeshLeft  *three.GoMesh
	MeshRight *three.GoMesh
	Mat       *materials.GoStandardMaterial
}

type GoNotesPiramid struct {
	C   *GoNote
	Csh *GoNoteSharp
	D   *GoNote
	Dsh *GoNoteSharp
	E   *GoNote
	F   *GoNote
	Fsh *GoNoteSharp
	G   *GoNote
	Gsh *GoNoteSharp
	A   *GoNote
	Ash *GoNoteSharp
	B   *GoNote
}

func NotesPiramid(scene *scene.GoScene, o int, x, y, z float64) *GoNotesPiramid {
	piramid := &GoNotesPiramid{}

	piramid.B = createNote(scene, nil, 7, o)
	piramid.B.Mesh.Move(point.Point{X: x, Y: y, Z: z})

	piramid.A = createNote(scene, piramid.B.Mesh, 6, o)
	piramid.Ash = createNoteSharp(scene, piramid.A, 6)

	piramid.G = createNote(scene, piramid.B.Mesh, 5, o)
	piramid.Gsh = createNoteSharp(scene, piramid.G, 5)

	piramid.F = createNote(scene, piramid.B.Mesh, 4, o)
	piramid.Fsh = createNoteSharp(scene, piramid.F, 4)

	piramid.E = createNote(scene, piramid.B.Mesh, 3, o)

	piramid.D = createNote(scene, piramid.B.Mesh, 2, o)
	piramid.Dsh = createNoteSharp(scene, piramid.D, 2)

	piramid.C = createNote(scene, piramid.B.Mesh, 1, o)
	piramid.Csh = createNoteSharp(scene, piramid.C, 1)

	return piramid
}

func createNote(scene *scene.GoScene, root *three.GoMesh, n, o int) *GoNote {
	size := Scale*4 + float64(n)*Scale*1.3
	pos := float64(7-n) * Scale

	geo := geometry.BoxGeometry(size, Scale, size)

	mat := materials.StandardMaterial("#556")
	mat.SetEmissiveIntensity(0)
	mat.SetEmissive("#FF0055")

	mesh := three.Mesh(geo, mat)
	mesh.Move(point.Point{Y: pos})

	scene.Add(mesh)
	if root != nil {
		root.Add(mesh)
	}

	note := string("CDEFGAB"[n-1])

	return &GoNote{
		Note:   note,
		Octave: o,

		Mesh: mesh,
		Mat:  mat,
	}
}

func createNoteSharp(scene *scene.GoScene, parent *GoNote, n int) *GoNoteSharp {
	mat := materials.StandardMaterial("#334")
	mat.SetEmissiveIntensity(0)
	mat.SetEmissive("#FF00FF")

	size := (Scale*4 + float64(n)*Scale*1.3) / 2
	osize := Scale * 0.5
	move := size + Scale*0.3

	geoFB := geometry.BoxGeometry(size, osize, osize)
	geoLR := geometry.BoxGeometry(osize, osize, size)

	meshF := three.Mesh(geoFB, mat)
	meshB := three.Mesh(geoFB, mat)
	meshL := three.Mesh(geoLR, mat)
	meshR := three.Mesh(geoLR, mat)

	meshF.Move(point.Point{Z: move})
	meshB.Move(point.Point{Z: -move})
	meshL.Move(point.Point{X: move})
	meshR.Move(point.Point{X: -move})

	scene.Add(meshF)
	scene.Add(meshB)
	scene.Add(meshL)
	scene.Add(meshR)
	if parent != nil {
		parent.Mesh.Add(meshF)
		parent.Mesh.Add(meshB)
		parent.Mesh.Add(meshL)
		parent.Mesh.Add(meshR)
	}

	return &GoNoteSharp{
		Note:   parent.Note + "sh",
		Octave: parent.Octave,

		MeshFront: meshF,
		MeshBack:  meshB,
		MeshLeft:  meshL,
		MeshRight: meshR,
		Mat:       mat,
	}
}

func (p *GoNote) Update(frames *models.Frames) {
	startFreq := notes.GetNoteFreq(p.Note, p.Octave)
	endFreq := startFreq*(1-NextBias) + notes.GetNextNoteFreq(p.Note, p.Octave)*NextBias

	aInt := frames.FreqGtEq(int(startFreq)).FreqLt(int(endFreq)).MaxIntencity()
	// p.Mat.SetEmissiveIntensity(aInt)
	eInt := p.Mat.GetEmissiveIntensity()
	if aInt > IntMax {
		eInt = eInt*(1-Inc) + Inc
	} else if aInt < IntMin {
		eInt = eInt * Dec
	}
	p.Mat.SetEmissiveIntensity(eInt)
}

func (p *GoNoteSharp) Update(frames *models.Frames) {
	startFreq := notes.GetNoteFreq(p.Note, p.Octave)
	endFreq := notes.GetNextNoteFreq(p.Note, p.Octave)

	aInt := frames.FreqGtEq(int(startFreq)).FreqLt(int(endFreq)).MaxIntencity()
	// p.Mat.SetEmissiveIntensity(aInt)
	eInt := p.Mat.GetEmissiveIntensity()
	if aInt > IntMax {
		eInt = eInt*(1-Inc) + Inc
	} else if aInt < IntMin {
		eInt = eInt * Dec
	}
	p.Mat.SetEmissiveIntensity(eInt)
}

func (p *GoNotesPiramid) Update(frames *models.Frames) {
	p.C.Update(frames)
	p.Csh.Update(frames)
	p.D.Update(frames)
	p.Dsh.Update(frames)
	p.E.Update(frames)
	p.F.Update(frames)
	p.Fsh.Update(frames)
	p.G.Update(frames)
	p.Gsh.Update(frames)
	p.A.Update(frames)
	p.Ash.Update(frames)
	p.B.Update(frames)
}
