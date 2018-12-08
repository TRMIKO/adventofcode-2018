package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func compara(st1 string, st2 string) int {
	r := 0
	for i := 0; i < len(st1); i++ {
		if st1[i] != st2[i] {
			r++
		}
	}
	return r
}
func main() {
	file, err := os.Open("../input.txt")
	check(err)
	defer file.Close()
	var s []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s = append(s, scanner.Text())

	}

	min := 50
	var pos [2]int

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i == j {
				continue
			}
			res := compara(s[i], s[j])
			if res < min {
				min = res
				pos[0] = i
				pos[1] = j
			}
		}
	}
	fmt.Println(min)
	fmt.Println(s[pos[0]])
	fmt.Println(s[pos[1]])

}
