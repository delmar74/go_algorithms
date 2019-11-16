package main

import (
	"fmt"
)

func generateLongestSequence(tree []int) (int, int) {
	pre := 0
	count := 1
	max := 1 // local maximum
	seq := 1 // [result] max sequence
	pos := 0 // [result] finding position

	for j := 0; j <= len(tree)-2; j++ {
		if tree[j] == tree[j+1] {
			count++
		} else {
			max = max + count

			if seq < max {
				seq = max
				pos = j - max + 1
			}

			if pre != tree[j+1] {
				max = count
			}

			count = 1
			pre = tree[j]
		}
	}
	return seq, pos
}

func main() {
	tree := []int{2, 4, 4, 2, 2, 4, 2, 4, 4, 2, 3, 2, 1, 1, 1, 2, 1, 2, 3, 4, 3, 2, 1}
	seq, pos := generateLongestSequence(tree)

	fmt.Println(tree)
	fmt.Println("seq = ", seq, ", pos = ", pos)
}
