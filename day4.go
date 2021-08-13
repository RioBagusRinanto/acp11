// package main

// import "fmt"

// //============ soal 1
// // import (
// // 	"fmt"
// // 	"strings"
// // )

// // func bandingkan(a, b string) string {
// // 	var result string

// // 	if len(a) > len(b) {
// // 		res := strings.Contains(a, b)
// // 		if res == true {

// // 			result = b
// // 		}
// // 	} else if len(a) < len(b) {
// // 		res := strings.Contains(b, a)
// // 		if res == true {

// // 			result = a
// // 		}
// // 	}

// // 	return result

// // }
// // func main() {
// // 	fmt.Println(bandingkan("aka", "akashi"))
// // 	fmt.Println(bandingkan("kangooro", "kang"))
// // 	fmt.Println(bandingkan("ki", "kijang"))
// // 	fmt.Println(bandingkan("kupu-kupu", "kupu"))
// // 	fmt.Println(bandingkan("ilalang", "ila"))
// // }

// //================== soal 2
// // func enco(offset int, input string) string {
// // 	offset = (offset%26 + 26) % 26 //[0,25]
// // 	b := make([]byte, len(input))

// // 	for i := 0; i < len(input); i++ {
// // 		t := input[i]
// // 		var a int

// // 		switch {
// // 		case 'a' <= t && t <= 'z':
// // 			a = 'a'
// // 		// case 'A' <= t && t <= 'Z':
// // 		// 	a = 'A'
// // 		default:
// // 			b[i] = t
// // 			continue
// // 		}
// // 		b[i] = byte(a + ((int(t)-a)+offset)%26)

// // 	}
// // 	return string(b)
// // }

// // func main() {
// // 	fmt.Println(enco(3, "abc"))
// // 	fmt.Println(enco(2, "alta"))
// // 	fmt.Println(enco(10, "alterraacademy"))
// // 	fmt.Println(enco(1, "abcdefghijklmnopqrstuvwxyz"))
// // 	fmt.Println(enco(1000, "abcdefghijklmnopqrstuvwxyz"))
// // }

// //================ soal 3
// // func arrayunik(arrA, arrB []int) []int {

// // 	var hasil []int = make([]int, 0)
// // 	res := 0

// // 	for i := 0; i < len(arrA); i++ {
// // 		res = 0

// // 		for j := 0; j < len(arrB); j++ {
// // 			if arrB[j] != arrA[i] {
// // 				res++
// // 			}
// // 		}
// // 		if res == len(arrB) {
// // 			hasil = append(hasil, arrA[i])
// // 		}
// // 	}
// // 	return hasil

// // }

// // func main() {
// // 	fmt.Println(arrayunik([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))
// // 	fmt.Println(arrayunik([]int{10, 20, 30, 40}, []int{5, 10, 15, 59}))
// // 	fmt.Println(arrayunik([]int{1, 3, 7}, []int{1, 3, 5}))
// // 	fmt.Println(arrayunik([]int{3, 8}, []int{2, 8}))
// // 	fmt.Println(arrayunik([]int{1, 2, 3}, []int{3, 2, 1}))
// // }

// //========= soal 4
// // func maxSum(k int, arr []int) int {
// // 	// res1 := 0
// // 	// res2 := 0
// // 	// for i := 0; i < len(arr); i++ {
// // 	// 	temp := arr[i]
// // 	// 	for j := i + 1; j < i+k; j++ {
// // 	// 		if j > len(arr)-1 {
// // 	// 			break
// // 	// 		}
// // 	// 		temp += arr[j]
// // 	// 	}
// // 	// 	res1 = temp
// // 	// 	if res1 > res2 {
// // 	// 		res2 = res1
// // 	// 	}
// // 	// }
// // 	// return res2
// // 	res := 0
// // 	for i := 0; i < len(arr); i++ {
// // 		if i+k < len(arr) {
// // 			temp := 0
// // 			for j := 0; j < k; j++ {
// // 				temp += arr[i+j]
// // 			}
// // 			if temp > res {
// // 				res = temp
// // 			}
// // 		}
// // 	}
// // 	return res
// // }

// // func main() {
// // 	fmt.Println(maxSum(3, []int{2, 1, 5, 1, 3, 2}))
// // 	fmt.Println(maxSum(2, []int{2, 3, 4, 1, 5}))
// // 	fmt.Println(maxSum(2, []int{2, 1, 4, 1, 1}))
// // 	fmt.Println(maxSum(3, []int{2, 1, 4, 1, 1}))
// // 	fmt.Println(maxSum(4, []int{2, 1, 4, 1, 1}))
// // }

// //============= soal 5
// func hapusdobel(arr []int) int {
// 	// res := 0
// 	// for i := 0; i < len(arr); i++ {
// 	// 	temp := arr[i]
// 	// 	temppos := 0
// 	// 	arrDobelVal := 0
// 	// 	for j := 0; j < len(arr); j++ {
// 	// 		if i == j {
// 	// 			continue
// 	// 		} else if temp == arr[j] {
// 	// 			temppos++
// 	// 		}

// 	// 	}
// 	// 	if temppos == 0 {
// 	// 		res++
// 	// 	} else if temppos != 0 {
// 	// 		arrDobelVal++
// 	// 	}

// 	// 	fmt.Println(arrDobelVal)

// 	// }
// 	// return res
// 	res := 1
// 	for i := 0; i < len(arr)-1; i++ {

// 		if arr[i] < arr[i+1] {
// 			res++
// 		}

// 	}

// 	return res
// }

// func main() {
// 	fmt.Println(hapusdobel([]int{2, 3, 3, 3, 6, 9, 9}))
// 	fmt.Println(hapusdobel([]int{2, 3, 4, 5, 6, 9, 9}))
// 	fmt.Println(hapusdobel([]int{2, 2, 2, 11}))
// 	fmt.Println(hapusdobel([]int{2, 2, 2, 11}))
// 	fmt.Println(hapusdobel([]int{1, 2, 3, 11, 11}))
// }
