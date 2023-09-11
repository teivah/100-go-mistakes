package stringset

type Set map[string]struct{}

func New(...string) Set { return nil }

func (s Set) Sort() []string { return nil }
