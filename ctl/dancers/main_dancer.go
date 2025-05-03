package dancers

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/materials"
)

func UpdateMainDancer(dancer js.Value, value float64) {
	if dancer.IsUndefined() {
		return
	}
	mat := dancer.Get("material")
	if mat.IsUndefined() {
		return
	}
	var goMat materials.GoPhongMaterial
	goMat.Material = mat
	goMat.SetEmissive("#00FF88")
	goMat.SetEmissiveIntensity(value)
}
