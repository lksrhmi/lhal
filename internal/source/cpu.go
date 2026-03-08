package source

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadCpuStat reads cpu times. gives back idle, total and error
func ReadCpuStat() (idle, total uint64, err error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// reads cpu line
		if len(fields) > 0 && fields[0] == "cpu" {
			var currentTotal uint64
			for i := 1; i < len(fields); i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					return 0, 0, err
				}
				currentTotal += val

				// Field 4 is 'idle' time
				if i == 4 {
					idle = val
				}
			}
			return idle, currentTotal, nil
		}
	}
	return 0, 0, fmt.Errorf("could not find cpu line in /proc/stat")
}
