package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(nearestTen(1, -2))
	fmt.Println(nearestTen(11, 12))
	fmt.Println(nearestTen(8, 12))

	fmt.Println(nearestTenMany(1, 1, 9, 19, 11, 4, 1))
	fmt.Println(nearestTenMany(1, 1, 2, 3, 4, 1, 1, 1, 2, 2, 2))
	fmt.Println(nearestTenMany(1, 1, 2, 3, 4, 1))

	fmt.Println(repeatThree(1, 1, 1, 2, 3, 4, 1))
	fmt.Println(repeatThree(1, 1, 2, 3, 4, 1, 1, 1, 2, 2, 2))
	fmt.Println(repeatThree(1, 1, 2, 3, 4, 1))

	var repeatThreeIIFE = func(numbers ...int) bool {
		for i := 0; i < len(numbers)-2; i++ {
			if numbers[i] == numbers[i+1] && numbers[i] == numbers[i+2] {
				return true
			}
		}

		return false
	}(1, 1, 1, 2, 3, 4, 1)

	fmt.Println(repeatThreeIIFE)

	var findLastDigit = lastDigit(1, 20, 937)
	fmt.Println(findLastDigit())

	var organisasi1 = Organisasi{nama: "BRI", alamat: "Sudirman"}
	var employee1 = Employee{nik: "1234567890", nama: "raffi", umur: 25, divisi: "DDB", organisasi: organisasi1}
	var employee2 = Employee{nik: "0987654321", nama: "iffar", umur: 52, divisi: "BDD", organisasi: organisasi1}

	var employees = Employees{employee1, employee2}

	employees.reflectEmployees()
	fmt.Println(employees.compareOlder())
}

// Day 6 Exercises
func nearestTen(num1 int, num2 int) int {
	x := math.Abs(float64(num1) - 10)
	y := math.Abs(float64(num2) - 10)

	if x < y {
		return num1
	} else if y < x {
		return num2
	} else {
		return 0
	}
}

func nearestTenMany(numbers ...int) int {
	var difference = map[int]float64{}
	var leastDifference = 999
	var nearest = 0

	for _, v := range numbers {
		difference[v] = math.Abs(float64(v) - 10)
	}

	for key, value := range difference {
		if value < float64(leastDifference) {
			leastDifference = int(value)
			nearest = key
		} else if value == float64(leastDifference) {
			return 0
		}
	}

	return nearest
}

func repeatThree(numbers ...int) bool {
	for i := 0; i < len(numbers)-2; i++ {
		if numbers[i] == numbers[i+1] && numbers[i] == numbers[i+2] {
			return true
		}
	}

	return false
}

func lastDigit(x ...int) func() []int {

	return func() []int {
		var lasts = []int{}

		for _, v := range x {
			var vAbs = int(math.Abs(float64(v)))

			if vAbs%10 != 0 {
				lasts = append(lasts, vAbs%10)
			}
		}

		return lasts
	}
}

type Employee struct {
	nik        string
	nama       string
	umur       int
	divisi     string
	organisasi Organisasi
}

type Organisasi struct {
	nama   string
	alamat string
}

type Employees []Employee

func (emps *Employees) reflectEmployees() {
	for i, e := range *emps {
		fmt.Println("Karyawan #", i)
		reflectValue := reflect.ValueOf(e)

		if reflectValue.Kind() == reflect.Ptr {
			reflectValue = reflectValue.Elem()
		}

		reflectType := reflectValue.Type()

		for i := 0; i < reflectValue.NumField(); i++ {
			fmt.Println("Nama      : ", reflectType.Field(i).Name)
			fmt.Println("Tipe Data : ", reflectType.Field(i).Type)
			fmt.Println("Nilai     : ", reflectValue.Field(i))
			fmt.Println("")
		}
		fmt.Println("")
	}
}

func (emps Employees) compareOlder() Employee {
	var oldestAge = 0
	var oldest Employee

	for _, v := range emps {
		if v.umur > oldestAge {
			oldestAge = v.umur
			oldest = v
		}
	}

	return oldest
}
