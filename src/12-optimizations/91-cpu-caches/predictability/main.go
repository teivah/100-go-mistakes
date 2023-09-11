package main

type node struct {
	value int64
	next  *node
}

func linkedList(n *node) int64 {
	var total int64
	for n != nil {
		total += n.value
		n = n.next
	}
	return total
}

func sum2(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 2 {
		total += s[i]
	}
	return total
}
