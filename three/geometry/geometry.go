package geometry

import "syscall/js"

type GoGeometry struct {
	Geometry js.Value
}

type Geometry interface {
	GetGeometry() *GoGeometry
}
