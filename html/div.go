package html

import (
	"math"
	"syscall/js"
)

type Div struct {
	Div js.Value
}

func NewDiv(div js.Value) *Div {
	d := &Div{
		Div: div,
	}
	div.Get("style").Set("flex", "1")
	div.Get("style").Set("border-radius", "0.3rem")
	d.BG("#CCCCCC")
	return d
}

func (p *Div) BG(c string) {
	p.Div.Get("style").Set("background", c)
}

func (p *Div) SetInt(i, sf, ff float64) {
	f := (sf + ff) / 2
	i = math.Log10(i*10) * f
	if i > 0.1 {
		p.BG("#ff0088")
	} else {
		p.BG("#000000")
	}
	// c := byte(i * 255)
	// p.BG(fmt.Sprintf("rgb(0, 0, %d)", c))
}
