package ctl

import (
	"github.com/julioguillermo/fit_anim/ctl/elements"
)

func (c *Controller) initScene() {
	c.MainElement = elements.CreateMainElement(c.Scene)

	OctInit := 2

	c.PiramidL1 = elements.NotesPiramid(c.Scene, OctInit, -1.3, -0.39, 0.6)
	c.PiramidL2 = elements.NotesPiramid(c.Scene, OctInit+1, -1.3, -0.39, -2.0)
	c.PiramidL3 = elements.NotesPiramid(c.Scene, OctInit+2, -1.3, -0.39, -3.6)
	c.PiramidL4 = elements.NotesPiramid(c.Scene, OctInit+3, -1.3, -0.39, -6)
	c.PiramidL5 = elements.NotesPiramid(c.Scene, OctInit+4, -1.3, -0.39, -7.6)

	c.PiramidR1 = elements.NotesPiramid(c.Scene, OctInit, 1.3, -0.39, 0.6)
	c.PiramidR2 = elements.NotesPiramid(c.Scene, OctInit+1, 1.3, -0.39, -2.0)
	c.PiramidR3 = elements.NotesPiramid(c.Scene, OctInit+2, 1.3, -0.39, -3.6)
	c.PiramidR4 = elements.NotesPiramid(c.Scene, OctInit+3, 1.3, -0.39, -6)
	c.PiramidR5 = elements.NotesPiramid(c.Scene, OctInit+4, 1.3, -0.39, -7.6)
}
