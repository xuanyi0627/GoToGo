package main

import (
	_ "fmt"
	"fmt"
)

func Flow_If(x int)  {

	fmt.Println("------------------------------- Flow_If")

	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	if m := computedValue(); m > 10 {
		fmt.Println("m is greater than 10")
	} else {
		fmt.Println("m is less than 10")
	}

	if x == 3 {
		fmt.Println("The integer is equal to 3")
	} else if x < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}
}

func computedValue() int {
	return 6
}

func Flow_Goto()  {

	fmt.Println("------------------------------- Flow_Goto")

	i := 0

Here:

	if i < 10 {

		fmt.Println(i)

		i++
	}

	goto Here
}

func Flow_Foo()  {

	fmt.Println("------------------------------- Flow_Foo")

	sum := 0

	for index := 0; index < 10 ; index++ {
		sum += index
	}

	fmt.Println("sum is equal to ", sum)

	sum = 1

	for ; sum < 1000 ;  {
		sum += sum
	}

	fmt.Println("sum is equal to ", sum)

	sum = 1

	for sum < 1000  {
		sum += sum
	}

	fmt.Println("sum is equal to ", sum)

	for index := 10; index > 0; index-- {
		if index == 5 {
			break
		}
		fmt.Println("break : ", index)
	}

	for index := 10; index > 0; index-- {
		if index == 5 {
			continue
		}
		fmt.Println("continue : ", index)
	}

	rating := map[string]float32{"C":5, "Go":4.5, "Python":4.5, "C++":2 }

	for k, v := range rating{
		fmt.Println("map's key:",k)
		fmt.Println("map's val:",v)
	}

	for _, v := range rating{
		fmt.Println("map's val:",v)
	}

}

func Flow_Switch()  {

	fmt.Println("------------------------------- Flow_Switch")

	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

func main()  {

	Flow_If(5)

	Flow_If(3)

	//Flow_Goto()

	Flow_Foo()

	Flow_Switch()
}