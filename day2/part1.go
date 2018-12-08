package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func analiza(str string, num int) bool {

	for i := 0; i < len(str); i++ {
		if strings.Count(str, string(str[i])) == num {
			return true
		}
	}
	return false
}

func main() {
	par := [2]int{0, 0}
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		if analiza(str, 2) {
			par[0]++
		}

		if analiza(str, 3) {
			par[1]++
		}

	}
	fmt.Println(par)
	fmt.Println(par[0] * par[1])
}
