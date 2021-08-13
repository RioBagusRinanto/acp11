package main

// import (
// 	"fmt"
// )

//========== soal 1
// func swap(a, b *int) {
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	temp := *a
// 	*a = *b
// 	*b = temp

// }

// func main() {
// 	a := 10
// 	b := 20

// 	swap(&a, &b)

// 	fmt.Println(a, b)

// }

//======= soal 2
// func getMinMax(num ...*int) (min int, max int) {
// 	max = 0
// 	min = int(^uint(0) >> 1)
// 	for _, val := range num {
// 		if *val > max {
// 			max = *val
// 		}
// 		if *val < min {
// 			min = *val
// 		}
// 	}

// 	return min, max
// }

// func main() {
// 	var a1, a2, a3, a4, a5, a6, min, max int
// 	fmt.Scan(&a1)
// 	fmt.Scan(&a2)
// 	fmt.Scan(&a3)
// 	fmt.Scan(&a4)
// 	fmt.Scan(&a5)
// 	fmt.Scan(&a6)
// 	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
// 	fmt.Println("nilai min ", min)
// 	fmt.Println("nilai max ", max)
// }

// ============= soal 3
// type Student struct {
// 	nama []string
// 	skor []int
// }

// func (s Student) Average() float64 {
// 	total := 0
// 	avg := 0.0
// 	for _, score := range s.skor {
// 		total += score
// 	}
// 	avg = float64(total) / float64(len(s.skor))

// 	return avg

// }
// func (s Student) Min() (min int, nama string) {
// 	resMin := int(^uint(0) >> 1)

// 	for posScore, score := range s.skor {
// 		if score < resMin {
// 			resMin = score
// 			nama = s.nama[posScore]
// 		}

// 	}
// 	min = resMin
// 	return
// }
// func (s Student) Max() (max int, nama string) {
// 	// resMax := 0

// 	for posScore, score := range s.skor {

// 		if score > max {
// 			max = score
// 			nama = s.nama[posScore]
// 		}

// 	}
// 	// max = resMax
// 	return
// }

// func main() {
// 	var a = Student{}

// 	for i := 0; i < 6; i++ {
// 		var name string
// 		fmt.Print("input " + string(i) + " students's name : ")
// 		fmt.Scan(&name)
// 		a.nama = append(a.nama, name)

// 		var score int
// 		fmt.Print("input " + name + " score : ")
// 		fmt.Scan(&score)
// 		a.skor = append(a.skor, score)
// 	}

// 	fmt.Println("\n\naverage score student is : ", a.Average())
// 	scoreMax, nameMax := a.Max()
// 	fmt.Println("max score student is : ", nameMax, " (", scoreMax, ") ")
// 	scoreMin, nameMin := a.Min()
// 	fmt.Println("min score student is : ", nameMin, " (", scoreMin, ") ")
// }

//============ soal 4

// type student struct {
// 	name       string
// 	nameEncode string
// 	score      int
// }

// type chiper interface {
// 	Encode() string
// 	Decode() string
// }

// func (s *student) Encode() string {
// 	var nameEncode = ""

// 	alp := "abcdefghijklmnopqrstuvwxyz"
// 	pla := "zyxwvutsrqponmlkjihgfedcba"

// 	var mapAlpabet = map[int]int{}
// 	var mapPla = map[int]int{}
// 	ran := 122
// 	for i := 0; i < len(alp); i++ {
// 		mapAlpabet[i] = ran
// 		ran--
// 	}
// 	ran2 := 0
// 	for i := 122; 0 < len(pla); i-- {
// 		mapPla[ran2] = i
// 		ran2++
// 	}
// 	res := make([]int, 0)
// 	for _, val := range s.name {
// 		res = append(res, mapPla[int(val)])
// 	}
// 	fmt.Println(res)

// 	return nameEncode
// }

// func (s *student) Decode() string {
// 	var nameDecode = ""

// 	return nameDecode
// }

// func main() {
// 	var menu int
// 	var a = student{}
// 	var c chiper = &a

// 	fmt.Print("[1] encrypt \n[2] decrypt \nchoose yout menu ?")
// 	fmt.Scan(&menu)
// 	if menu == 1 {
// 		fmt.Print("\ninput nama murid : ")
// 		fmt.Scan(&a.name)
// 		fmt.Print("\nencode dari : " + a.name + " is : " + c.Encode())
// 	} else if menu == 2 {
// 		fmt.Print("\ninput encode nama murid : ")
// 		fmt.Scan(&a.nameEncode)
// 		fmt.Print("\ndecode dari : " + a.nameEncode + " is : " + c.Decode())
// 	} else {
// 		fmt.Print("salah menu")
// 	}

// }
