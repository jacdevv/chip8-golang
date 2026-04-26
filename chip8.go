package main

type Chip8 struct {
	registers  [16]uint8
	memory     [4096]uint8
	index      uint16
	pc         uint16
	stack      [16]uint16
	sp         uint16
	delayTimer uint8
	soundTimer uint8
	keypad     [16]uint8
	video      [64 * 32]uint32
	opcode     uint16
}