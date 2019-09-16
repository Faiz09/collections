package collections

import "github.com/Faiz09/collections/methods"

type Collection struct {
	data interface{}
}

func (c Collection) where(name string, key interface{}) Collection {
	c.data = methods.Where(c.data, name, key)
	return c
}

func (c Collection) get() interface{} {
	return c.data
}
