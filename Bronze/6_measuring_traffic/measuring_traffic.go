package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type SensorInfo struct {
	loc string
	lb  int
	ub  int
}

func getStartEndEstimate(sensors []SensorInfo) (int, int, int, int) {
	agg_lb, agg_ub := 0, 0
	start_lb, start_ub := 0, math.MaxInt
	end_lb, end_ub := 0, math.MaxInt
	for _, val := range sensors {
		if val.loc == "on" {
			agg_lb -= val.ub
			agg_ub -= val.lb

		} else if val.loc == "off" {
			start_lb = max(start_lb, val.lb+agg_ub)

			agg_lb += val.lb
			agg_ub += val.ub

		} else {
			start_lb = max(start_lb, val.lb+agg_lb)
			start_ub = min(start_ub, val.ub+agg_ub)
		}
	}
	end_lb = max(end_lb, start_lb-agg_ub)
	end_ub = min(end_ub, start_ub-agg_lb)

	return start_lb, start_ub, end_lb, end_ub
}

func main() {
	var n int

	vals := []SensorInfo{}

	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)

	for range n {
		if scanner.Scan() {
			line := scanner.Text()
			parts := strings.Fields(line)

			lb, _ := strconv.Atoi(parts[1])
			ub, _ := strconv.Atoi(parts[2])
			vals = append(vals, SensorInfo{loc: parts[0], lb: lb, ub: ub})
		}
	}

	start_lb, start_ub, end_lb, end_ub := getStartEndEstimate(vals)

	fmt.Println(start_lb, start_ub)
	fmt.Println(end_lb, end_ub)
}
