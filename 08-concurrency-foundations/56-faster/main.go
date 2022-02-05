package main

import (
	"sync"
)

func sequentialMergesort(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2
	sequentialMergesort(s[:middle])
	sequentialMergesort(s[middle:])
	merge(s, middle)
}

func parallelMergesortV1(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[:middle])
	}()

	go func() {
		defer wg.Done()
		parallelMergesortV1(s[middle:])
	}()

	wg.Wait()
	merge(s, middle)
}

const max = 2048

func parallelMergesortV2(s []int) {
	if len(s) <= 1 {
		return
	}

	if len(s) <= max {
		sequentialMergesort(s)
	} else {
		middle := len(s) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			parallelMergesortV2(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergesortV2(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

func merge(s []int, middle int) {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}
