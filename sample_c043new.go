package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	id          int
	item_counts int
	max_flag    bool
	last_flag   bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	used_item_count, _ := strconv.Atoi(scanner.Text())
	persons := make([]Person, used_item_count)
	scanner.Scan()
	item_ids := strings.Split(scanner.Text(), " ")

	var s int
	var id_number int
	for _, v := range item_ids {
		s, _ = strconv.Atoi(v)
		for i := 0; i < used_item_count; i++ {
			if persons[i].id == 0 {
				persons[i].id = s
				persons[i].item_counts++
				id_number++
				break
			} else if persons[i].id == s {
				persons[i].item_counts++
				break
			}
		}
	}

	max := persons[0].item_counts
	persons[0].max_flag = true
	persons[0].last_flag = true
	for i := 1; i < id_number; i++ {
		if max < persons[i].item_counts {
			persons[i].max_flag = true
			persons[i].last_flag = true
			max = persons[i].item_counts
			for j := i - 1; j >= 0; j-- {
				persons[j].max_flag = false
				persons[j].last_flag = false
			}
		} else if max == persons[i].item_counts {
			persons[i].max_flag = true
			persons[i].last_flag = true
			for j := i - 1; j >= 0; j-- {
				persons[j].last_flag = false
			}
		}
	}

	var max_flag_count int
	for i := 0; i < id_number; i++ {
		if persons[i].max_flag == true {
			max_flag_count++
		}
	}

	output := make([]int, max_flag_count)
	for i, j := 0, 0; i < id_number; i++ {
		if persons[i].max_flag == true {
			output[j] = persons[i].id
			j++
		}
	}
	sort.Ints(output)

	for i := 0; i < max_flag_count; i++ {
		fmt.Print(output[i])
		if i != max_flag_count-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}
