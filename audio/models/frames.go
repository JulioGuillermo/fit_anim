package models

import "syscall/js"

type Frames struct {
	Frames []Frame
}

func (f *Frames) GetJS() js.Value {
	val := js.Value{}

	val.Set("chanEq", GetNumFunc(func(p float64) any {
		return f.ChanEq(int(p)).GetJS()
	}))
	val.Set("chanNotEq", GetNumFunc(func(p float64) any {
		return f.ChanNotEq(int(p)).GetJS()
	}))
	val.Set("chanGt", GetNumFunc(func(p float64) any {
		return f.ChanGt(int(p)).GetJS()
	}))
	val.Set("chanGtEq", GetNumFunc(func(p float64) any {
		return f.ChanGtEq(int(p)).GetJS()
	}))
	val.Set("chanLt", GetNumFunc(func(p float64) any {
		return f.ChanGt(int(p)).GetJS()
	}))
	val.Set("chanLtEq", GetNumFunc(func(p float64) any {
		return f.ChanGtEq(int(p)).GetJS()
	}))

	val.Set("freqEq", GetNumFunc(func(p float64) any {
		return f.FreqEq(int(p)).GetJS()
	}))
	val.Set("freqNotEq", GetNumFunc(func(p float64) any {
		return f.FreqNotEq(int(p)).GetJS()
	}))
	val.Set("freqGt", GetNumFunc(func(p float64) any {
		return f.FreqGt(int(p)).GetJS()
	}))
	val.Set("freqGtEq", GetNumFunc(func(p float64) any {
		return f.FreqGtEq(int(p)).GetJS()
	}))
	val.Set("freqLt", GetNumFunc(func(p float64) any {
		return f.FreqGt(int(p)).GetJS()
	}))
	val.Set("freqLtEq", GetNumFunc(func(p float64) any {
		return f.FreqGtEq(int(p)).GetJS()
	}))

	val.Set("intEq", GetNumFunc(func(p float64) any {
		return f.IntEq(p).GetJS()
	}))
	val.Set("intNotEq", GetNumFunc(func(p float64) any {
		return f.IntNotEq(p).GetJS()
	}))
	val.Set("intGt", GetNumFunc(func(p float64) any {
		return f.IntGt(p).GetJS()
	}))
	val.Set("intGtEq", GetNumFunc(func(p float64) any {
		return f.IntGtEq(p).GetJS()
	}))
	val.Set("intLt", GetNumFunc(func(p float64) any {
		return f.IntGt(p).GetJS()
	}))
	val.Set("intLtEq", GetNumFunc(func(p float64) any {
		return f.IntGtEq(p).GetJS()
	}))

	val.Set("chans", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.Channels()
	}))
	val.Set("freqs", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.Frequencies()
	}))
	val.Set("ints", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.Intencities()
	}))

	val.Set("chanCount", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.CountChannels()
	}))
	val.Set("freqCount", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.CountFrequencies()
	}))
	val.Set("intCount", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.CountIntencities()
	}))

	val.Set("chan", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.AvgChannels()
	}))
	val.Set("freq", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.AvgFrequencies()
	}))
	val.Set("int", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.AvgIntencities()
	}))

	val.Set("chanSum", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.SumChannels()
	}))
	val.Set("freqSum", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.SumFrequencies()
	}))
	val.Set("intSum", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.SumIntencities()
	}))

	val.Set("chanFirst", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.FirstChannel()
	}))
	val.Set("freqFirst", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.FirstFrequency()
	}))
	val.Set("intFirst", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.FirstIntencity()
	}))

	val.Set("chanLast", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.FirstChannel()
	}))
	val.Set("freqLast", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.FirstFrequency()
	}))
	val.Set("intLast", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.FirstIntencity()
	}))

	val.Set("chanMax", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.MaxChannel()
	}))
	val.Set("freqMax", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.MaxFrequency()
	}))
	val.Set("intMax", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.MaxIntencity()
	}))

	val.Set("chanMin", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.MinChannel()
	}))
	val.Set("freqMin", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.MinFrequency()
	}))
	val.Set("intMin", js.FuncOf(func(this js.Value, args []js.Value) any {
		return f.MinIntencity()
	}))

	return val
}

