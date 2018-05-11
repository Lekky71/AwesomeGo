package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Printf("Hello world!\n")
	const complexNum = 5+4i
	fmt.Printf("Value is: %v",complexNum)
	test()

}

var chineseHello string
var englishHello string

func test(){
	//no, yes, maybe := "no", "yes", "maybe"
	chineseHello = "NiHao"
	englishHello = "Hello"
	//to change the index of a string
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)
	//end
	//String concatenation
	fmt.Printf("%s\n", chineseHello+ " "+ englishHello)
	//multi-line strings
	m := `hi
			,how are you?`
	fmt.Printf("%s\n", m)

	//error types
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}

//var name = "Leke"
//var age = 85
//var score = 67.54
//var isActive bool
//isActive = true

const (
	name = "Leke"
	age = 76
	score = 76.23
)

//var (
//	name string
//	age int
//	score float32
//)

const (
	x = iota //x == 0
	y = iota //y == 1
	z = iota //z = 2
	w   //w = 3
)

const v = iota //v == 0

const (
	e, f, g = iota, iota, iota //e == f == g == 0
)