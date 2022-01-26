package main

import "strings"

func concat1(values []string) string {
	s := ""
	for _, value := range values {
		s += value
	}
	return s
}

func concat2(values []string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}

func concat3(values []string) string {
	total := 0
	for i := 0; i < len(values); i++ {
		total += len(values[i])
	}

	sb := strings.Builder{}
	sb.Grow(total)
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}
