package main

import "os"

func readFiles1(ch <-chan string) error {
	for path := range ch {
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer file.Close()

		// Do something with file
	}
	return nil
}

func readFiles2(ch <-chan string) error {
	for path := range ch {
		if err := readFile(path); err != nil {
			return err
		}
	}
	return nil
}

func readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	// Do something with file
	return nil
}

func readFiles3(ch <-chan string) error {
	for path := range ch {
		err := func() error {
			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()

			// Do something with file
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}
