package main

import "testing"

func TestName(t *testing.T) {

	tree := []int{2, 4, 4, 4, 2, 3, 2, 1, 1, 1, 2, 1, 2, 3, 4, 3, 2, 1}
  seq := 7
  pos := 6


	res_seq, res_pos := generateLongestSequence(tree)
	if res_seq != seq || res_pos != pos {
		t.Error("Respone from generateLongestSequence is unexpected value")
	}
}
