package main

import "strings"

func concat1(ids []string) string {
	s := ""
	for _, id := range ids {
		s += id
	}
	return s
}

func concat2(ids []string) string {
	sb := strings.Builder{}
	for _, id := range ids {
		_, _ = sb.WriteString(id)
	}
	return sb.String()
}

func concat3(ids []string) string {
	total := 0
	for i := 0; i < len(ids); i++ {
		total += len(ids[i])
	}

	sb := strings.Builder{}
	sb.Grow(total)
	for _, id := range ids {
		_, _ = sb.WriteString(id)
	}
	return sb.String()
}
