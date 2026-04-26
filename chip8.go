package main

import "os"

type Chip8 struct {
	registers  [16]byte
	memory     [4096]byte
	index      uint16
	pc         uint16
	stack      [16]uint16
	sp         uint16
	delayTimer byte
	soundTimer byte
	keypad     [16]byte
	video      [64 * 32]uint32
	opcode     uint16
}

func NewChip8() *Chip8 {
	return &Chip8{
		pc: 0x200,
	}
}

func (chip *Chip8) loadROM(filename string) {
	index := 0x200
	data, err := os.ReadFile(filename); 
	if err != nil {
		panic(err)
	}

	copy(chip.memory[index:], data)
}