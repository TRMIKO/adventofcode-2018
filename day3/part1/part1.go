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
func cuenta(arr [1000][1000]int) int {
	cont := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			//fmt.Println(i, j)
			if arr[i][j] > 1 {
				cont++
			}
		}
	}
	return cont
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func row(str []string) (int64, int64, int64, int64) {
	v1, _ := strconv.ParseInt(str[1], 0, 64)
	v2, _ := strconv.ParseInt(str[2], 0, 64)
	v3, _ := strconv.ParseInt(str[3], 0, 64)
	v4, _ := strconv.ParseInt(str[4], 0, 64)
	return v1, v2, v3, v4
}

func main() {
	var arr [1000][1000]int
	// example
	// recorre(1, 3, 4, 4, &arr)
	// recorre(3, 1, 4, 4, &arr)
	// recorre(5, 5, 2, 2, &arr)
	file, err := os.Open("../input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		r, _ := regexp.Compile("[0-9]+")
		txt := r.FindAllString(scanner.Text(), -1)
		v1, v2, v3, v4 := row(txt)

		recorre(v1, v2, v3, v4, &arr)

	}
	// for j := 0; j < 1000; j++ {
	// 	fmt.Println(arr[j])
	// }

	fmt.Println(cuenta(arr))

}
