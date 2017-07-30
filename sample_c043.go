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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	first_flag := true
	member := 0
	i := 0
	var countries []country

	for scanner.Scan() {
		if first_flag == true {
			member, _ = strconv.Atoi(scanner.Text())
			countries = make([]country, member)
			first_flag = false
		} else {
			s := strings.Split(scanner.Text(), " ")
			l, _ := strconv.Atoi(s[0])
			m, _ := strconv.Atoi(s[1])
			n, _ := strconv.Atoi(s[2])
			countries[i] = country{gold: l, silver: m, bronze: n}
			i++
		}
	}

	sort.Slice(countries, func(i, j int) bool {
		if countries[i].gold < countries[j].gold {
			return false
		}

		if countries[i].gold == countries[j].gold {
			if countries[i].silver < countries[j].silver {
				return false
			}
			if countries[i].silver == countries[j].silver {
				if countries[i].bronze < countries[j].bronze {
					return false
				}
			}
		}
		return true
	})

	for i := 0; i < member; i++ {
		fmt.Println(countries[i].gold, countries[i].silver, countries[i].bronze)
	}
}
