package main

func main() {
	newWatcher()

	// Run the application
}

func newWatcher() {
	w := watcher{}
	go w.watch()
}

type watcher struct { /* Some resources */
}

func (w watcher) watch() {}
