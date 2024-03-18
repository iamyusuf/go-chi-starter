package main

import "log"

func main() {
	s := NewServer()
	log.Fatal(s.Start(":3000"))
}
