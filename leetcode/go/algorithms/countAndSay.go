package algorithms

import (
	"strconv"
)

// Recursive solution
// Time Complexity: O(n) or O(m*n) with m being string length? not sure
// Space Complexity: O(n) - Recursive call stack is n
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	lastSaid := CountAndSay(n - 1)

	var count int
	var num string
	var say string

	for _, v := range lastSaid {
		if num == "" {
			count = 1
			num = string(v)
			continue
		}

		if num == string(v) {
			count++
			continue
		}

		say += strconv.Itoa(count)
		say += num
		count = 1
		num = string(v)
	}

	say += strconv.Itoa(count)
	say += num

	return say
}

/*
1.     1
2.     11
3.     21
4.     1211
5.     111221
*/
