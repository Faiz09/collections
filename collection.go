package collections

import "github.com/Faiz09/collections/methods"

type Collection struct {
	Data interface{}
}

func (c Collection) Where(name string, key interface{}) Collection {
	c.Data = methods.Where(c.Data, name, key)
	return c
}

func (c Collection) Avg(name string) int {
	 return methods.Avg(c.Data, name).(int)
}

func (c Collection) Exists(name string, key interface{}) bool {
	 return methods.Exists(c.Data, name, key)
}

func (c Collection) Count() int {
	 return methods.Count(c.Data)
}

func (c Collection) Each(f func(in interface{}) interface{}) interface{} {
	 return methods.Each(c.Data, f)
}

func (c Collection) get() interface{} {
	return c.Data
}
