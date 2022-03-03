package main

import "fmt"

func main() {
	// цикл без условия, while(true) OR for(;;;)
	for {
		fmt.Println("loop iteration")
		break
	}

	// цикл без условия, while(isRun)
	isRun := true
	if isRun {
		fmt.Println("loop iteration with condition")
		isRun = false
	}

	// цикл с условием и блоком инициализации
	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration", i)
		if i == 1 {
			continue
		}
	}

	// операции по slice
	sl := []int{1, 2, 3}
	idx := 0

	for idx < len(sl) {
		fmt.Println("while-stype loop, idx:", idx, "value:", sl[idx])
		idx++
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c-style loop", i, sl[i])
	}

	for index := range sl {
		fmt.Println("range slice by index", index)
	}

	for index, value := range sl {
		fmt.Println("range slice by idx-value", index, value)
	}

	// операции по map
	profile := map[int]string{1: "Vasily", 2: "Romanov"}

	for key := range profile {
		fmt.Println("range map by key", key)
	}

	for key, value := range profile {
		fmt.Println("range map by key-value", key, value)
	}

	for _, value := range profile {
		fmt.Println("range map by value", value)
	}

	str := "Привет, Мир!"
	for pos, char := range str {
		fmt.Printf("%#U at pos %d\n", char, pos)
	}
}
