package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type country struct {
	gold   int
	silver int
	bronze int
}

type countries []country

func lessf(c countries) func(int, int) bool {
	return func(i, j int) bool {
		if c[i].gold < c[j].gold {
			return false
		}

		if c[i].gold == c[j].gold {
			if c[i].silver < c[j].silver {
				return false
			}
			if c[i].silver == c[j].silver {
				if c[i].bronze < c[j].bronze {
					return false
				}
			}
		}
		return true
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	first_flag := true
	member := 0
	i := 0
	var c countries

	for scanner.Scan() {
		if first_flag == true {
			member, _ = strconv.Atoi(scanner.Text())
			c = make([]country, member)
			first_flag = false
		} else {
			s := strings.Split(scanner.Text(), " ")
			l, _ := strconv.Atoi(s[0])
			m, _ := strconv.Atoi(s[1])
			n, _ := strconv.Atoi(s[2])
			c[i] = country{gold: l, silver: m, bronze: n}
			i++
		}
	}
	lessfunc := lessf(c)
	sort.Slice(c, lessfunc)

	for i := 0; i < member; i++ {
		fmt.Println(c[i].gold, c[i].silver, c[i].bronze)
	}
}
