package main

//=========== soal 1 belum
// func SimpleEquation(a, b, c int) {
// 	z := a / 2
// 	x := 0
// 	y := 0
// 	for i := z - 1; i >= 1; i-- {
// 		if z-i > 1 {
// 			y = i
// 			x = z - i
// 			if x+y+z == a && x*y*z == b && (x*x)+(y*y)+(z*z) == c {
// 				break
// 			} else if x+y+z != a && x*y*z != b && (x*x)+(y*y)+(z*z) != c {
// 				z--
// 				continue
// 			}
// 		}
// 	}
// 	if x+y+z != a || x*y*z != b || (x*x)+(y*y)+(z*z) != c {
// 		fmt.Println("no solution")
// 	} else if x+y+z == a && x*y*z == b && (x*x)+(y*y)+(z*z) == c {
// 		fmt.Println("x : ", x, " ,y : ", y, " , z : ", z)
// 	}

// 	// fmt.Println("x+y+z ", x+y+z)
// 	// fmt.Println("x.y.z ", x*y*z)
// 	// fmt.Println("x^+y^+z^ ", (x*x)+(y*y)+(z*z))

// }

// func main() {
// 	SimpleEquation(1, 2, 3)     //no solution
// 	SimpleEquation(6, 6, 14)    //1 2 3
// 	SimpleEquation(16, 126, 94) //367
// }

//=============== soal 2 selesai
// func moneyChange(money int) []int {
// 	var pecahan []int = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
// 	var res []int = []int{}
// 	for i := 0; i < len(pecahan); i++ {
// 		temp := money - pecahan[i]
// 		if temp > 0 {
// 			// fmt.Println(pecahan[i])
// 			res = append(res, pecahan[i])
// 			money = money - pecahan[i]
// 			if money >= pecahan[i] {
// 				i--
// 			}
// 		}
// 		if temp == 1 {
// 			temp = temp - 1
// 			res = append(res, 1)
// 		}
// 		if money == 0 {
// 			break
// 		}
// 	}
// 	return res

// }

// func main() {
// 	fmt.Println(moneyChange(123))
// 	fmt.Println(moneyChange(432))
// 	fmt.Println(moneyChange(543))
// 	fmt.Println(moneyChange(7752))
// 	fmt.Println(moneyChange(15321))
// }

//================ soal 3 kurang sorting
// func DragonOfLoowater(dragonHead, knightHeight []int) {
// 	// sort.Ints(dragonHead)
// 	// sort.Ints(knightHeight)
// 	for i := 0; i < len(dragonHead); i++ {
// 		for j := i + 1; j < len(dragonHead); j++ {
// 			if dragonHead[j] < dragonHead[i] {
// 				temp := dragonHead[i]
// 				dragonHead[i] = dragonHead[j]
// 				dragonHead[j] = temp
// 			}
// 		}
// 	}

// 	for i := 0; i < len(knightHeight); i++ {
// 		for j := i + 1; j < len(knightHeight); j++ {
// 			if knightHeight[j] < knightHeight[i] {
// 				temp := knightHeight[i]
// 				knightHeight[i] = knightHeight[j]
// 				knightHeight[j] = temp
// 			}
// 		}
// 	}
// 	j := 0
// 	res := 0
// 	chop := 0
// 	for i := 0; i < len(knightHeight); i++ {
// 		if j == len(dragonHead) {
// 			break
// 		} else if dragonHead[j] > knightHeight[i] && j < len(dragonHead) {

// 			continue
// 		} else if dragonHead[j] <= knightHeight[i] {
// 			// fmt.Println(knightHeight[i])
// 			res = res + knightHeight[i]
// 			j++
// 			chop++
// 		}
// 	}
// 	if chop < len(dragonHead) {
// 		fmt.Println("Knight fall")
// 	} else if chop == len(dragonHead) {
// 		fmt.Println(res)
// 	}
// 	// fmt.Println(knightHeight)
// }

// func main() {
// 	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
// 	DragonOfLoowater([]int{5, 10}, []int{5})
// 	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})
// 	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5})

// }

//=========== soal 4 selesai
// func BinarySearch(array []int, x int) {

// 	min := 0
// 	max := len(array) - 1
// 	res := 0

// 	for i := min; i < max+1; i++ {
// 		med := (min + max) / 2
// 		// fmt.Print(res)
// 		if array[med] == x {
// 			res = med
// 		} else if array[med] != x {
// 			res = -1
// 		}
// 		if array[med] < x {
// 			min = med + 1
// 		} else if array[med] > x {
// 			max = med - 1
// 		}

// 	}

// 	fmt.Println(res)

// }

// func main() {
// 	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)                  //2
// 	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)                 //3
// 	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  //6
// 	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) //-1
// }

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> part 2
//=============== soal 1 done

// func fibo(n int) int {
// 	var res int
// 	if n <= 1 {
// 		res = n
// 		return res
// 	}
// 	res = fibo(n-1) + fibo(n-2)
// 	return res

// }

// func main() {

// 	fmt.Println(fibo(0))
// 	fmt.Println(fibo(1))
// 	fmt.Println(fibo(2))
// 	fmt.Println(fibo(3))
// 	fmt.Println(fibo(5))
// 	fmt.Println(fibo(6))
// 	fmt.Println(fibo(7))
// 	fmt.Println(fibo(9))
// 	fmt.Println(fibo(10))
// }

//============ soal 2 belum
// func fibo(n int) int {

// }

// func main() {

// 	fmt.Println(fibo(0))
// 	fmt.Println(fibo(1))
// 	fmt.Println(fibo(2))
// 	fmt.Println(fibo(3))
// 	fmt.Println(fibo(5))
// 	fmt.Println(fibo(6))
// 	fmt.Println(fibo(7))
// 	fmt.Println(fibo(9))
// 	fmt.Println(fibo(10))
// }

//================= soal 3 belom
// func frog(jumps []int) int {

// }

// func main() {
// 	fmt.Println(frog([]int{10, 30, 40, 20}))         //30
// 	fmt.Println(frog([]int{30, 10, 60, 10, 60, 50})) //40
// }

//================= soal 4 belum
// func romanNum(value int) string {
// 	// var roman = map[int]string{1000: "M", 500: "D", 100: "C", 50: "L", 10: "X", 5: "V", 1: "I"}
// 	var pecahan []int = []int{1000, 500, 100, 50, 10, 5, 1}
// 	var roma []string = []string{"M", "D", "C", "L", "X", "V", "I"}
// 	var res string
// 	if value <= 0 {
// 		return "invalid number"
// 	} else {
// 		// for nilai, huruf := range roman {
// 		// 	if value >= nilai {
// 		// 		res += huruf
// 		// 		value = value - nilai
// 		// 	}
// 		// }
// 		for i := 0; i < len(pecahan); i++ {
// 			if value >= pecahan[i] {
// 				res = res + roma[i]
// 				value = value - pecahan[i]
// 				i--
// 			}
// 		}
// 	}
// 	return res

// }

// func main() {
// 	fmt.Println(romanNum(6))
// 	fmt.Println(romanNum(9))
// 	fmt.Println(romanNum(23))
// 	fmt.Println(romanNum(2021))
// 	fmt.Println(romanNum(1646))
// }
