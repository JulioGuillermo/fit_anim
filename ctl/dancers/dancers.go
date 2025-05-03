package dancers

import (
	"syscall/js"

	"github.com/julioguillermo/fit_anim/three/materials"
)

func UpdateDancer(dancer js.Value, value float64) {
	if dancer.IsUndefined() {
		return
	}
	mesh := dancer.Call("getObjectByName", "Stormtrooper", true)
	if mesh.IsUndefined() {
		return
	}
	mat := mesh.Get("material")
	if mat.IsUndefined() {
		return
	}

	var goMat materials.GoPhongMaterial
	goMat.Material = mat
	goMat.SetEmissive("#FFAA00")
	goMat.SetEmissiveIntensity(value)
}
