package main

import "fmt"

type person struct {
	name string
	age int
}

func test()  {
	var p person
	p.name = "xx"
	p.age = 28

	fmt.Printf("The person's name is %s \n", p.name)

	pp := person{"xiaoming", 10}

	fmt.Printf("The person's name is %s \n", pp.name)

	ppp := person{name:"tom", age:14}

	fmt.Printf("The person's name is %s \n", ppp.name)

	P := new(person)

	P.name = "xx"
	P.age = 28

	fmt.Printf("The person's name is %s \n", P.name)

	older, age := Older(p, pp)
	fmt.Printf("The person's name is %s age is %d\n", older.name, age)
}

func Older(p1, p2 person) (olderman person, age int)  {

	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main()  {

	test()
}
