package main

type Foo1 struct {
	b1 byte
	i  int64
	b2 byte
}

func sum1(foos []Foo1) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}

type Foo2 struct {
	i  int64
	b1 byte
	b2 byte
}

func sum2(foos []Foo2) int64 {
	var s int64
	for i := 0; i < len(foos); i++ {
		s += foos[i].i
	}
	return s
}
