package sampledb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAllTemperatureReports(t *testing.T) {
	r, err := db.GetAllTemperatureReports()
	require.NoError(t, err)
	rlen := len(r)
	dataLen := len(db.db.Data)
	require.Equal(t, dataLen, rlen)
	require.Equal(t, r[0].ID, 0)
}

func TestGetTemperatureReport(t *testing.T) {
	r, err := db.GetTemperatureReport(0)
	require.NoError(t, err)
	require.Equal(t, r.ID, 0)

	_, err = db.GetTemperatureReport(6)
	require.Error(t, err)
}
