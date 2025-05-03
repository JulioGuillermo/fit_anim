package main

import (
	"github.com/julioguillermo/fit_anim/audio/models"
	"github.com/julioguillermo/fit_anim/ctl"
	"github.com/julioguillermo/fit_anim/three"
)

func loadAudioData(c *ctl.Controller) {
	audio := models.GetAudioFromUrls(
		"/example/output/output_channel_0",
		// "/example/output/output_channel_1",
	)
	c.Audio = audio
}

func main() {
	// js.Global().Set("loadAudioData", js.FuncOf(func(this js.Value, args []js.Value) any {
	// 	codes := []string{}
	// 	for _, a := range args {
	// 		if a.Type() == js.TypeString {
	// 			codes = append(codes, a.String())
	// 		}
	// 	}
	// 	return models.GetAudio(codes...).GetJS()
	// }))
	c := &ctl.Controller{}
	three.InitGoThree(c.InitScene, c.RenderScene)

	loadAudioData(c)

	// h := html.NewGoHTML(c.Audio)
	// js.Global().Set("updateHTML", js.FuncOf(func(this js.Value, args []js.Value) any {
	// 	h.Update()
	// 	return nil
	// }))

	<-make(chan bool)
}
