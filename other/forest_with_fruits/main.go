package main

import (
	"fmt"
)

func generateLongestSequence(tree []int) int {
	pre := -1
	count := 1

	max := 0 // local maximum
	seq := 1 // [result] max sequence
	//pos := 0 // [result] finding position

	if len(tree) < 3 {
		return len(tree)
	}
	for i := 0; i <= len(tree)-2; i++ {

		if tree[i] == tree[i+1] {
			count++
			if seq < max+count {
				seq = max + count
			}
		} else {

			if i == len(tree)-2 && (pre == -1 || pre == tree[i+1]) {
				if seq < max+count+1 {
					seq = max + count + 1
				}
			} else if seq < max+count {
				seq = max + count
			}

			if pre == -1 || pre == tree[i+1] {
				max = max + count
			} else {
				max = count
			}

			count = 1
			pre = tree[i]
		}

		//pos = i - max + 2
	}
	return seq
}

func main() {

	tree := []int{4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 4, 2, 2, 4, 2, 4, 4, 2, 3, 2, 1, 1, 1, 2, 1, 2, 3, 4, 3, 2, 1}

	seq := generateLongestSequence(tree)

	fmt.Println(tree)
	fmt.Println("seq = ", seq)
}
