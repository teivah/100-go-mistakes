package main

func handleOperations1(id string) {
	operations := getOperations(id)
	if operations != nil {
		handle(operations)
	}
}

func handleOperations2(id string) {
	operations := getOperations(id)
	if len(operations) != 0 {
		handle(operations)
	}
}

func getOperations(id string) []float32 {
	operations := make([]float32, 0)

	if id == "" {
		return operations
	}

	// Add elements to operations

	return operations
}

func handle(operations []float32) {}
