package mips64asm

import (
	"encoding/binary"
	"fmt"
)

//type instArgs [5]instArg
type instArgs [5]byte
type Inst [4]byte

var (
	errShort   = fmt.Errorf("truncated instruction")
	errUnknown = fmt.Errorf("unknown instruction")
)

func init() {
}

// Decode decodes the 4 bytes in src as a single instruction.
func Decode(src []byte) (inst Inst, err error) {
	if len(src) < 4 {
		return Inst{}, errShort
	}

	x := binary.LittleEndian.Uint32(src)

	return Inst{}, errUnknown
}
