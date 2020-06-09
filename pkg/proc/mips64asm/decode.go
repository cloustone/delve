package mips64asm

import (
	"encoding/binary"
	"fmt"
)

type instArgs [5]instArg

// An instFormat describes the format of an instruction encoding.
// An instruction with 32-bit value x matches the format if x&mask == value
// and the predicator: canDecode(x) return true.
type instFormat struct {
	mask  uint32
	value uint32
	op    Op
	// args describe how to decode the instruction arguments.
	// args is stored as a fixed-size array.
	// if there are fewer than len(args) arguments, args[i] == 0 marks
	// the end of the argument list.
	args      instArgs
	canDecode func(instr uint32) bool
}

var (
	errShort   = fmt.Errorf("truncated instruction")
	errUnknown = fmt.Errorf("unknown instruction")
)

var decoderCover []bool

func init() {
	decoderCover = make([]bool, len(instFormats))
}

// Decode decodes the 4 bytes in src as a single instruction.
func Decode(src []byte) (inst Inst, err error) {
	if len(src) < 4 {
		return Inst{}, errShort
	}

	x := binary.LittleEndian.Uint32(src)

Search:
	for i := range instFormats {
		f := &instFormats[i]
		if x&f.mask != f.value {
			continue
		}
		if f.canDecode != nil && !f.canDecode(x) {
			continue
		}
		// Decode args.
		var args Args
		for j, aop := range f.args {
			if aop == 0 {
				break
			}
			arg := decodeArg(aop, x)
			if arg == nil { // Cannot decode argument
				continue Search
			}
			args[j] = arg
		}
		decoderCover[i] = true
		inst = Inst{
			Op:   f.op,
			Args: args,
			Enc:  x,
		}
		return inst, nil
	}
	return Inst{}, errUnknown
}

// decodeArg decodes the arg described by aop from the instruction bits x.
// It returns nil if x cannot be decoded according to aop.
func decodeArg(aop instArg, x uint32) Arg {
	switch aop {
	default:
		return nil
	}

}