func (f *Frames) Extends(frames *Frames) *Frames {
	return &Frames{
		Frames: append(f.Frames, frames.Frames...),
	}
}

func (f *Frames) FilterChannels(fun func(Frame) bool) *Frames {
	if f == nil {
		return nil
	}
	frames := []Frame{}
	for _, fr := range f.Frames {
		if fun(fr) {
			frames = append(frames, fr)
		}
	}
	return &Frames{
		Frames: frames,
	}
}

func (f *Frames) FilterFrequencies(fun func(Frequency) bool) *Frames {
	if f == nil {
		return nil
	}
	frames := []Frame{}
	for _, fr := range f.Frames {
		frequencies := []Frequency{}
		for _, fq := range fr.Frequencies {
			if fun(fq) {
				frequencies = append(frequencies, fq)
			}
		}
		if len(frequencies) == 0 {
			continue
		}
		frame := Frame{
			Channel: fr.Channel,
			Frame:   fr.Frame,
			Start:   fr.Start,
			End:     fr.End,

			Frequencies: frequencies,
		}
		frames = append(frames, frame)
	}
	return &Frames{
		Frames: frames,
	}
}

func (f *Frames) ChanEq(c int) *Frames {
	return f.FilterChannels(func(f Frame) bool {
		return f.Channel == c
	})
}

func (f *Frames) ChanNotEq(c int) *Frames {
	return f.FilterChannels(func(f Frame) bool {
		return f.Channel != c
	})
}

func (f *Frames) ChanGt(c int) *Frames {
	return f.FilterChannels(func(f Frame) bool {
		return f.Channel > c
	})
}

func (f *Frames) ChanGtEq(c int) *Frames {
	return f.FilterChannels(func(f Frame) bool {
		return f.Channel >= c
	})
}

func (f *Frames) ChanLt(c int) *Frames {
	return f.FilterChannels(func(f Frame) bool {
		return f.Channel < c
	})
}

func (f *Frames) ChanLtEq(c int) *Frames {
	return f.FilterChannels(func(f Frame) bool {
		return f.Channel <= c
	})
}

func (p *Frames) FreqEq(f int) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Frequency == f
	})
}

func (p *Frames) FreqNotEq(f int) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Frequency != f
	})
}

func (p *Frames) FreqGt(f int) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Frequency > f
	})
}

func (p *Frames) FreqGtEq(f int) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Frequency >= f
	})
}

func (p *Frames) FreqLt(f int) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Frequency < f
	})
}

func (p *Frames) FreqLtEq(f int) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Frequency <= f
	})
}

func (p *Frames) IntEq(i float64) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Intencity == i
	})
}

func (p *Frames) IntNotEq(i float64) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Intencity != i
	})
}

func (p *Frames) IntGt(i float64) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Intencity > i
	})
}

func (p *Frames) IntGtEq(i float64) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Intencity >= i
	})
}

func (p *Frames) IntLt(i float64) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Intencity < i
	})
}

func (p *Frames) IntLtEq(i float64) *Frames {
	return p.FilterFrequencies(func(fq Frequency) bool {
		return fq.Intencity <= i
	})
}

func (p *Frames) Channels() []int {
	channels := []int{}
	if p == nil {
		return channels
	}
	for _, f := range p.Frames {
		channels = append(channels, f.Channel)
	}
	return channels
}

func (p *Frames) CountChannels() int {
	chans := p.Channels()
	return len(chans)
}

func (p *Frames) FirstChannel() int {
	chans := p.Channels()
	if len(chans) == 0 {
		return 0
	}
	return chans[0]
}

func (p *Frames) LastChannel() int {
	chans := p.Channels()
	if len(chans) == 0 {
		return 0
	}
	return chans[len(chans)-1]
}

func (p *Frames) MaxChannel() int {
	chans := p.Channels()
	if len(chans) == 0 {
		return 0
	}
	max := chans[0]
	for _, c := range chans {
		if c > max {
			max = c
		}
	}
	return max
}

