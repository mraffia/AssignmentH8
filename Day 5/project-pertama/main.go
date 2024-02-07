package main

import (
	"fmt"
	"strings"
)

func main() {
	monyetMasalah("Senyum", "Cemberut")
	berkokok(true, 12)

	var arr [4][5]string = [4][5]string{
		{"alvin", "arif", "reza", "rinaldi", "nina"},      //0
		{"noel", "dilla", "rosa", "juan michel", "teguh"}, //1
		{"septyan", "farras", "giyanda", "afin", "azwar"}, //2
		{"dionysius", "rifki", "raffi", "zain"},           //3
	}
	kelas(arr)
}

// Day 5 Exercises
func monyetMasalah(monyet1 string, monyet2 string) {
	if strings.ToLower(monyet1) == strings.ToLower(monyet2) {
		fmt.Println("Masalah")
	} else {
		fmt.Println("Aman")
	}
}

func berkokok(isBerkokok bool, jam int) {
	if isBerkokok && (jam == 24 || (jam >= 0 && jam <= 3)) {
		fmt.Println("Saya kesambet")
	} else {
		fmt.Println("Saya tidur nyenyak")
	}
}

func kelas(arr [4][5]string) {
	for i, arr2 := range arr {
		for j, value := range arr2 {
			if value == "raffi" {
				fmt.Println("Iterator 1:", i, "|", "Iterator 2:", j)
				fmt.Println("Nama:", value, "|", "Array nama ditemukan:", arr2)
			}
		}
	}
}
