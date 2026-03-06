package cpu

import (
	"cmd/lhal/internal/source"
	_ "fmt"
	"time"
)

func GetUsage() (float64, error) {

	usage, _:= calculate()

	return usage, nil
}

func calculate() (float64, error) {
	const wait = 50 * time.Millisecond

	rawCpu1, err := source.ReadCPUStat()
	if err != nil {
		return 0, err
	}

	time.Sleep(wait)

	rawCpu2, err := source.ReadCPUStat()
	if err != nil {
		return 0, err
	}

	totalDelta := rawCpu2.Total - rawCpu1.Total
	idleDelta := rawCpu2.Idle - rawCpu1.Idle

	if totalDelta == 0 {
		return 0, nil
	}

	usage := float64(totalDelta-idleDelta) / float64(totalDelta) * 100

	return usage, nil
}