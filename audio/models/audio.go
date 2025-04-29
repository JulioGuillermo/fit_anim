package models

import (
	"io"
	"log"
	"net/http"
	"syscall/js"
)

type Audio struct {
	Channels []Channel
}

func GetAudioFromUrls(urls ...string) *Audio {
	codes := []string{}
	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println(err)
			continue
		}
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			continue
		}
		code := string(bytes)
		codes = append(codes, code)
	}

	return GetAudio(codes...)
}

func GetAudio(codes ...string) *Audio {
	channelsList := []Channel{}
	for _, code := range codes {
		var channel Channel
		channel.LoadChannel(code)
		channelsList = append(channelsList, channel)
	}

	return &Audio{
		Channels: channelsList,
	}
}

func (a *Audio) GetChannel(c int) Channel {
	return a.Channels[c]
}

func (a *Audio) GetJS() js.Value {
	val := js.Value{}

	val.Set("getFrames", GetNumFunc(func(p float64) any {
		return a.GetTimeFrames(uint64(p)).GetJS()
	}))

	return val
}

func (a *Audio) GetTimeFrames(msec uint64) *Frames {
	frames := make([]Frame, len(a.Channels))
	for i, c := range a.Channels {
		frames[i] = c.GetFrameByTime(msec)
	}
	return &Frames{
		Frames: frames,
	}
}
