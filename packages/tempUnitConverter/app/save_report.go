package app

import (
	"tempUnitConverter/domain"
)

func (c *Temperature) SaveFtoC(f domain.F) (int, error) {
	return c.db.SaveTemperatureReport(f.ToC())
}
