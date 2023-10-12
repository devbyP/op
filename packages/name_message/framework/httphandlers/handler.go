package httphandlers

import "tempUnitConverter/app"

type Adapter struct {
	app app.AppPort
}

func NewAdapter(a app.AppPort) *Adapter {
	return &Adapter{app: a}
}
