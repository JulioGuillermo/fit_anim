package ctl

import (
	"time"

	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/ctl/elements"
	"github.com/julioguillermo/fit_anim/three/camera"
	"github.com/julioguillermo/fit_anim/three/renderer"
	"github.com/julioguillermo/fit_anim/three/scene"
)

type Controller struct {
	Audio *models.Audio

	Scene    *scene.GoScene
	Camera   *camera.GoCamera
	Renderer *renderer.GoRenderer

	FreqBoxs0 *elements.GoFreqBoxs
	FreqBoxs1 *elements.GoFreqBoxs

	FreqSphs0 *elements.GoFreqSphs
	FreqSphs1 *elements.GoFreqSphs

	MainElement *elements.MainElement
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

	ch0 := frames.ChanEq(0)
	ch1 := frames.ChanEq(1)

	c.FreqBoxs0.Update(ch0)
	c.FreqBoxs1.Update(ch1)

	c.FreqSphs0.Update(ch0)
	c.FreqSphs1.Update(ch1)
}
