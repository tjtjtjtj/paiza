package gotest

import (
	"testing"
)

func Testdead_live(t *testing.T) {
	bill := new(student)
	bill.subject = "l"
	bill.subject_scores = []int{10, 20, 30, 40, 50}

	if bill.dead_live() {
		t.Error("合計の計算がおかしい")
	}
}
