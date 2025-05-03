package dancers

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/materials"
)

func UpdateLights(light js.Value, value float64) {
	if light.IsUndefined() {
		return
	}
	mat := light.Get("material")
	if mat.IsUndefined() {
		return
	}
	var goMat materials.GoPhongMaterial
	goMat.Material = mat
	goMat.SetEmissiveIntensity(value)
}
