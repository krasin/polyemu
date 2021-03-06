package dcpu16

import (
	"github.com/krasin/polyemu/emu"
	"testing"
)

var tests = []emu.Test{
	{
		Mem:     []byte{0x61, 0x88}, // SET X, 1
		WantReg: []uint64{RX: 1, PC: 1},
		N:       1,
	},
	{
		Mem:     []byte{0x61, 0x8c, 0x64, 0x0c}, // SET X, 2 ; MUL X, X
		WantReg: []uint64{RX: 4, PC: 2},
		N:       2,
	},
	{
		Mem:     []byte{0x81, 0xfc}, // SET Y, 30
		WantReg: []uint64{RY: 30, PC: 1},
		N:       1,
	},
	{
		Mem: []byte{
			0x81, 0xfc, // SET Y, 30
			0x81, 0x99, // SET [Y], 5
			0x61, 0x78, 0x1e, 0x00, // SET X, [30]
		},
		WantReg: []uint64{RX: 5, RY: 30, PC: 4},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x98, // SET A, 5
			0x01, 0xa2, 0x01, 0x00, // SET [A+1], 7
			0x21, 0x78, 0x06, 0x00, // SET B, [6]
		},
		WantReg: []uint64{RA: 5, RB: 7, PC: 5},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0xaf, // SET PUSH, 10
			0x61, 0x78, 0xFF, 0xFF, // SET X, [0xFFFF]
		},
		WantReg: []uint64{RX: 10, SP: 0xFFFF, PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x88, // SET X, 1
			0x65, 0x80, // MLI X, -1
		},
		WantReg: []uint64{RX: 0xFFFF, EX: 0xFFFF, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0xa8, // SET X, 9
			0x66, 0x90, // DIV X, 3
		},
		WantReg: []uint64{RX: 3, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0xa8, // SET X, 9
			0xa1, 0x8b, // SET EX, 1
			0x66, 0x84, // DIV X, 0
		},
		WantReg: []uint64{PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x61, 0xa8, // SET X, 9
			0x67, 0x7c, 0xfd, 0xff, // DVI X, -3
		},
		WantReg: []uint64{RX: 0xFFFD, PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0xa8, // SET X, 9
			0x67, 0x7c, 0xfe, 0xff, // DVI X, -2
		},
		WantReg: []uint64{RX: 0xFFFC, PC: 3, EX: 0x8000},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0xa8, // SET X, 9
			0x67, 0x7c, 0xfc, 0xff, // DVI X, -2
		},
		WantReg: []uint64{RX: 0xFFFE, PC: 3, EX: 0xC000},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0xa8, // SET X, 9
			0xa1, 0x8b, // SET EX, 1
			0x67, 0x84, // DVI X, 0
		},
		WantReg: []uint64{PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0xa1, 0x8b, // SET EX, 1
		},
		WantReg: []uint64{EX: 1, PC: 1},
		N:       1,
	},
	{
		Mem: []byte{
			0xa1, 0x8b, // SET EX, 1
			0xe2, 0x8b, 0x01, 0x00, // ADD 1, 1
		},
		WantReg: []uint64{PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0xe2, 0x8b, 0xFF, 0xFF, // ADD 0xFFFF, 1
		},
		WantReg: []uint64{PC: 2, EX: 1},
		N:       1,
	},
	{
		Mem: []byte{
			0x61, 0x90, // SET X, 3
			0x62, 0x94, // ADD X, 4
		},
		WantReg: []uint64{RX: 7, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x60, 0xea, // SET X, 60000
			0x62, 0x7c, 0x10, 0x27, // ADD X, 10000
		},
		WantReg: []uint64{RX: 4464, PC: 4, EX: 1},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x98, // SET X, 5
			0x63, 0x90, // SUB X, 3
		},
		WantReg: []uint64{RX: 2, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x98, // SET X, 5
			0x63, 0xa0, // SUB X, 7
		},
		WantReg: []uint64{RX: 0xFFFE, PC: 2, EX: 0xFFFF},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x98, // SET X, 5
			0xa1, 0x8b, // SET EX, 1
			0x63, 0x98, // SUB X, 5
		},
		WantReg: []uint64{PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x61, 0x98, // SET X, 5
			0xa1, 0x8b, // SET EX, 1
			0x64, 0x0c, // MUL X, X
		},
		WantReg: []uint64{RX: 25, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x61, 0x98, // SET X, 5
			0x64, 0x7c, 0x20, 0x4e, // MUL X, 20000
		},
		WantReg: []uint64{RX: 34464, PC: 3, EX: 1},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x20, 0x4e, // SET X, 20000
			0x64, 0x0c, // MUL X, X
		},
		WantReg: []uint64{RX: 33792, PC: 3, EX: 6103},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x20, 0x4e, // SET X, 20000
			0x65, 0x0c, // MLI X, X
		},
		WantReg: []uint64{RX: 33792, PC: 3, EX: 6103},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x98, // SET X, 5
			0x65, 0x80, // MLI X, -1
		},
		WantReg: []uint64{RX: 0xFFFB, PC: 2, EX: 0xFFFF},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x98, // SET X, 5
			0xa1, 0x8b, // SET EX, 1
			0x65, 0x0c, // MLI X, X
		},
		WantReg: []uint64{RX: 25, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x65, 0x0c, // MLI X, X
		},
		WantReg: []uint64{RX: 1604, PC: 3, EX: 11},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x68, 0xb8, // MOD X, 13
		},
		WantReg: []uint64{RX: 5, PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x80, // SET X, 65535
			0x68, 0x98, // MOD X, 5
		},
		WantReg: []uint64{PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x68, 0x84, // MOD X, 0
		},
		WantReg: []uint64{PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x69, 0xb8, // MDI X, 13
		},
		WantReg: []uint64{RX: 5, PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0xf9, 0xff, // SET X, -7
			0x69, 0xc4, // MDI X, 16
		},
		WantReg: []uint64{RX: 0xFFF9, PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x80, // SET X, 65535
			0x69, 0x98, // MDI X, 5
		},
		WantReg: []uint64{RX: 0xFFFF, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x69, 0x84, // MDI X, 0
		},
		WantReg: []uint64{PC: 3},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x6a, 0x7c, 0xd1, 0x06, // AND X, 1745
		},
		WantReg: []uint64{RX: 592, PC: 4},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x6b, 0x7c, 0xd1, 0x06, // BOR X, 1745
		},
		WantReg: []uint64{RX: 2003, PC: 4},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x6c, 0x7c, 0xd1, 0x06, // XOR X, 1745
		},
		WantReg: []uint64{RX: 1411, PC: 4},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x6d, 0x7c, 0xd1, 0x06, // SHR X, 1745
		},
		WantReg: []uint64{PC: 4},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x6d, 0x7c, 0x05, 0x00, // SHR X, 5
		},
		WantReg: []uint64{RX: 26, PC: 4, EX: 36864},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x52, 0x03, // SET X, 850
			0x6d, 0x7c, 0xff, 0xff, // SHR X, 65535
		},
		WantReg: []uint64{PC: 4},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0xff, 0xff, // SET X, 65535
			0x6d, 0x7c, 0x05, 0x00, // SHR X, 5
		},
		WantReg: []uint64{RX: 2047, PC: 4, EX: 63488},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0xff, 0xff, // SET X, -1
			0x6e, 0x7c, 0x05, 0x00, // ASR X, 5
		},
		WantReg: []uint64{RX: 65535, PC: 4, EX: 63488},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x00, 0x70, // SET X, 28672
			0x6e, 0x7c, 0x05, 0x00, // ASR X, 5
		},
		WantReg: []uint64{RX: 896, PC: 4},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0xff, 0xff, // SET X, 65535
			0x6f, 0x7c, 0x05, 0x00, // SHL X, 5
		},
		WantReg: []uint64{RX: 65504, PC: 4, EX: 31},
		N:       2,
	},
	{
		Mem: []byte{
			0x61, 0x7c, 0x20, 0x4e, // SET X, 20000
			0x6f, 0x7c, 0x05, 0x00, // SHL X, 5
		},
		WantReg: []uint64{RX: 50176, PC: 4, EX: 9},
		N:       2,
	},
	{
		Mem: []byte{
			0xa1, 0x83, // SET EX, 0xFFFF
			0x61, 0x80, // SET X, 0xFFFF
			0x7a, 0x80, // ADX X, 0xFFFF
		},
		WantReg: []uint64{RX: 0xFFFD, PC: 3, EX: 1},
		N:       3,
	},
	{
		Mem: []byte{
			0xa1, 0x8b, // SET EX, 1
			0x41, 0x8c, // SET C, 2
			0x3a, 0x08, // ADX B, C
		},
		WantReg: []uint64{RB: 3, RC: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x41, 0x8c, // SET C, 2
			0x3a, 0x08, // ADX B, C
		},
		WantReg: []uint64{RB: 2, RC: 2, PC: 2},
		N:       2,
	},
	// SBX
	{
		Mem: []byte{
			0xa1, 0x83, // SET EX, 0xFFFF
			0x61, 0x80, // SET X, 0xFFFF
			0x7b, 0x80, // SBX X, 0xFFFF
		},
		WantReg: []uint64{RX: 0xFFFF, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0xa1, 0x8b, // SET EX, 1
			0x41, 0x8c, // SET C, 2
			0x3b, 0x08, // SBX B, C
		},
		WantReg: []uint64{RB: 0xFFFF, RC: 2, PC: 3, EX: 0xFFFF},
		N:       3,
	},
	{
		Mem: []byte{
			0x41, 0x8c, // SET C, 2
			0x3b, 0x08, // SBX B, C
		},
		WantReg: []uint64{RB: 0xFFFE, RC: 2, PC: 2, EX: 0xFFFF},
		N:       2,
	},
	// STI
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x3e, 0x00, // STI B, A
		},
		WantReg: []uint64{RA: 1, RB: 1, RI: 1, RJ: 1, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0xc1, 0x88, // SET I, 1
			0xfe, 0x18, // STI J, I
		},
		WantReg: []uint64{RI: 2, RJ: 2, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0xc1, 0x7c, 0x64, 0x00, // SET I, 100
			0xe1, 0x7c, 0xc8, 0x00, // SET J, 200
			0xe1, 0xc1, // SET [J], 15
			0xde, 0x3e, 0x01, 0x00, // STI [I+1], [J]
			0x01, 0x38, // SET A, [I]
		},
		WantReg: []uint64{RA: 15, RI: 101, RJ: 201, PC: 8},
		N:       5,
	},
	// STD
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x3f, 0x00, // STD B, A
		},
		WantReg: []uint64{RA: 1, RB: 1, RI: 0xFFFF, RJ: 0xFFFF, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0xc1, 0x88, // SET I, 1
			0xff, 0x18, // STD J, I
		},
		WantReg: []uint64{RI: 0, RJ: 0, PC: 2},
		N:       2,
	},
	{
		Mem: []byte{
			0xc1, 0x7c, 0x64, 0x00, // SET I, 100
			0xe1, 0x7c, 0xc8, 0x00, // SET J, 200
			0xe1, 0xc1, // SET [J], 15
			0xdf, 0x3e, 0xFF, 0xFF, // STD [I-1], [J]
			0x01, 0x38, // SET A, [I]
		},
		WantReg: []uint64{RA: 15, RI: 99, RJ: 199, PC: 8},
		N:       5,
	},
	// IFB
	{
		Mem: []byte{
			0x10, 0x04, // IFB A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RY: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x10, 0x04, // IFB A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x21, 0x88, // SET B, 1
			0x10, 0x04, // IFB A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RB: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x80, // SET A, 0xFFFF
			0x10, 0x04, // IFB A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 0xFFFF, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x80, // SET A, 0xFFFF
			0x21, 0x88, // SET B, 1
			0x10, 0x04, // IFB A, B
			0x61, 0x88, // SET X, 1
		},
		WantReg: []uint64{RA: 0xFFFF, RB: 1, RX: 1, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x21, 0x88, // SET B, 1
			0x10, 0x04, // IFB A, B
			0x61, 0x88, // SET X, 1
		},
		WantReg: []uint64{RA: 1, RB: 1, RX: 1, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x98, // SET A, 5
			0x21, 0x9c, // SET B, 6
			0x10, 0x04, // IFB A, B
			0x61, 0x88, // SET X, 1
		},
		WantReg: []uint64{RA: 5, RB: 6, RX: 1, PC: 4},
		N:       4,
	},
	// IFC
	{
		Mem: []byte{
			0x11, 0x04, // IFC A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RX: 1, RY: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x11, 0x04, // IFC A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 1, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x21, 0x88, // SET B, 1
			0x11, 0x04, // IFC A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RB: 1, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x80, // SET A, 0xFFFF
			0x11, 0x04, // IFC A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 0xFFFF, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x80, // SET A, 0xFFFF
			0x21, 0x88, // SET B, 1
			0x11, 0x04, // IFC A, B
			0x61, 0x88, // SET X, 1
		},
		WantReg: []uint64{RA: 0xFFFF, RB: 1, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x21, 0x88, // SET B, 1
			0x11, 0x04, // IFC A, B
			0x61, 0x88, // SET X, 1
		},
		WantReg: []uint64{RA: 1, RB: 1, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x98, // SET A, 5
			0x21, 0x9c, // SET B, 6
			0x11, 0x04, // IFC A, B
			0x61, 0x88, // SET X, 1
		},
		WantReg: []uint64{RA: 5, RB: 6, PC: 4},
		N:       4,
	},

	// IFE
	{
		Mem: []byte{
			0x12, 0x04, // IFE A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RX: 1, RY: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x12, 0x04, // IFE A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 1, RY: 2, PC: 4},
		N:       4,
	},
	// IFN
	{
		Mem: []byte{
			0x13, 0x04, // IFN A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RY: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x13, 0x04, // IFN A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 1, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	// IFG
	{
		Mem: []byte{
			0x14, 0x04, // IFG A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RY: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x14, 0x04, // IFG A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 1, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x80, // SET A, 0xFFFF
			0x14, 0x04, // IFG A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 0xFFFF, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	// IFA
	{
		Mem: []byte{
			0x15, 0x04, // IFA A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RY: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x15, 0x04, // IFA A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 1, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x80, // SET A, 0xFFFF
			0x15, 0x04, // IFA A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 0xFFFF, RY: 2, PC: 4},
		N:       4,
	},
	// IFL
	{
		Mem: []byte{
			0x16, 0x04, // IFL A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RY: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x16, 0x04, // IFL A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x21, 0x88, // SET B, 1
			0x16, 0x04, // IFL A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RB: 1, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x80, // SET A, 0xFFFF
			0x16, 0x04, // IFL A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 0xFFFF, RY: 2, PC: 4},
		N:       4,
	},
	// IFU
	{
		Mem: []byte{
			0x17, 0x04, // IFU A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RY: 2, PC: 3},
		N:       3,
	},
	{
		Mem: []byte{
			0x01, 0x88, // SET A, 1
			0x17, 0x04, // IFU A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x21, 0x88, // SET B, 1
			0x17, 0x04, // IFU A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RB: 1, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	{
		Mem: []byte{
			0x01, 0x80, // SET A, 0xFFFF
			0x17, 0x04, // IFU A, B
			0x61, 0x88, // SET X, 1
			0x81, 0x8c, // SET Y, 2
		},
		WantReg: []uint64{RA: 0xFFFF, RX: 1, RY: 2, PC: 4},
		N:       4,
	},
	// PC
	{
		Mem: []byte{
			0x61, 0x70, // SET X, PC
			0x81, 0x70, // SET Y, PC
			0x81, 0x87, // SET PC, 0
		},
		WantReg: []uint64{RX: 1, RY: 2, PC: 0},
		N:       3,
	},
	// PC = 0xFFFF
	{
		Mem: []byte{
			131070: 0x61, 131071: 0x88, // SET X, 1
		},
		Reg:     []uint64{PC: 0xFFFF},
		WantReg: []uint64{RX: 1, PC: 0},
		N:       1,
	},
	// This test case is correct, but many other emulators fail on it.
	{
		Mem: []byte{
			0x61, 0x72, 0x10, 0x00, // SET [X+16], PC
			0x81, 0x78, 0x10, 0x00, // SET Y, [16]
		},
		WantReg: []uint64{RY: 2, PC: 4},
		N:       2,
	},
	// JSR
	{
		Mem: []byte{
			0x20, 0x8c, // JSR 2
			0x61, 0x88, // SET X, 1
			0x81, 0x64, // SET Y, [SP]
		},
		WantReg: []uint64{RY: 1, PC: 3, SP: 0xFFFF},
		N:       2,
	},
	// Return from function
	{
		Mem: []byte{
			0x20, 0x94, // JSR 4
			0x61, 0x88, // SET X, 1
			0x61, 0x84, // SET X, 0
			0x61, 0x84, // SET X, 0
			0x81, 0x88, // SET Y, 1
			0x81, 0x63, // SET PC, POP
		},
		WantReg: []uint64{RX: 1, RY: 1, PC: 2},
		N:       4,
	},
}

func TestSet(t *testing.T) {
	for testInd, tt := range tests {
		tt.Reg = append(tt.Reg, make([]uint64, RegCount)...)
		tt.Mem = append(tt.Mem, make([]byte, 2*65536)...)
		e := new(Emulator)
		emu.RunTest(testInd, t, e, tt)
	}
}
