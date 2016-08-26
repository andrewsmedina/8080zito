package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func disassemble(b []byte, pc int) int {
	var op int = 1
	code := b[pc]
	switch code {
	case 0x00:
		fmt.Println("NOP")
	case 0x01:
		fmt.Printf("LXI    B,#$%02x%02x\n", b[pc+2], b[pc+1])
		op = 3
	case 0x02:
		fmt.Println("STAX   B")
	case 0x03:
		fmt.Println("INX    B")
	case 0x04:
		fmt.Println("INR    B")
	case 0x05:
		fmt.Println("DCR    B")
	case 0x06:
		fmt.Printf("MVI    B,#$%02x\n", b[pc+1])
		op = 2
	case 0x07:
		fmt.Println("RLC")
	case 0x08:
		fmt.Println("NOP")
	case 0x3e:
		fmt.Printf("MVI    A,#0x%02x\n", b[pc+1])
		op = 2
	case 0xc3:
		fmt.Printf("JMP    $%02x%02x\n", b[pc+2], b[pc+1])
		op = 3
	default:
		fmt.Printf("MISSING  %x\n", code)
	}
	return op
}

func main() {
	file, err := os.Open("invaders.h")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var pc int
	for pc < len(b) {
		pc += disassemble(b, pc)
	}
}
