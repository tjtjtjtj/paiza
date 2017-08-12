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
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")
	day := strings.Split(s[0], "/")
	time := strings.Split(s[1], ":")
	day_int, _ := strconv.Atoi(day[1])
	time_int, _ := strconv.Atoi(time[0])
	l := time_int / 24
	day_int = day_int + l
	m := time_int % 24
	fmt.Printf("%s/%02d %02d:%s", day[0], day_int, m, time[1])
}
