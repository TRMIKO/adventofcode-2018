package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func recorre(x int64, y int64, w int64, h int64, arr *[1000][1000]int) {
	for i := y; i < y+h; i++ {
		for j := x; j < x+w; j++ {
			//fmt.Println(i, j)
			arr[i][j]++
		}
	}
}
func verifica(x int64, y int64, w int64, h int64, arr [1000][1000]int) bool {
	for i := y; i < y+h; i++ {
		for j := x; j < x+w; j++ {
			if arr[i][j] > 1 {
				return false
			}
		}
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func row(str []string) (int64, int64, int64, int64, int64) {
	v0, _ := strconv.ParseInt(str[0], 0, 64)
	v1, _ := strconv.ParseInt(str[1], 0, 64)
	v2, _ := strconv.ParseInt(str[2], 0, 64)
	v3, _ := strconv.ParseInt(str[3], 0, 64)
	v4, _ := strconv.ParseInt(str[4], 0, 64)
	return v0, v1, v2, v3, v4
}

func main() {
	var arr [1000][1000]int

	// example
	// recorre(1, 3, 4, 4, &arr)
	// recorre(3, 1, 4, 4, &arr)
	// recorre(5, 5, 2, 2, &arr)
	// if verifica(1, 3, 4, 4, arr) {
	// 	fmt.Println(1)
	// }
	// if verifica(3, 1, 4, 4, arr) {
	// 	fmt.Println(2)
	// }
	// if verifica(5, 5, 2, 2, arr) {
	// 	fmt.Println(3)
	// }
	file, err := os.Open("../input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		r, _ := regexp.Compile("[0-9]+")
		txt := r.FindAllString(scanner.Text(), -1)
		_, v1, v2, v3, v4 := row(txt)

		recorre(v1, v2, v3, v4, &arr)

	}
	file.Seek(0, 0)

	scanner = bufio.NewScanner(file)

	defer file.Close()
	for scanner.Scan() {
		r, _ := regexp.Compile("[0-9]+")
		txt := r.FindAllString(scanner.Text(), -1)
		v0, v1, v2, v3, v4 := row(txt)

		if verifica(v1, v2, v3, v4, arr) {
			fmt.Println(v0)
		}
	}

}
