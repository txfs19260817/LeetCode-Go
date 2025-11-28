package k

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogEntry(t *testing.T) {
	line := "44776.619 KTB918 310E MAINROAD"
	entry, err := NewLogEntry(line)
	assert.NoError(t, err)
	assert.InDelta(t, 44776.619, entry.Timestamp, 0.0001)
	assert.Equal(t, "KTB918", entry.LicensePlate)
	assert.Equal(t, 310, entry.Location)
	assert.Equal(t, "EAST", entry.Direction)
	assert.Equal(t, "MAINROAD", entry.BoothType)
}

func TestCountJourneys(t *testing.T) {
	logData := `90750.191 JOX304 250E ENTRY
91081.684 JOX304 260E MAINROAD
91483.251 JOX304 270E MAINROAD
91874.493 JOX304 280E EXIT`

	lf, err := NewLogFile(strings.NewReader(logData))
	assert.NoError(t, err)
	assert.Equal(t, 1, lf.CountJourneys())
}

func TestCatchSpeeders(t *testing.T) {
	// Case 1: 130km/h rule
	// 10km in 275s = 130.9 km/h -> Speeding
	logData1 := `1000.000 TST002 270W ENTRY
1275.000 TST002 260W EXIT`
	lf1, _ := NewLogFile(strings.NewReader(logData1))
	speeders1 := lf1.CatchSpeeders()
	assert.Contains(t, speeders1, "TST002")

	// Case 2: 120km/h rule (2 segments)
	// Segment 1: 10km in 290s = 124 km/h
	// Segment 2: 10km in 290s = 124 km/h
	// Total: 2 segments >= 120 -> Speeding
	logData2 := `1000.000 TST003 270W ENTRY
1290.000 TST003 260W MAINROAD
1580.000 TST003 250W EXIT`
	lf2, _ := NewLogFile(strings.NewReader(logData2))
	speeders2 := lf2.CatchSpeeders()
	assert.Contains(t, speeders2, "TST003")

	// Case 3: Normal driving
	// 10km in 400s = 90 km/h
	logData3 := `1000.000 OK001 270W ENTRY
1400.000 OK001 260W EXIT`
	lf3, _ := NewLogFile(strings.NewReader(logData3))
	speeders3 := lf3.CatchSpeeders()
	assert.Empty(t, speeders3)
}
