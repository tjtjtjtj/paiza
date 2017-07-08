package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var total_num int
		var m int
		for i := 0; i < len(s); i++ {
			m, _ = strconv.Atoi(s[i])
			total_num = total_num + m
		}
		var average float64
		average = float64(total_num) / float64(len(s))
		fmt.Println(strconv.FormatFloat(average, 'f', 1, 32))
	}
}
