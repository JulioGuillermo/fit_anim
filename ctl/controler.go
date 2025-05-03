package ctl

import (
	"time"

	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/ctl/dancers"
	"github.com/julioguillermo/fit_anim/ctl/elements"
	"github.com/julioguillermo/fit_anim/three/camera"
	"github.com/julioguillermo/fit_anim/three/light"
	"github.com/julioguillermo/fit_anim/three/renderer"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type Controller struct {
	Audio *models.Audio

	Scene    *scene.GoScene
	Camera   *camera.GoCamera
	Renderer *renderer.GoRenderer

	TopLight *light.GoDirectionalLight

	MainElement *elements.MainElement

	PiramidL1 *elements.GoNotesPiramid
	PiramidL2 *elements.GoNotesPiramid
	PiramidL3 *elements.GoNotesPiramid
	PiramidL4 *elements.GoNotesPiramid
	PiramidL5 *elements.GoNotesPiramid

	PiramidR1 *elements.GoNotesPiramid
	PiramidR2 *elements.GoNotesPiramid
	PiramidR3 *elements.GoNotesPiramid
	PiramidR4 *elements.GoNotesPiramid
	PiramidR5 *elements.GoNotesPiramid
}

func (c *Controller) InitScene(scene *scene.GoScene, camera *camera.GoCamera, renderer *renderer.GoRenderer, start time.Time) {
	c.Scene = scene
	c.Camera = camera
	c.Renderer = renderer

	c.initCamera()
	c.initLights()
	c.initScene()
}

func (c *Controller) RenderScene(
	scene *scene.GoScene,
	camera *camera.GoCamera,
	renderer *renderer.GoRenderer,
	start time.Time,
	last time.Time,
	delta float64,
	audioTime float64,
) {
	var frames *models.Frames = nil
	if c.Audio != nil {
		msec := uint64(audioTime * 1000)
		frames = c.Audio.GetTimeFrames(msec)
	}

	c.MainElement.RenderScene(
		scene,
		camera,
		renderer,
		start,
		last,
		delta,
		frames,
	)

	c.PiramidL1.Update(frames.ChanEq(0))
	c.PiramidL2.Update(frames.ChanEq(0))
	c.PiramidL3.Update(frames.ChanEq(0))
	c.PiramidL4.Update(frames.ChanEq(0))
	c.PiramidL5.Update(frames.ChanEq(0))

	c.PiramidR1.Update(frames.ChanEq(1))
	c.PiramidR2.Update(frames.ChanEq(1))
	c.PiramidR3.Update(frames.ChanEq(1))
	c.PiramidR4.Update(frames.ChanEq(1))
	c.PiramidR5.Update(frames.ChanEq(1))

	maxInt := frames.MaxIntencity()

	c.TopLight.SetIntencity(maxInt)

	dancers.UpdateLights(c.Scene.GetJSObjByName("TopLights"), maxInt)
	dancers.UpdateMainDancer(c.Scene.GetJSObjByName("Alpha_Surface"), maxInt)
	dancers.UpdateDancer(c.Scene.GetJSObjByName("DancerL"), maxInt)
	dancers.UpdateDancer(c.Scene.GetJSObjByName("DancerR"), maxInt)
}
