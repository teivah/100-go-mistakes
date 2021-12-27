package concepts

type Store struct {
	m map[string]*Foo
}

func (s Store) Put(id string, foo *Foo) {
	s.m[id] = foo
	// ...
}

type Foo struct{}

func updateMapValue(mapValue map[string]LargeStruct, id string) {
	value := mapValue[id]
	value.foo = "bar"
	mapValue[id] = value
}

func updateMapPointer(mapPointer map[string]*LargeStruct, id string) {
	mapPointer[id].foo = "bar"
}

type LargeStruct struct {
	foo string
	_   [1024]int64
}
