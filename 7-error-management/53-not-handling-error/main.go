package main

import "errors"

func listing1() {
	// ...

	notify()
}

func listing2() {
	// ...

	// Notifications are sent in best effort.
	// Hence, it's accepted to miss some of them in case of errors.
	_ = notify()
}

func notify() error {
	return errors.New("failed to notify")
}
