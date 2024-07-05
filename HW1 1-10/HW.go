package main

import (
	"fmt"
	"math"
	"strings"
)

type Rectangle struct {
	width  float64
	height float64
}

func task1() {
	fmt.Println("Привет, мир!")
}
func task2() {
	fmt.Println("Введите два числа для задания 2")
	var x, y int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	fmt.Println(x + y)
}
func task3() {
	fmt.Println("Введите число для задания 3")
	var x int
	fmt.Scanf("%d", &x)
	if x%2 == 0 {
		fmt.Println("Чётное")
	} else {
		fmt.Println("Нечётное")
	}
}
func task4() {
	fmt.Println("Введите три числа для задания 4")
	var x, y, z int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	fmt.Scanf("%d", &z)
	fmt.Println(max(x, y, z))
}
func task5() {
	fmt.Println("Введите число для задания 5")
	var x int
	fmt.Scanf("%d", &x)
	for i := x - 1; i != 0; i-- {
		x *= i
	}
	fmt.Println(x)
}
func task6() {
	fmt.Println("Введите букву для задания 6")
	vowels := "aeiouAEIOU"
	var ch string
	fmt.Scanln(&ch)

	result := "consonant"
	if strings.Contains(vowels, ch) {
		result = "vowel"
	}
	fmt.Println("%s is %s\n", ch, result)
}
func task7() {
	fmt.Println("Введите число для задания 7")
	var x int
	fmt.Scan(&x)
	if x < 2 {
		fmt.Println("Нет простых чисел < 2")
		return
	}

	isPrime := make([]bool, x+1)
	for i := 2; i <= x; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if isPrime[i] {
			for j := i * i; j <= x; j += i {
				isPrime[j] = false
			}
		}
	}

	for i := 2; i <= x; i++ {
		if isPrime[i] {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
func task8() {
	fmt.Println("Введите строку для задания 8")
	var x string
	fmt.Scanln(&x)
	result := ""
	for _, v := range x {
		result = string(v) + result
	}
	fmt.Println(result)
}
func task9(arr []int) {
	fmt.Println("Задание 9")
	sum := 0
	for _, num := range arr {
		sum += num
	}
	fmt.Printf("Сумма элементов: %d\n", sum)

}
func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	task6()
	task7()
	task8()

	fog := []int{5, 2, 7, 1, 9, 3}
	task9(fog)

	//task10
	var x Rectangle
	fmt.Print("Введите ширину и длину для задания 10: ")
	fmt.Scan(&x.width)
	fmt.Scan(&x.height)
	res := x.Area()
	fmt.Println(res)
}
