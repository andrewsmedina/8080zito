package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("invaders")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var pc int
	var s state
	for pc < len(b) {
		pc += emulator(&s, b)
	}
}
