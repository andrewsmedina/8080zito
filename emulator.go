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
	case 0x00: // NOP
	case 0x01: // LXI B,word
		s.c = int(b[s.pc+1])
		s.b = int(b[s.pc+2])
		op = 3
	case 0x05: // DCR B
		s.b = s.b - 1
		s.cc.z = s.b == 0
	case 0x06: // MVI B,byte
		s.b = int(b[s.pc+1])
		op = 2
	case 0x0d: // DCR C
		s.c = s.c - 1
		s.cc.z = s.b == 0
	case 0x0e: // MVI C,byte
		s.c = int(b[s.pc+1])
		op = 2
	case 0x11: // LXI D,word
		s.d = int(b[s.pc+1])
		s.e = int(b[s.pc+2])
		op = 3
	case 0x21: // LXI H,word
		s.h = int(b[s.pc+1])
		s.l = int(b[s.pc+2])
		op = 3
	case 0x26: // MVI H,byte
		s.h = int(b[s.pc+1])
		op = 2
	case 0x3e: // MVI A,byte
		s.a = int(b[s.pc+1])
		op = 2
	default:
		fmt.Printf("MISSING  %x\n", code)
	}
	s.pc += op
	return op
}
