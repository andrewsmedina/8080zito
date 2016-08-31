package main

import "fmt"

type conditionCodes struct {
	z   bool
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
	case 0x05:
		s.b = s.b - 1
		s.cc.z = s.b == 0
	case 0x06:
		s.b = int(b[s.pc+1])
		op = 2
	case 0x0d:
		s.c = s.c - 1
		s.cc.z = s.b == 0
	case 0x0e:
		s.c = int(b[s.pc+1])
		op = 2
	case 0x15:
		s.d = s.d - 1
		s.cc.z = s.b == 0
	case 0x1d:
		s.e = s.e - 1
		s.cc.z = s.b == 0
	case 0x25:
		s.h = s.h - 1
		s.cc.z = s.b == 0
	case 0x26:
		s.h = int(b[s.pc+1])
		op = 2
	case 0x2d:
		s.l = s.l - 1
		s.cc.z = s.b == 0
	case 0x3d:
		s.a = s.a - 1
		s.cc.z = s.b == 0
	case 0x3e:
		s.a = int(b[s.pc+1])
		op = 2
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
	case 0x58:
		s.e = s.b
	case 0x59:
		s.e = s.c
	case 0x5a:
		s.e = s.d
	case 0x5c:
		s.e = s.h
	case 0x5d:
		s.e = s.l
	case 0x5f:
		s.e = s.a
	default:
		fmt.Printf("MISSING  %x\n", code)
	}
	s.pc += op
	return op
}
