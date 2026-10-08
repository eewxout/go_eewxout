package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

//  1
func task1() {
	fmt.Println(" Задача 1 ")
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, n := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("Квадрат %d = %d\n", num, num*num)
		}(n)
	}

	wg.Wait()
}

// 2
func task2() {
	fmt.Println("\n Задача 2")
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, t := range temps {
		var key int
		if t >= 0 {
			key = int(t) / 10 * 10
		} else {
			key = int(t)/10*10 - 10
			if int(t)%10 == 0 {
				key = int(t)
			}
		}
		groups[key] = append(groups[key], t)
	}

	for k, v := range groups {
		fmt.Printf("%d: %v\n", k, v)
	}
}

// 3 
func task3() {
	fmt.Println("\n Задача 3")
	input := "главрыба — абырвалг"

	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Println(string(runes))
}

// 4
func task4() {
	fmt.Println("\n Задача 4 ")
	input := "snow dog sun"

	words := strings.Fields(input)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	fmt.Println(strings.Join(words, " "))
}

//   5 
func task5() {
	fmt.Println("\n Задача 5 ")
	testStrings := []string{"abcd", "abCda", "Hello", "world"}

	for _, s := range testStrings {
		fmt.Printf("%q -> %v\n", s, isUnique(s))
	}
}

func isUnique(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		r = unicode.ToLower(r)
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}


func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// демо чтение для задач
func taskFromStdin() {
	fmt.Println("\n Ввод из stdin задачи 3 - 5 ")
	input := readInput()

	// 3
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println("Перевёрнутая строка:", string(runes))

	// 4
	words := strings.Fields(input)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	fmt.Println("Перевёрнутые слова:", strings.Join(words, " "))

	// 5
	fmt.Println("Все символы уникальны:", isUnique(input))

	_ = strconv.Itoa 
}

func main() {
	task1()
	task2()
	task3()
	task4()
	task5()
	taskFromStdin()
}