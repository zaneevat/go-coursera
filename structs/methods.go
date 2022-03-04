package main

import "fmt"

type PersonTwo struct {
	Id   int
	Name string
}

// UpdateName не изменит оригинальной структуры, для которой вызван
func (p PersonTwo) UpdateName(name string) {
	p.Name = name
}

// SetName изменяет оригинальную структуру
func (p *PersonTwo) SetName(name string) {
	p.Name = name
}

type AccountTwo struct {
	Id   int
	Name string
	PersonTwo
}

type MySlice []int

func (s1 *MySlice) Add(val int) {
	*s1 = append(*s1, val)
}

func (s1 *MySlice) Count() int {
	return len(*s1)
}

func main() {
	pers := PersonTwo{1, "Maxim"}
	pers.SetName("Maxim Semkin")
	fmt.Printf("updated person: %#v\n", pers)

	var acc = AccountTwo{
		Id:   1,
		Name: "smax",
		PersonTwo: PersonTwo{
			Id:   2,
			Name: "Maxim Semkin",
		},
	}

	acc.SetName("maxim semkin")
	fmt.Printf("%#v \n", acc)

	s1 := MySlice([]int{1, 2})
	s1.Add(4)
	fmt.Println(s1.Count(), s1)
}
