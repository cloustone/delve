// TODO: disassembler support should be compiled in unconditionally,
// instead of being decided by the build-target architecture, and be
// part of the Arch object instead.

package proc

import (
	"github.com/go-delve/delve/pkg/proc/mips64asm"
)

func mips64AsmDecode(asmInst *AsmInstruction, mem []byte, regs Registers, memrw MemoryReadWriter, bi *BinaryInfo) error {
	asmInst.Size = 4
	asmInst.Bytes = mem[:asmInst.Size]

	inst, err := mips64asm.Decode(mem)
	if err != nil {
		asmInst.Inst = (*mips64ArchInst)(nil)
		return err
	}

	asmInst.Inst = (*mips64ArchInst)(&inst)
	asmInst.Kind = OtherInstruction

	switch inst.Op {
	case mips64asm.BL, mips64asm.BLR:
		asmInst.Kind = CallInstruction
	case mips64asm.RET, mips64asm.ERET:
		asmInst.Kind = RetInstruction
	case mips64asm.B, mips64asm.BR:
		asmInst.Kind = JmpInstruction
	}

	asmInst.DestLoc = resolveCallArgMips64(&inst, asmInst.Loc.PC, asmInst.AtPC, regs, memrw, bi)

	return nil
}

func resolveCallArgMips64(inst *mips64asm.Inst, instAddr uint64, currentGoroutine bool, regs Registers, mem MemoryReadWriter, bininfo *BinaryInfo) *Location {
	switch inst.Op {
	case mips64asm.BL, mips64asm.BLR, mips64asm.B, mips64asm.BR:
		//ok
	default:
		return nil
	}

	var pc uint64
	var err error

	switch arg := inst.Args[0].(type) {
	case mips64asm.Imm:
		pc = uint64(arg.Imm)
	case mips64asm.Reg:
		if !currentGoroutine || regs == nil {
			return nil
		}
		pc, err = regs.Get(int(arg))
		if err != nil {
			return nil
		}
	case mips64asm.PCRel:
		pc = uint64(instAddr) + uint64(arg)
	default:
		return nil
	}

	file, line, fn := bininfo.PCToLine(pc)
	if fn == nil {
		return &Location{PC: pc}
	}
	return &Location{PC: pc, File: file, Line: line, Fn: fn}
}

// Possible stacksplit prologues are inserted by stacksplit in
// $GOROOT/src/cmd/internal/obj/mips64/obj7.go.
var prologuesMips64 []opcodeSeq

func init() {
	var tinyStacksplit = opcodeSeq{uint64(mips64asm.MOV), uint64(mips64asm.CMP), uint64(mips64asm.B)}
	var smallStacksplit = opcodeSeq{uint64(mips64asm.SUB), uint64(mips64asm.CMP), uint64(mips64asm.B)}
	var bigStacksplit = opcodeSeq{uint64(mips64asm.CMP), uint64(mips64asm.B), uint64(mips64asm.ADD), uint64(mips64asm.SUB), uint64(mips64asm.MOV), uint64(mips64asm.CMP), uint64(mips64asm.B)}
	var unixGetG = opcodeSeq{uint64(mips64asm.LDR)}

	prologuesMips64 = make([]opcodeSeq, 0, 3)
	for _, getG := range []opcodeSeq{unixGetG} {
		for _, stacksplit := range []opcodeSeq{tinyStacksplit, smallStacksplit, bigStacksplit} {
			prologue := make(opcodeSeq, 0, len(getG)+len(stacksplit))
			prologue = append(prologue, getG...)
			prologue = append(prologue, stacksplit...)
			prologuesMips64 = append(prologuesMips64, prologue)
		}
	}
}

type mips64ArchInst mips64asm.Inst

func (inst *mips64ArchInst) Text(flavour AssemblyFlavour, pc uint64, symLookup func(uint64) (string, uint64)) string {
	if inst == nil {
		return "?"
	}

	var text string

	switch flavour {
	case GNUFlavour:
		text = mips64asm.GNUSyntax(mips64asm.Inst(*inst))
	default:
		text = mips64asm.GoSyntax(mips64asm.Inst(*inst), pc, symLookup, nil)
	}

	return text
}

func (inst *mips64ArchInst) OpcodeEquals(op uint64) bool {
	if inst == nil {
		return false
	}
	return uint64(inst.Op) == op
}
