package models

import (
	"fmt"
	"strings"
)

type Channel struct {
	Channel   int
	NumFrames int
	MinFreq   float64
	MaxFreq   float64

	Frames []Frame
}

func (c *Channel) LoadChannel(code string) {
	lines := strings.Split(code, "\n")
	c.loadMetaData(lines[0])

	c.Frames = make([]Frame, 0, len(lines)-1)
	for i := range c.NumFrames {
		c.loadFrame(lines[i+1])
	}
}

func (c *Channel) loadMetaData(line string) {
	fmt.Sscanf(
		line,
		"CHAN [%d] -> Frames %d MinFreq %f MaxFreq %f",
		&c.Channel,
		&c.NumFrames,
		&c.MinFreq,
		&c.MaxFreq,
	)
}

func (c *Channel) loadFrame(line string) {
	var frame Frame
	frame.Channel = c.Channel
	frame.LoadFrame(line)
	c.Frames = append(c.Frames, frame)
}

func (c *Channel) MaxDuration() uint64 {
	return c.Frames[len(c.Frames)-1].End
}

func (c *Channel) GetFrameByTime(msec uint64) Frame {
	msec = msec % c.MaxDuration()
	for _, f := range c.Frames {
		if f.Start <= msec && f.End > msec {
			return f
		}
	}
	return c.Frames[0]
}
