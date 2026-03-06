package cli

import (
	"cmd/lhal/internal/souce"
	"fmt"
)

func Render(cpuData source.CPUStat) {
	fmt.Println(cpuData)
}