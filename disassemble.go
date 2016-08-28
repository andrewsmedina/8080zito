package main

import "fmt"

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
	case 0x09:
		fmt.Println("DAD    B")
	case 0x0a:
		fmt.Println("LDAX   B")
	case 0x0b:
		fmt.Println("DCX    B")
	case 0x0c:
		fmt.Println("INR    C")
	case 0x0d:
		fmt.Println("DCR    C")
	case 0x0e:
		fmt.Printf("MVI    C,#0x%02x\n", b[pc+1])
		op = 1
	case 0x0f:
		fmt.Println("RRC")
	case 0x10:
		fmt.Println("-")
	case 0x11:
		fmt.Printf("LXI    D,#$%02x%02x\n", b[pc+2], b[pc+1])
		op = 3
	case 0x12:
		fmt.Println("STAX    D")
	case 0x13:
		fmt.Println("INX    D")
	case 0x14:
		fmt.Println("INR    D")
	case 0x15:
		fmt.Println("DCR    D")
	case 0x16:
		fmt.Printf("MVI    D,#0x%02x\n", b[pc+1])
		op = 2
	case 0x17:
		fmt.Println("RAL")
	case 0x18:
		fmt.Println("-")
	case 0x19:
		fmt.Println("DAD    D")
	case 0x1a:
		fmt.Println("LDAX    D")
	case 0x1b:
		fmt.Println("DCX    D")
	case 0x1c:
		fmt.Println("INR    E")
	case 0x1d:
		fmt.Println("DCR    E")
	case 0x1e:
		fmt.Printf("MVI E,#0x%02x\n", b[pc+1])
		op = 2
	case 0x1f:
		fmt.Println("RAR")
	case 0x3e:
		fmt.Printf("MVI    A,#0x%02x\n", b[pc+1])
		op = 2
	case 0x38:
		fmt.Println("-")
	case 0x78:
		fmt.Println("MOV    A,B")
	case 0x8d:
		fmt.Println("ADC    L")
	case 0x93:
		fmt.Println("SUB    E")
	case 0xc3:
		fmt.Printf("JMP    $%02x%02x\n", b[pc+2], b[pc+1])
		op = 3
	case 0xd6:
		fmt.Printf("SUI    #0x%02x\n", b[pc+1])
		op = 2
	case 0xff:
		fmt.Println("RST    7")
	default:
		fmt.Printf("MISSING  %x\n", code)
	}
	return op
}
