package main

import (
	"fmt"
	"strconv"
)

 var wds map[int64]string = map[int64]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "fourty",
	50: "fifty",
	60: "sixy",
	70: "seventy",
	80: "eighty",
	90: "ninety",
	100: "hundred",
	1000: "thousand",
	1000000: "million",
	1000000000: "billion",
	1000000000000: "trillion",
}

func getWords(n int64) string{
	if wds[n] != ""{
		return wds[n]
	}

	h := n / 100
	rem := n % 100

	if h <= 0 {
		if wds[n] != ""{
			return wds[n]
		}

		t := (n / 10) * 10
		r := n % 10

		return wds[t] + " " + wds[r]
	}
	
	str := wds[h] + " hundred"

	if rem != 0 {
		str += " and " + getWords(rem)
	}

	return str
}

func getString(n int64) string {
	
	str := fmt.Sprint(n)
	nums := []int64{}

	start := 0
	end := 0
	// num := 0
	for i := len(str); i > 0; i-=3{
		start = i - 3
		end = i

		if start < 0 {
			start = 0
		}
		num, _  := strconv.Atoi(str[start: end])
		nums = append(nums, int64(num))
	}

	dens := [4]int64{1000, 1000000, 1000000000, 1000000000000}
	strs := []string{}

	str = ""
	index := 0
	for i := 0; i < len(nums); i++ {
		str = getWords(nums[i])
		if i > 0 {
			str += " " + getWords(dens[index])
			index++
		}
		strs = append(strs, str)
	}

	str = ""
	for i := len(strs) - 1; i >= 0; i-- {
		if len(strs) - 1 > i {
			str += ", "
		}

		str +=  strs[i]
	}

	return str
}

func main() {
	var num int64 = 916

	fmt.Println(getString(num))
	fmt.Println(getString(0))
}