func (p *Frames) MinChannel() int {
	chans := p.Channels()
	if len(chans) == 0 {
		return 0
	}
	min := chans[0]
	for _, c := range chans {
		if c < min {
			min = c
		}
	}
	return min
}

func (p *Frames) SumChannels() int {
	chans := p.Channels()
	sum := 0
	for _, c := range chans {
		sum += c
	}
	return sum
}

func (p *Frames) AvgChannels() float64 {
	chans := p.Channels()
	sum := 0
	for _, c := range chans {
		sum += c
	}
	return float64(sum) / float64(len(chans))
}

func (p *Frames) Frequencies() []int {
	freqs := []int{}
	if p == nil {
		return freqs
	}
	for _, f := range p.Frames {
		for _, fq := range f.Frequencies {
			freqs = append(freqs, fq.Frequency)
		}
	}
	return freqs
}

func (p *Frames) CountFrequencies() int {
	freqs := p.Frequencies()
	return len(freqs)
}

func (p *Frames) FirstFrequency() int {
	freqs := p.Frequencies()
	if len(freqs) == 0 {
		return 0
	}
	return freqs[0]
}

func (p *Frames) LastFrequency() int {
	freqs := p.Frequencies()
	if len(freqs) == 0 {
		return 0
	}
	return freqs[len(freqs)-1]
}

func (p *Frames) MaxFrequency() int {
	freqs := p.Frequencies()
	if len(freqs) == 0 {
		return 0
	}
	max := freqs[0]
	for _, c := range freqs {
		if c > max {
			max = c
		}
	}
	return max
}

func (p *Frames) MinFrequency() int {
	freqs := p.Frequencies()
	if len(freqs) == 0 {
		return 0
	}
	min := freqs[0]
	for _, c := range freqs {
		if c < min {
			min = c
		}
	}
	return min
}

func (p *Frames) SumFrequencies() int {
	freqs := p.Frequencies()
	sum := 0
	for _, c := range freqs {
		sum += c
	}
	return sum
}

func (p *Frames) AvgFrequencies() float64 {
	freqs := p.Frequencies()
	if len(freqs) == 0 {
		return 0
	}
	sum := 0
	for _, c := range freqs {
		sum += c
	}
	return float64(sum) / float64(len(freqs))
}

func (p *Frames) Freq() float64 {
	return p.AvgFrequencies()
}

func (p *Frames) Intencities() []float64 {
	ints := []float64{}
	if p == nil {
		return ints
	}
	for _, f := range p.Frames {
		for _, fq := range f.Frequencies {
			ints = append(ints, fq.Intencity)
		}
	}
	return ints
}

func (p *Frames) CountIntencities() int {
	ints := p.Intencities()
	return len(ints)
}

func (p *Frames) FirstIntencity() float64 {
	ints := p.Intencities()
	if len(ints) == 0 {
		return 0
	}
	return ints[0]
}

func (p *Frames) LastIntencity() float64 {
	ints := p.Intencities()
	if len(ints) == 0 {
		return 0
	}
	return ints[len(ints)-1]
}

func (p *Frames) MaxIntencity() float64 {
	ints := p.Intencities()
	if len(ints) == 0 {
		return 0
	}
	max := ints[0]
	for _, c := range ints {
		if c > max {
			max = c
		}
	}
	return max
}

func (p *Frames) MinIntencity() float64 {
	ints := p.Intencities()
	if len(ints) == 0 {
		return 0
	}
	min := ints[0]
	for _, c := range ints {
		if c < min {
			min = c
		}
	}
	return min
}

func (p *Frames) SumIntencities() float64 {
	ints := p.Intencities()
	sum := float64(0)
	for _, c := range ints {
		sum += c
	}
	return sum
}

func (p *Frames) AvgIntencities() float64 {
	ints := p.Intencities()
	if len(ints) == 0 {
		return 0
	}
	sum := float64(0)
	for _, c := range ints {
		sum += c
	}
	return float64(sum) / float64(len(ints))
}

func (p *Frames) Int() float64 {
	return p.AvgIntencities()
}
