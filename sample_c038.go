package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type machine struct {
	number     int
	bunnpaisuu int
	amari      int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	first_flag := true
	var m1, m2 machine
	var s []string
	var total_num int
	var machine_num int

	for scanner.Scan() {
		if first_flag == true {
			s = strings.Split(scanner.Text(), " ")
			total_num, _ = strconv.Atoi(s[1])
			first_flag = false
		} else {
			machine_num++
			m1.bunnpaisuu, _ = strconv.Atoi(scanner.Text())
			m1.amari = total_num % m1.bunnpaisuu
			if m2.bunnpaisuu == 0 {
				m2.amari = m1.amari
				m2.bunnpaisuu = m1.bunnpaisuu
				m2.number = machine_num
			} else if m1.amari < m2.amari {
				m2.amari = m1.amari
				m2.bunnpaisuu = m1.bunnpaisuu
				m2.number = machine_num
			} else if m1.amari == m2.amari && m1.bunnpaisuu > m2.bunnpaisuu {
				m2.bunnpaisuu = m1.bunnpaisuu
				m2.number = machine_num
			}
		}
	}
	fmt.Print(m2.number)
}
