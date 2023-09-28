package app

import (
	"tempUnitConverter/domain"
)

func (t *Temperature) FToC(f domain.F) domain.C {
	return f.ToC()
}

func (t *Temperature) CToF(c domain.C) domain.F {
	return c.ToF()
}
