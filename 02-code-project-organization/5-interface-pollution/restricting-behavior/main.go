package main

type IntConfig struct {
	value int
}

func (c *IntConfig) Get() int {
	return c.value
}

func (c *IntConfig) Set(value int) {
	c.value = value
}

type intConfigGetter interface {
	Get() int
}

type Foo struct {
	threshold intConfigGetter
}

func NewFoo(threshold intConfigGetter) Foo {
	return Foo{threshold: threshold}
}

func (f Foo) Bar() {
	threshold := f.threshold.Get()
	_ = threshold
}

func main() {
	foo := NewFoo(&IntConfig{value: 42})
	foo.Bar()
}
