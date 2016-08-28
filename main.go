package main

import (
	"io/ioutil"
	"log"
	"os"
)

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
		pc += emulator(b, pc)
	}
}
