package main

import "fmt"

//type Person struct {
//	Name string
//	Age  int
//}
//
//const NUM int = 100
//
//const NUM1 = 100
//
//type Gender string
//
//const (
//	Male   Gender = "Male"
//	Female Gender = "Female"
//)
//
//func main() {
//	// 声明并初始化匿名结构体
//	person := struct {
//		Name string
//		Age  int
//	}{
//		Name: "张三",
//		Age:  30,
//	}
//
//	fmt.Printf("姓名: %s, 年龄: %d\n", person.Name, person.Age)
//
//	person1 := Person{
//		Name: "张三11",
//		Age:  30,
//	}
//
//	person1.Print()
//
//	// 匿名结构体作为字段
//	type Employee struct {
//		ID   int
//		Info struct {
//			Address string
//			Phone   string
//		}
//	}
//
//	var emp Employee
//	emp.ID = 1001
//	emp.Info.Address = "北京市海淀区"
//	emp.Info.Phone = "123-4567-8901"
//
//	fmt.Printf("员工ID: %d, 地址: %s, 电话: %s\n", emp.ID, emp.Info.Address, emp.Info.Phone)
//
//	// 匿名结构体切片
//	people := []struct {
//		Name string
//		Age  int
//	}{
//		{"李四", 25},
//		{"王五", 35},
//		{"赵六", 40},
//	}
//
//	for _, p := range people {
//		fmt.Printf("姓名: %s, 年龄: %d\n", p.Name, p.Age)
//	}
//
//	fmt.Println(NUM)
//	fmt.Println(NUM1)
//	fmt.Println(Male)
//
//	gender := Male
//	fmt.Println(gender.String())
//	fmt.Println(gender.IsMale())
//
//	const (
//		f = 2
//		g = iota
//		h
//		i
//	)
//	fmt.Println(f, g, h, i)
//}
//
//func (p *Person) Print() {
//	fmt.Printf("姓名1: %s, 年龄1: %d\n", p.Name, p.Age)
//}
//
//func (g *Gender) String() string {
//	switch *g {
//	case Male:
//		return "Male"
//	case Female:
//		return "Female"
//	default:
//		return "Unknown"
//	}
//}
//
//func (g *Gender) IsMale() bool {
//	fmt.Println("IsMale:", *g)
//	return *g == Male
//}

type A struct {
	i int
}

func (a *A) add(v int) int {
	a.i += v
	return a.i
}

// 声明函数变量
var function1 func(int) int

// 声明闭包
var squart2 func(int) int = func(p int) int {
	p *= p
	return p
}

// 声明闭包
var str1 = func(str string) string {
	return str
}("hello world")

func main() {
	a := A{1}
	// 把方法赋值给函数变量
	function1 = a.add

	// 声明一个闭包并直接执行
	// 此闭包返回值是另外一个闭包（带参闭包）
	returnFunc := func() func(int, string) (int, string) {
		fmt.Println("this is a anonymous function")
		return func(i int, s string) (int, string) {
			return i, s
		}
	}()

	// 执行returnFunc闭包并传递参数
	ret1, ret2 := returnFunc(1, "test")
	fmt.Println("call closure function, return1 = ", ret1, "; return2 = ", ret2)

	fmt.Println("a.i = ", a.i)
	fmt.Println("after call function1, a.i = ", function1(1))
	fmt.Println("a.i = ", a.i)

	squart1 := squart2(2)
	fmt.Println("squart1 =", squart1)

	fmt.Println(str1)
}
