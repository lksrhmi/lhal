package main

import (
	source "cmd/lhal/internal/souce"
	"cmd/lhal/internal/ui/cli"
)

func main() {
	cpuData, _ := source.ReadCPUStat()

	cli.Render(cpuData)
}