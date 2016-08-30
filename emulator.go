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
	case 0x48:
		s.c = s.b
	case 0x4a:
		s.c = s.d
	case 0x4b:
		s.c = s.e
	case 0x4c:
		s.c = s.h
	case 0x4d:
		s.c = s.l
	case 0x4f:
		s.c = s.a
	case 0x50:
		s.d = s.b
	case 0x51:
		s.d = s.c
	case 0x53:
		s.d = s.e
	case 0x54:
		s.d = s.h
	case 0x55:
		s.d = s.l
	case 0x57:
		s.d = s.a
	default:
		fmt.Printf("MISSING  %x\n", code)
	}
	s.pc += op
	return op
}
