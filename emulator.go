package main

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
	memory int
	cc     conditionCodes
	enable bool
}

func emulator(b []byte, pc int) int {
	var op int = 1
	return op
}
