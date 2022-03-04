package main

import "fmt"

func main() {
	// анонимная функция
	func(in string) {
		fmt.Println("anon func out:", in)
	}("nobody")

	printer := func(in string) {
		fmt.Println("printer outs:", in)
	}

	printer("as variable")

	// определяем тип функции
	type strFuncType func(string)

	// функция принимает коллбек
	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(printer)

	// функция возвращает замыкание
	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s \n", prefix, in)
		}
	}

	successLogger := prefixer("SUCCESS")
	successLogger("expected behaviour")
}
