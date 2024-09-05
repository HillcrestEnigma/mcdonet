package main

func main() {
	s := NewServer()
	s.Listen("localhost:2000")
}