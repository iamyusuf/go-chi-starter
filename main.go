package main

func main() {
	s := NewServer()
	err := s.Start(":3000")

	if err != nil {
		return
	}
}
