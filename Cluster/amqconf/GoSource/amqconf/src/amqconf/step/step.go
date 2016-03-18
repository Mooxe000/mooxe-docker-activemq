package step

import (
	gd "github.com/Mooxe000/GoDash"
	"github.com/clbanning/mxj"
)

var (
	pln    = gd.Pln
	prf    = gd.Prf
	dd     = gd.Dd
	typeof = gd.TypeOf
)

type Step struct {
	Value *mxj.Map
}
