go collections *Experimental
--

Helper functions for multidimensional arrays - inspired by laravel collections

    type person struct {
        name string
        age int64
    }
    
    func main()  {
       p := []person{
            {name: "Dinesh", age:  28},
            {name: "Gilf", age:  32},
            {name: "Richard", age:  20},
            {name: "Erlic", age:  58},
            {name: "Bar", age:  25},
            {name: "Dinesh", age:  28},
        }
       
        c := Collection{p}
       
        res := c.Where("name", "Dinesh").Where("age", 28).get().([]person)
       
        fmt.Println(res)
       
        fmt.Println(c.Avg("age"))
        fmt.Println(c.Exists("age", 1000))
        fmt.Println(c.Count())
        fmt.Println(c.Where("name", "Dinesh").Count())
       
       
        e := c.Each(func(p interface{}) interface{} {
            person := p.(person)
            person.name = person.name + " ** "
            return person
        }).([]person)
    
        fmt.Println(e)
    
        f := c.Filter(func(p interface{}) interface{} {
            person := p.(person)
            return person.age > 30
        }).([]person)
    
        fmt.Println(f)
    }
    
  
  output
    
    [{Dinesh 28} {Dinesh 28}]
    31
    false
    6
    2
    [{Dinesh **  28} {Gilf **  32} {Richard **  20} {Erlic **  58} {Bar **  25} {Dinesh **  28}]
    [{Gilf 32} {Erlic 58}]

