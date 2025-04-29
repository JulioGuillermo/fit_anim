package html

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/audio/models"
)

const Rows = 10
const Keys = 40

type GoHTML struct {
	Audio  *models.Audio
	Player js.Value

	Elements []*Div
}

func NewGoHTML(audio *models.Audio) *GoHTML {
	h := &GoHTML{}
	h.Player = js.Global().Get("document").Call("getElementById", "audio")
	h.Audio = audio
	h.InitComponent()
	return h
}

func (p *GoHTML) InitComponent() {
	box := js.Global().Get("document").Call("createElement", "div")
	box.Get("style").Set("background", "#333333")
	box.Get("style").Set("position", "fixed")
	box.Get("style").Set("left", "0")
	box.Get("style").Set("right", "0")
	box.Get("style").Set("bottom", "0")

	js.Global().Get("document").Get("body").Call("appendChild", box)

	for _ = range Rows {
		row := js.Global().Get("document").Call("createElement", "div")
		row.Get("style").Set("display", "flex")
		row.Get("style").Set("gap", "0.2rem")
		row.Get("style").Set("padding", "0.1rem")
		row.Get("style").Set("height", "2rem")
		box.Call("appendChild", row)

		for _ = range Keys {
			div := js.Global().Get("document").Call("createElement", "div")
			row.Call("appendChild", div)

			d := NewDiv(div)
			p.Elements = append(p.Elements, d)
		}
	}
}

func (p *GoHTML) Update() {
	audioTime := p.Player.Get("currentTime").Float()
	frames := p.Audio.GetTimeFrames(uint64(audioTime * 1000))

	const MinF = 20
	const MaxF = 5000

	keys := len(p.Elements)
	IncF := (5000 - 20) / keys

	for i := range keys {
		StartF := IncF * i
		EndF := StartF + IncF
		fInt := frames.
			FreqGtEq(StartF).
			FreqLt(EndF).
			Int()
		p.Elements[i].SetInt(fInt, float64(StartF), float64(EndF))
	}
}
