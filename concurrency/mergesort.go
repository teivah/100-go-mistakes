package concurrency

import "sync"

func mergesortSequential(s []int) {
	if len(s) > 1 {
		middle := len(s) / 2
		mergesortSequential(s[:middle])
		mergesortSequential(s[middle:])
		merge(s, middle)
	}
}

func mergeSortConcurrentV1(s []int) {
	if len(s) > 1 {
		middle := len(s) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			mergeSortConcurrentV1(s[:middle])
		}()

		go func() {
			defer wg.Done()
			mergeSortConcurrentV1(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

const max = 1 << 20

func mergeSortConcurrentV2(s []int) {
	if len(s) > 1 {
		if len(s) <= max {
			// Sequential
			mergesortSequential(s)
		} else {
			// Concurrent
			middle := len(s) / 2

			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				defer wg.Done()
				mergeSortConcurrentV2(s[:middle])
			}()

			go func() {
				defer wg.Done()
				mergeSortConcurrentV2(s[middle:])
			}()

			wg.Wait()
			merge(s, middle)
		}
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
