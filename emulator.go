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
	case 0x01:
		s.c = int(b[s.pc+1])
		s.b = int(b[s.pc+2])
		op = 3
	case 0x41:
		s.b = s.c
	case 0x42:
		s.b = s.d
	case 0x43:
		s.b = s.e
	case 0x44:
		s.b = s.h
	case 0x45:
		s.b = s.l
	case 0x47:
		s.b = s.a
	default:
		fmt.Printf("MISSING  %x\n", code)
	}
	s.pc += op
	return op
}
