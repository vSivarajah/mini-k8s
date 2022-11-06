package main

func main() {
	s := NewAPIServer(":3000")
	s.Start()
}
