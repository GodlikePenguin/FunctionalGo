package slices

import (
	"fmt"
	"strings"
)

func IntExample() {
	// Return the sum of the squares of the even numbers
	is := FromIntSlice([]int{1, 2, 3, 4})
	result :=
		is.Filter(func(item int) bool {
			return item%2 == 0
		}).Map(func(item int) int {
			return item * item
		}).Reduce(func(acc, item int) int {
			return acc + item
		}, 0)
	fmt.Println(result)
	// Prints 20
}

func StringExample() {
	// Take a sentance, and return only the words which are longer than 3 characters
	input := "The quick brown fox jumps over the lazy dog"
	s := strings.Split(input, " ")
	ss := FromStringSlice(s)
	result :=
		ss.Filter(func(s string) bool {
			return len(s) > 3
		}).Reduce(func(acc string, s string) string {
			if acc != "" {
				return acc + " " + s
			}
			return s
		}, "")
	fmt.Println(result)
	// Prints "quick brown jumps over lazy"
}
