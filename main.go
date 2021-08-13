package main

import "fmt"

// soal 2.1 faktor bilangan
// func main() {
// 	//hasil perulangan nanti print dari terkecil, maka perulangan for dari kecil dan terus bertambah
// 	var bil int

// 	fmt.Println("masukkan bilangan : ...")

// 	fmt.Scanf("%d", &bil)
// 	fmt.Println("daftar faktor bilangannya adalah :")

// 	//perulangan yang dimaksud
// 	for i := 1; i <= bil; i++ {
// 		if bil%i == 0 {
// 			fmt.Println(i)
// 		}
// 	}

// }

// soal 2.2 faktor bilangan
// func main() {
// 	//karena hasil akhir yg diharapkan terbalik urutannya, maka perulangan for dari yg terbesar dan terus berkurang
// 	var bil int

// 	fmt.Println("masukkan bilangan : ...")

// 	fmt.Scanf("%d", &bil)
// 	fmt.Println("daftar faktor bilangannya adalah (sorting terbalik) :")
// 	//perulangan yg dimaksud
// 	for i := bil; i >= 1; i-- {
// 		if bil%i == 0 {
// 			fmt.Println(i)
// 		}
// 	}

// }
//---------------------------------------------//
// soal 3
func primeNum(number int) bool {
	//jumlah bilangan yg habis untuk membagi input
	var kali int

	//hasil akhir
	var hasil bool

	//perulangan sebanyak jumlah input, apabila habis dibagi maka variabel kali akan bertambah
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			kali++
		}
	}

	// karena bil prima adalah bilangan yg habis dibagi hanya dengan 1 dan bil itu sendiri, maka jumlah bil di variabel kali harus 2 agar bisa true
	if kali == 2 {
		hasil = true
	} else if kali != 2 {
		hasil = false
	}
	return hasil
}

func main() {
	fmt.Println(primeNum(11)) //true
	fmt.Println(primeNum(13)) //true
	fmt.Println(primeNum(17)) //true
	fmt.Println(primeNum(20)) //false
	fmt.Println(primeNum(35)) //false
}

//---------------------------------------//
// soal 4
// func palindrome(input string) bool {
// 	//hasil akhir yg diharapkan (true/false) defaultnya false
// 	var kembalian bool

// 	//menampung jumlah karakter yang sama
// 	var nilai int

// 	//panjang dari inputan ( jumlah karakternya )
// 	var karakter int = len(input)

// 	for i := 0; i < karakter; i++ {
// 		//misalkan karakter pertama dan terakhir sama, maka variabel nilai bertambah
// 		if input[i] == input[karakter-(i+1)] {
// 			nilai++
// 		}
// 	}
// 	//jika jumlah karakter input dan jumlah karakter yg sama sesuai maka true
// 	if nilai == karakter {
// 		kembalian = true
// 	}
// 	return kembalian
// }

// func main() {
// 	fmt.Println(palindrome("civic"))
// 	fmt.Println(palindrome("katak"))
// 	fmt.Println(palindrome("kasur rusak"))
// 	fmt.Println(palindrome("kupu-kupu"))
// 	fmt.Println(palindrome("lion"))
// }

//---------------------------------------//
// soal 5
func pangkat(base, pangkat int) int {
	//variabel untuk menampung hasil perkalian
	var hasil int = base

	//perkalian yang dilakukan dengan bilangan base sebanyak pangkat
	for i := 1; i < pangkat; i++ {
		hasil = hasil * base
	}
	return hasil
}

// func main() {

// 	fmt.Println(pangkat(2, 3))
// 	fmt.Println(pangkat(5, 3))
// 	fmt.Println(pangkat(10, 2))
// 	fmt.Println(pangkat(2, 5))
// 	fmt.Println(pangkat(7, 3))

// }

//-------------------------------------------//
// soal 3.1
// func playWithAsterik(n int) {
// 	//perulangan baris
// 	for i := 0; i < n; i++ {
// 		//perulangan kolom
// 		for spasi := n - i; spasi >= 0; spasi-- {
// 			//print sebanyak nilai spasi
// 			fmt.Print(" ")

