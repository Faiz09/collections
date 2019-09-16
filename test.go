package collections

import "fmt"

type person struct {
	name string
	age int64
}

func Test()  {

	p := []person{
		{name: "Chris", age:  20},
		{name: "Marrie", age:  28},
		{name: "David", age:  20},
		{name: "Dr", age:  100},
		{name: "Brew", age:  2},
		{name: "Chris", age:  30},
		{name: "Bar", age:  130},
	}

	c := Collection{p}

	res := c.where("name", "Chris").where("age", 20).get().([]person)

	fmt.Printf("%+v\n", res)
	fmt.Printf("%+T\n", res)
}