package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var a int64
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		num, _ := strconv.ParseInt(scanner.Text(), 0, 32)
		//fmt.Println(num)
		a += num
	}
	fmt.Println(a)
}
