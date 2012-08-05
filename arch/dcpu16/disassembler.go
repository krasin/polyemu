package dcpu16

import (
	"github.com/krasin/polyemu/emu"

	"fmt"
)

func opcodeStr(opcode int) string {
	switch opcode {
	case SPECIAL_OP:
		return "SPECIAL_OP"
	case SET_OP:
		return "SET"
	case ADD_OP:
		return "ADD"
	case SUB_OP:
		return "SUB"
	case MUL_OP:
		return "MUL"
	case MLI_OP:
		return "MLI"
	case DIV_OP:
		return "DIV"
	case DVI_OP:
		return "DVI"
	case MOD_OP:
		return "MOD"
	case MDI_OP:
		return "MDI"
	case AND_OP:
		return "AND"
	case BOR_OP:
		return "BOR"
	case XOR_OP:
		return "XOR"
	case SHR_OP:
		return "SHR"
	case ASR_OP:
		return "ASR"
	case SHL_OP:
		return "SHL"
	case IFB_OP:
		return "IFB"
	case IFC_OP:
		return "IFC"
	case IFE_OP:
		return "IFE"
	case IFN_OP:
		return "IFN"
	case IFG_OP:
		return "IFG"
	case IFA_OP:
		return "IFA"
	case IFL_OP:
		return "IFL"
	case IFU_OP:
		return "IFU"
	case ADX_OP:
		return "ADX"
	case SBX_OP:
		return "SBX"
	case STI_OP:
		return "STI"
	case STD_OP:
		return "STD"
	}
	return fmt.Sprintf("OPCODE:%d", opcode)
}

func regStr(ind int) string {
	switch ind {
	case RA:
		return "A"
	case RB:
		return "B"
	case RC:
		return "C"
	case RX:
		return "X"
	case RY:
		return "Y"
	case RZ:
		return "Z"
	case RI:
		return "I"
	case RJ:
		return "J"
	case PC:
		return "PC"
	case SP:
		return "SP"
	case EX:
		return "EX"
	}
	return fmt.Sprintf("REG:%d", ind)
}

func argStr(ar arg) string {
	switch ar.mode {
	case REG_ARG:
		return regStr(int(ar.val))
	case REG_ADDR_ARG: // [register]
		return fmt.Sprintf("[%s]", regStr(int(ar.val)))
	case REG_ADDR_WORD_ARG: // [register + next word]
		return fmt.Sprintf("[%s + 0x%04x]", regStr(int(ar.val)), ar.val2)
	case PUSH_ARG: // (PUSH / [--SP])
		return fmt.Sprintf("PUSH")
	case POP_ARG: // POP / [SP++]
		return fmt.Sprintf("POP")
	case ADDR_WORD_ARG: // [next word]
		return fmt.Sprintf("[0x%04x]", ar.val)
	case WORD_ARG: // next word (literal)
		return fmt.Sprintf("0x%04x", ar.val)
	case LITERAL_ARG: // literal (-1..30)
		return fmt.Sprintf("0x%04x", ar.val)
	}
	return fmt.Sprintf("[mode=%d, val=%d, val2=%d]", ar.mode, ar.val, ar.val2)
}

func Disassemble(mem []byte) (string, emu.Code) {
	st16 := &state{
		mem: &memory{a: mem, diff: make(map[uint64]byte)},
		reg: &regState{a: make([]uint64, 30), diff: make(map[uint64]uint64)},
	}
	if _, code := st16.Step(); code != emu.OK {
		return "", code
	}
	if st16.opcode == SPECIAL_OP {
		return "SPECIAL_OP", emu.OK
	}
	return fmt.Sprintf("%v %v, %v", opcodeStr(st16.opcode), argStr(st16.argB), argStr(st16.argA)), emu.OK
}
