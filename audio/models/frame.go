package models

import (
	"fmt"
	"log"
	"strings"
)

type Frame struct {
	Channel int

	Frame    uint64
	Duration uint64
	Start    uint64
	End      uint64

	Frequencies []Frequency
}

func (f *Frame) LoadFrame(code string) {
	codes := strings.Split(code, " -> ")
	if len(codes) < 1 {
		log.Fatalln("Fail to parse frame")
	}
	f.loadMetadata(codes[0])
	if len(codes) > 1 {
		f.loadFrequencies(codes[1])
	}
}

func (f *Frame) loadMetadata(code string) {
	fmt.Sscanf(code, "%d [%d - %d]", &f.Frame, &f.Start, &f.Duration)
	f.End = f.Start + f.Duration
}

func (f *Frame) loadFrequencies(code string) {
	codes := strings.Split(code, " ")
	f.Frequencies = make([]Frequency, len(codes))
	for i, code := range codes {
		var freq Frequency
		fmt.Sscanf(code, "%d:%f", &freq.Frequency, &freq.Intencity)
		f.Frequencies[i] = freq
	}
}
