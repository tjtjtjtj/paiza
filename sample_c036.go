package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type round struct {
	no1 int
	no2 int
}

func main() {
	var tour = make([]round, 3)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s1 := strings.Split(scanner.Text(), " ")
	tour[0].no1, _ = strconv.Atoi(s1[0])
	tour[0].no2, _ = strconv.Atoi(s1[1])

	scanner.Scan()
	s2 := strings.Split(scanner.Text(), " ")
	tour[1].no1, _ = strconv.Atoi(s2[0])
	tour[1].no2, _ = strconv.Atoi(s2[1])

	scanner.Scan()
	m := strings.Split(scanner.Text(), " ")
	var mint [4]int
	for k, v := range m {
		mint[k], _ = strconv.Atoi(v)
	}

	if mint[tour[0].no1-1] < mint[tour[0].no2-1] {
		tour[2].no1 = tour[0].no1
	} else {
		tour[2].no1 = tour[0].no2
	}
	if mint[tour[1].no1-1] < mint[tour[1].no2-1] {
		tour[2].no2 = tour[1].no1
	} else {
		tour[2].no2 = tour[1].no2
	}

	scanner.Scan()
	l := strings.Split(scanner.Text(), " ")
	l1, _ := strconv.Atoi(l[0])
	l2, _ := strconv.Atoi(l[1])
	if tour[2].no1 < tour[2].no2 {
		if l1 < l2 {
			fmt.Println(tour[2].no1)
			fmt.Println(tour[2].no2)
		} else {
			fmt.Println(tour[2].no2)
			fmt.Println(tour[2].no1)
		}
	} else {
		if l1 < l2 {
			fmt.Println(tour[2].no2)
			fmt.Println(tour[2].no1)
		} else {
			fmt.Println(tour[2].no1)
			fmt.Println(tour[2].no2)
		}
	}
}
