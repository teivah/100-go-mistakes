package main

func main() {
	var i any

	i = 42
	i = "foo"
	i = struct {
		s string
	}{
		s: "bar",
	}
	i = f

	_ = i
}

func f() {}
