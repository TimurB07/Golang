package main

import (
	"fmt"
	"math"
)
type Shape interface {
	Area() float64
}


type Circle struct {
	Radius float64
}


func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}


type Rectangle struct {
	Width, Height float64
}


func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}


func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}


	PrintArea(circle)
	PrintArea(rectangle)
}
/*type Shape interface {
	Area() float64
}


type Circle struct {
	Radius float64
}


func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}


type Rectangle struct {
	Width, Height float64
}


func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}


func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}


	PrintArea(circle)
	PrintArea(rectangle)
}
*/

/*type Employee struct {
	Name string
	ID   int
}

type Manager struct {
	Employee   
	Department string
}

func (e Employee) Work() {
	fmt.Println(e.Name, e.ID)
}

func main() {
	manager := Manager{
		Employee: Employee{
			Name: "Leo Messi",
			ID:   10,
		},
		Department: "Sport",
	}

	manager.Work()
}
*/

/*type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	person := Person{Name: "Timur", Age: 20}

	person.Greet()
}
*/


/*func add(a int, b int) int {
	return a + b
}

func swap(str1, str2 string) (string, string) {
	return str2, str1
}

func qr(div1, div2 int) (int, int) {
	q := div1 / div2
	r := div1 % div2
	return q, r
}

func main() {
	num1, num2 := 5, 3
	fmt.Println(add(num1, num2))

	var str1, str2 string
	fmt.Scanln(&str1, &str2)
	swappedStr1, swappedStr2 := swap(str1, str2)
	fmt.Println(swappedStr1, swappedStr2)

	div1, div2 := 10, 3
	q, r := qr(div1, div2)
	fmt.Println(div1, div2, q, r)
}*/

/*func main() {
	var num int
	fmt.Print()
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("positive")
	} else if num < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	var day int
	fmt.Scan(&day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}
}

*/

/*var n1 int = 5
var n2 float64 = 32.15
var n3 string = "Messi"
var n4 bool = true

func main() {
	n5 := 17
	fmt.Printf(" %T, %d\n", n1, n1)
	fmt.Printf(" %T, %.2f\n", n2, n2)
	fmt.Printf(" %T, %s\n", n3, n3)
	fmt.Printf(" %T, %t\n", n4, n4)
	fmt.Printf(" %T, %d\n", n5, n5)
}
*/

/* func main() {
	fmt.Println("Hello, World!")
}
*/
