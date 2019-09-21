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
            {name: "Erlic", age:  28},
            {name: "Big Head", age:  25},
            {name: "Big Head", age:  10},
        }
    
        c := Collection{p}
    
        fmt.Println(c.get())
    
        w := c.Where("name", "Dinesh").Where("age", 28).get().([]person)
    
        fmt.Println(w)
    
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
    
        fmt.Println(c.First().(person).name)
        fmt.Println(c.Last().(person).age)
    
        res := c.GroupBy("name").get().(map[string][]interface{})
    
        m := res["Big Head"]
    
        for j:=0;j<=len(f) ;j++  {
            fmt.Println(m[j].(person).name)
        }
    
        fmt.Println(c.Sort("name").get())
    }
    
  
  output
    
    [{Dinesh 28} {Gilf 32} {Richard 20} {Erlic 28} {Big Head 25} {Big Head 10}]
    [{Dinesh 28}]
    23
    false
    6
    1
    [{Dinesh **  28} {Gilf **  32} {Richard **  20} {Erlic **  28} {Big Head **  25} {Big Head **  10}]
    [{Gilf 32}]
    Dinesh
    10
    Big Head
    Big Head
    [{Big Head 25} {Big Head 10} {Dinesh 28} {Erlic 28} {Gilf 32} {Richard 20}]



