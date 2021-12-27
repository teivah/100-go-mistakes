package main

type Config struct {
	Port int
}

func NewServer(addr string, cfg Config) {
}

func main() {
	NewServer("localhost", Config{})
}
