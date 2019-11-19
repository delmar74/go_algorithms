package main

import (
	"fmt"
)

func generateLongestSequence(tree []int) (int, int) {
	pre := 0
	count := 1
	max := 0 // local maximum
	seq := 1 // [result] max sequence
	pos := 0 // [result] finding position

	for i := 0; i <= len(tree)-2; i++ {

		if tree[i] == tree[i+1] {
			count++

			if i == len(tree)-2 {
				max = max + count
				if seq < max {
					seq = max
					pos = i - max + 2
				}
			}
			fmt.Println("i=", i, ", len=", len(tree), ", value=", tree[i], ", count=", count, ", max=", max, ", seq=", seq, ", pos=", pos)
		} else {
			max = max + count

			if seq < max {
				seq = max
				if pre == 0 {
					seq = seq + 1
				}
				pos = i - max + 1
			}

			if pre != tree[i+1] {
				max = count
			}
			fmt.Println("i=", i, ", len=", len(tree), ", value=", tree[i], ", count=", count, ", max=", max, ", seq=", seq, ", pos=", pos, ", pre=", pre, ", tree[i+1]=", tree[i+1])
			count = 1
			pre = tree[i]
		}
	}
	return pos, seq
}

func main() {
	//tree := []int{2, 4, 4, 2, 2, 4, 2, 4, 4, 2, 3, 2, 1, 1, 1, 2, 1, 2, 3, 4, 3, 2, 1}
	//tree := []int{2, 4, 4, 2, 2, 3, 1, 4, 5, 4, 5}
	//tree := []int{1, 1, 1, 2,0}
	//tree := []int{1, 1, 1}
	tree := []int{1, 0, 1}
	//tree := []int{1,1}
	//tree := []int{1}
	//tree := []int{}
	pos, seq := generateLongestSequence(tree)

	fmt.Println(tree)
	fmt.Println("pos = ", pos, ", seq = ", seq)
}
