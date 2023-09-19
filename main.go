package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	//NG Challenge 2 : Conditional 1
	var name string
	fmt.Println("Masukan Nama")
	fmt.Scanln(&name)

	// fmt.Println(name)

	randomInt := rand.Intn(100)
	fmt.Println(randomInt)
	switch {
	case randomInt > 80:
		fmt.Println("Selamat", name+" adna sangat beruntung")
	case randomInt > 60:
		fmt.Println("Selamat", name+" anda beruntung")
	case randomInt > 40:
		fmt.Println("Mohon maaf", name+" anda kurang beruntung")
	default:
		fmt.Println("Mohon maaf", name+" anda sial")
	}

	//NG Challenge 2 : Conditional 2
	var nama string
	var umurStr string

	fmt.Println("Masukan Nama")
	fmt.Scanln(&nama)

	fmt.Println("Masukan Umur")
	fmt.Scanln(&umurStr)

	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		fmt.Println("umur bukan merupakan angka")
		os.Exit(1)
	}
	if umur < 0 || umur > 100 {
		fmt.Println("umur harus dalam rentang 0-100")
		os.Exit(1)
	}
	if umur > 18 {
		fmt.Println("Silakan masuk,", nama)
	} else {
		fmt.Println("Dilarang masuk, maksimal umur 18.")
	}
}

//strings string.TrimSpace
