package main

import "fmt"

func countDuplicate(nums []int32) int32 {

	// var i int32
	var count int32
	var temp []int32
	for len(nums) > 1 {

		temp = append(temp, nums[0])
		nums = nums[1:]

		for _, eval := range nums {

			if eval == temp[len(temp)-1] {
				count++
				nums = removeEl(nums[0], nums)
				break
			}
		}

	}

	return count
}

func removeEl(num int32, nums []int32) []int32 {
	var out []int32
	for _, el := range nums {
		if el == num {
			continue
		} else {
			out = append(out, el)
		}
	}
	return out
}

func main() {
	nums := []int32{1, 10, 2, 3, 4, 10}
	nums2 := []int32{1, 2, 3, 4, 3, 9, 9, 1}

	fmt.Println(countDuplicate(nums))
	fmt.Println(countDuplicate(nums2))

}
