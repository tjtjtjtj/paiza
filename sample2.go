package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		s := scanner.Text()
		n, _ := strconv.Atoi(s)
		for i:=0 ;i < n;i++{
			fmt.Print("*")
		}
	}
	if err := scanner.Err(); err != nil{
		fmt.Fprintln(os.Stderr,"reading standart input:",err)
	}
}
	
