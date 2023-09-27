package sampledb

import (
	"testing"

	"tempUnitConverter/domain"

	"github.com/stretchr/testify/require"
)

func TestSaveTemperatureReport(t *testing.T) {
	testTem := domain.C(33.5)
	id, err := db.SaveTemperatureReport(testTem)
	require.NoError(t, err)
	require.Equal(t, id, 1)

	r, _ := db.GetTemperatureReport(id)
	require.Equal(t, r.ID, 1)
	require.Equal(t, r.TempInC, testTem.Float64())
}
