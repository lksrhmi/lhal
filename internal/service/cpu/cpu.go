package cpu

import (
	"cmd/lhal/internal/source"
	_ "fmt"
	"time"
	_ "time"
)


// returns a float64. CPU usage in percent 0-100
func GetUsage() (float64, error) {
	intervall := time.Millisecond * 100

	idle1, total1, getCpuTimesErr1 := source.ReadCpuStat()
	if getCpuTimesErr1 != nil {
		return 0, getCpuTimesErr1
	}

	// wait given time
	time.Sleep(intervall)

	idle2, total2, getCpuTimesErr2 := source.ReadCpuStat()
	if getCpuTimesErr2 != nil {
		return 0, getCpuTimesErr2
	}

	totalDelta := total2 - total1
	idleDelta := idle2 - idle1

	if totalDelta == 0 {
		return 0, nil
	}

	usage := float64(totalDelta-idleDelta) / float64(totalDelta) * 100
	
	return usage, nil
}