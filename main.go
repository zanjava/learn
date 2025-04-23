package main

import (
	"fmt"
	"time"
)

const s string = "123"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	var str string = "initial"
	fmt.Println(str)
	var d = true
	fmt.Println(d)
	b := "ff"
	fmt.Println(b)
	var i, j int = 1, 2
	fmt.Println(i, j)

	fmt.Println(s)

	var c uint8 = 1
	fmt.Println(c)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	default:
		fmt.Println("Good afternoon")
	}

	whatAmi := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Println("i don’t know type:%T", t)
		}
	}

	whatAmi(true)
}
