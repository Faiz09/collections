go collections
--

Helper functions for multidimensional arrays

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
    }
    
  
  output
  
      [{Dinesh 28} {Dinesh 28}]
      31
      false
