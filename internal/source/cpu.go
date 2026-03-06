package source

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type CpuStat struct {
	User    uint64
	Nice    uint64
	System  uint64
	Idle    uint64
	Iowait  uint64
	IRQ     uint64
	SoftIRQ uint64

	Total uint64
}

func ReadCPUStat() (CpuStat, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return CpuStat{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "cpu ") {
			fields := strings.Fields(line)

			user, _ := strconv.ParseUint(fields[1], 10, 64)
			nice, _ := strconv.ParseUint(fields[2], 10, 64)
			system, _ := strconv.ParseUint(fields[3], 10, 64)
			idle, _ := strconv.ParseUint(fields[4], 10, 64)
			iowait, _ := strconv.ParseUint(fields[5], 10, 64)
			irq, _ := strconv.ParseUint(fields[6], 10, 64)
			softirq, _ := strconv.ParseUint(fields[7], 10, 64)

			return CpuStat{
				User:    user,
				Nice:    nice,
				System:  system,
				Idle:    idle,
				Iowait:  iowait,
				IRQ:     irq,
				SoftIRQ: softirq,
				
				Total: user + nice + system + idle + iowait + irq + softirq,
			}, nil
		}
	}

	return CpuStat{}, nil
}