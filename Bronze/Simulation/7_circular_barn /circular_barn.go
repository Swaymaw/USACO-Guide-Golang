package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getMinDistance(nums []int) int {
	preSum := nums[0]
	preDistance := nums[0] * (len(nums) - 1)
	postSum := 0
	postDistance := 0
	ans := math.MaxInt

	for idx, val := range nums[1:] {
		postSum += val
		postDistance += (idx + 1) * val
	}

	ans = min(ans, postDistance)

	for idx := range nums[1:] {
		postDistance -= postSum
		postSum -= nums[idx+1]

		ans = min(ans, preDistance+postDistance)

		preDistance += nums[idx+1] * (len(nums) - idx)
		preDistance -= preSum
		preSum += nums[idx+1]
	}

	return ans
}

func main() {
	var n int
	nums := make([]int, 0, n)

	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)

	for range n {
		if scanner.Scan() {
			val := scanner.Text()
			int_val, _ := strconv.Atoi(val)
			nums = append(nums, int_val)
		}
	}

	result := getMinDistance(nums)

	fmt.Println("Min distance travelled by cows:", result)

}
