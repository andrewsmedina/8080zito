package main

import "fmt"

type conditionCodes struct {
	z   int
	s   int
	p   int
	cy  int
	ac  int
	pad int
}

type state struct {
	a      int
	b      int
	c      int
	d      int
	e      int
	h      int
	l      int
	sp     int
	pc     int
	cc     conditionCodes
	enable bool
}

func emulator(s *state, b []byte) int {
	op := 1
	code := b[s.pc]
	switch code {
	default:
		fmt.Printf("MISSING  %x\n", code)
	}
	s.pc += op
	return op
}
