package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type student struct {
	subject        string
	subject_scores [5]int
}

func (t *student) dead_live() bool {
	var total_socre int
	for _, v := range t.subject_scores {
		total_socre = total_socre + v
	}

	if total_socre >= 350 {
		if t.subject == "l" {
			if t.subject_scores[3]+t.subject_scores[4] >= 160 {
				return true
			}
		} else {
			if t.subject_scores[1]+t.subject_scores[2] >= 160 {
				return true
			}
		}
	} else {
		return false
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	number, _ := strconv.Atoi(scanner.Text())
	students := make([]student, number)

	for i := 0; i < number; i++ {
		scanner.Scan()
		n := strings.Split(scanner.Text(), " ")
		students[i].subject = n[0]
		for j := 1; j <= len(students[i].subject_scores); j++ {
			students[i].subject_scores[j-1], _ = strconv.Atoi(n[j])
		}
	}

	var live_number int
	for _, v := range students {
		if v.dead_live() {
			live_number++
		}
	}

	fmt.Println(live_number)
}
