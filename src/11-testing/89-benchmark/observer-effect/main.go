package main

func calculateSum512(s [][512]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			sum += s[i][j]
		}
	}
	return sum
}

func calculateSum513(s [][513]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			sum += s[i][j]
		}
	}
	return sum
}
