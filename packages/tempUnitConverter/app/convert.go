package app

import (
	"tempUnitConverter/domain"
)

func (t *Temperature) ToC(f domain.F) domain.C {
	return f.ToC()
}

func (t *Temperature) ToF(c domain.C) domain.F {
	return c.ToF()
}