// 		}
// 		for bintang := 0; bintang <= i; bintang++ {
// 			//print sebanyak nilai bintang (yang di print "* ")
// 			fmt.Print("*", " ")
// 		}
// 		//buat baris baru
// 		fmt.Print("\n")
// 	}
// }
// func main() {
// 	playWithAsterik(5)
// }

// func main() {
// 	var in int = 4
// 	var lamp bool = false

// 	for i := 1; i <= in; i++ {

// 		if i == 1 {
// 			lamp = true
// 		} else {
// 			if in%i == 0 {
// 				if lamp == true {
// 					lamp = false
// 				} else if lamp == false {
// 					lamp = true
// 				}
// 			}
// 		}
// 	}

// 	fmt.Println(lamp)
// }

//============== day3 ================//
//soal 1
// func bilPrim(input uint) bool {
// 	if input == 2 {
// 		return true
// 	} else if input <= 13 {
// 		for i := 1; i < int(input); i++ {
// 			if input%uint(i) == 0 {
// 				return true
// 			}
// 		}
// 	} else {
// 		for i := 2; i <= 13; i++ {
// 			if input%uint(i) == 0 {
// 				return false
// 			}
// 		}
// 	}
// 	return true

// }
// func main() {
// 	fmt.Println(bilPrim(1000000007))
// 	fmt.Println(bilPrim(1500450271))
// 	fmt.Println(bilPrim(1000000000))
// 	fmt.Println(bilPrim(10000000019))
// 	fmt.Println(bilPrim(10000000033))
// 	fmt.Println(bilPrim(7))
// 	fmt.Println(bilPrim(3))
// }

//soal2
// func pangkat(x, n int) int {
// 	var res int = x * x
// 	var kali2 int = x * x

// 	for i := 0; i < (n-2)/2; i++ {
// 		res *= kali2
// 	}
// 	if (n/2)*2 != n {
// 		res *= x
// 	}

// 	return res
// }

// func main() {
// 	fmt.Println(pangkat(2, 3))
// 	fmt.Println(pangkat(7, 2))
// 	fmt.Println(pangkat(10, 5))
// 	fmt.Println(pangkat(17, 6))
// 	fmt.Println(pangkat(5, 3))
// }

//soal 3 -- belum
// func gabungTanpaDup(arrA, arrB []string) []string {
// 	var res = make([]string, 0)
// 	for

// 	return res

// }
// func main() {
// 	fmt.Println(gabungTanpaDup([]string{"apel", "anggur"}, []string{"lemon", "leci", "nanas"}))
// 	fmt.Println(gabungTanpaDup([]string{"samsung", "apple"}, []string{"apple", "sony", "xiaomi"}))
// 	fmt.Println(gabungTanpaDup([]string{"football", "basketball"}, []string{"basketball", "football"}))
// }

//soal 4
// import "strconv"

// func sekali(angka string) int  {
// 	var arrInt []int = make([]int, 0)

// }

// func main()  {
// 	fmt.Println(sekali("1234123"))
// 	fmt.Println(sekali("76523752"))
// 	fmt.Println(sekali("12345"))
// 	fmt.Println(sekali("1122334455"))
// 	fmt.Println(sekali("0872504"))
// }

//soal 5
// func pairSum(arr []int, target int) []int {
// 	var res []int
// 	// for index, value := range arr {
// 	// 	var pengurangan int
// 	// 	for i := 0; i < len(arr)/2; i++ {
// 	// 		pengurangan = target - arr[i]
// 	// 		if pengurangan == value {
// 	// 			res = append(res, i, index)
// 	// 		}
// 	// 	}
// 	// }

// 	return res
// }

// func main() {
// 	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6))
// 	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))
// 	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12))
// 	fmt.Println(pairSum([]int{1, 4, 6, 8}, 10))
// 	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6))

// }
