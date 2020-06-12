package mips64asm

type mipsOperandType uint16

/* Enumerates the various types of MIPS operand.  */
const (
	/* Described by mips_int_operand.  */
	OP_INT mipsOperandType = iota

	/* Described by mips_mapped_int_operand.  */
	OP_MAPPED_INT

	/* Described by mips_msb_operand.  */
	OP_MSB

	/* Described by mips_reg_operand.  */
	OP_REG

	/* Like OP_REG, but can be omitted if the register is the same as the
	   previous operand.  */
	OP_OPTIONAL_REG

	/* Described by mips_reg_pair_operand.  */
	OP_REG_PAIR

	/* Described by mips_pcrel_operand.  */
	OP_PCREL

	/* A performance register.  The field is 5 bits in size, but the supported
	   values are much more restricted.  */
	OP_PERF_REG

	/* The final operand in a microMIPS ADDIUSP instruction.  It mostly acts
	   as a normal 9-bit signed offset that is multiplied by four, but there
	   are four special cases:

	   -2 * 4 => -258 * 4
	   -1 * 4 => -257 * 4
		0 * 4 =>  256 * 4
		1 * 4 =>  257 * 4.  */
	OP_ADDIUSP_INT

	/* The target of a (D)CLO or (D)CLZ instruction.  The operand spans two
	   5-bit register fields, both of which must be set to the destination
	   register.  */
	OP_CLO_CLZ_DEST

	/* A register list for a microMIPS LWM or SWM instruction.  The operand
	   size determines whether the 16-bit or 32-bit encoding is required.  */
	OP_LWM_SWM_LIST

	/* The register list for an emulated MIPS16 ENTRY or EXIT instruction.  */
	OP_ENTRY_EXIT_LIST

	/* The register list and frame size for a MIPS16 SAVE or RESTORE
	   instruction.  */
	OP_SAVE_RESTORE_LIST

	/* A 10-bit field VVVVVNNNNN used for octobyte and quadhalf instructions:

	   V      Meaning
	   -----  -------
	   0EEE0  8 copies of $vN[E], OB format
	   0EE01  4 copies of $vN[E], QH format
	   10110  all 8 elements of $vN, OB format
	   10101  all 4 elements of $vN, QH format
	   11110  8 copies of immediate N, OB format
	   11101  4 copies of immediate N, QH format.  */
	OP_MDMX_IMM_REG

	/* A register operand that must match the destination register.  */
	OP_REPEAT_DEST_REG

	/* A register operand that must match the previous register.  */
	OP_REPEAT_PREV_REG

	/* $pc, which has no encoding in the architectural instruction.  */
	OP_PC

	/* $28, which has no encoding in the MIPS16e architectural instruction.  */
	OP_REG28

	/* A 4-bit XYZW channel mask or 2-bit XYZW index; the size determines
	   which.  */
	OP_VU0_SUFFIX

	/* Like OP_VU0_SUFFIX, but used when the operand's value has already
	   been set.  Any suffix used here must match the previous value.  */
	OP_VU0_MATCH_SUFFIX

	/* An index selected by an integer, e.g. [1].  */
	OP_IMM_INDEX

	/* An index selected by a register, e.g. [$2].  */
	OP_REG_INDEX

	/* The operand spans two 5-bit register fields, both of which must be set to
	   the source register.  */
	OP_SAME_RS_RT

	/* Described by mips_prev_operand.  */
	OP_CHECK_PREV

	/* A register operand that must not be zero.  */
	OP_NON_ZERO_REG
)

type mipsRegOperandType uint16

/* Enumerates the types of MIPS register.  */
const (
	/* General registers $0-$31.  Software names like $at can also be used.  */
	OP_REG_GP mipsRegOperandType = iota

	/* Floating-point registers $f0-$f31.  */
	OP_REG_FP

	/* Coprocessor condition code registers $cc0-$cc7.  FPU condition codes
	   can also be written $fcc0-$fcc7.  */
	OP_REG_CCC

	/* FPRs used in a vector capacity.  They can be written $f0-$f31
	   or $v0-$v31, although the latter form is not used for the VR5400
	   vector instructions.  */
	OP_REG_VEC

	/* DSP accumulator registers $ac0-$ac3.  */
	OP_REG_ACC

	/* Coprocessor registers $0-$31.  Mnemonic names like c0_cause can
	   also be used in some contexts.  */
	OP_REG_COPRO

	/* Hardware registers $0-$31.  Mnemonic names like hwr_cpunum can
	   also be used in some contexts.  */
	OP_REG_HW

	/* Floating-point registers $vf0-$vf31.  */
	OP_REG_VF

	/* Integer registers $vi0-$vi31.  */
	OP_REG_VI

	/* R5900 VU0 registers $I, $Q, $R and $ACC.  */
	OP_REG_R5900_I
	OP_REG_R5900_Q
	OP_REG_R5900_R
	OP_REG_R5900_ACC

	/* MSA registers $w0-$w31.  */
	OP_REG_MSA

	/* MSA control registers $0-$31.  */
	OP_REG_MSA_CTRL
)

/* Base class for all operands.  */
type mipsOperand struct {
	operandType mipsOperandType // The type of the operand
	size        uint16          // The operand occupies SIZE bits of the instruction, starting at LSB
	lsb         uint16
}

/* Describes an integer operand with a regular encoding pattern.  */
type mipsIntOperand struct {
	mipsOperand
	/* The low ROOT.SIZE bits of MAX_VAL encodes (MAX_VAL + BIAS) << SHIFT.
	   The cyclically previous field value encodes 1 << SHIFT less than that,
	   and so on.  E.g.

	   - for { { T, 4, L }, 14, 0, 0 }, field values 0...14 encode themselves,
	     but 15 encodes -1.

	   - { { T, 8, L }, 127, 0, 2 } is a normal signed 8-bit operand that is
	     shifted left two places.

	   - { { T, 3, L }, 8, 0, 0 } is a normal unsigned 3-bit operand except
	     that 0 encodes 8.

	   - { { ... }, 0, 1, 3 } means that N encodes (N + 1) << 3.  */
	maxVal   uint16
	bias     uint16
	shift    uint16
	printHex bool // True if the operand should be printed as hex rather than decimal
}

/* Uses a lookup table to describe a small integer operand.  */
type mipsMappedIntOperand struct {
	mipsOperand
	intMap    int  /* Maps each encoding value to the integer that it represents.  */
	printHhex bool /* True if the operand should be printed as hex rather than decimal.  */
}

/* An operand that encodes the most significant bit position of a bitfield.
   Given a bitfield that spans bits [MSB, LSB], some operands of this type
   encode MSB directly while others encode MSB - LSB.  Each operand of this
   type is preceded by an integer operand that specifies LSB.

   The assembly form varies between instructions.  For some instructions,
   such as EXT, the operand is written as the bitfield size.  For others,
   such as EXTS, it is written in raw MSB - LSB form.  */
type mipsMsbOperand struct {
	mipsOperand
	bias   int    /* The assembly-level operand encoded by a field value of 0.  */
	addLsb bool   /* True if the operand encodes MSB directly, false if it encodes MSB - LSB.  */
	opsize uint32 /* The maximum value of MSB + 1.  */
}

/* Describes a single register operand.  */
type mipsRegOperand struct {
	mipsOperand
	regType mipsRegOperandType /* The type of register.  */
	regMap  []byte
	/* If nonnull, REG_MAP[N] gives the register associated with encoding N,
	   otherwise the encoding is the same as the register number.  */
}

/* Describes an operand that which must match a condition based on the
   previous operand.  */
type mipsCheckPrevOperand struct {
	mipsOperand
	greaterThan bool
	lessThan    bool
	equal       bool
	zero        bool
}

/* Describes an operand that encodes a pair of registers.  */
type mipsRegPairOperand struct {
	mipsOperand
	regType mipsRegOperandType /* The type of register.  */
	reg1Map []byte             /* Encoding N represents REG1_MAP[N], REG2_MAP[N].  */
	reg2Map []byte
}

/* Describes an operand that is calculated relative to a base PC.
   The base PC is usually the address of the following instruction,
   but the rules for MIPS16 instructions like ADDIUPC are more complicated.  */
type mipsPcrelOperand struct {
	mipsOperand
	/* The low ALIGN_LOG2 bits of the base PC are cleared to give PC',
	which is then added to the offset encoded by ROOT.  */
	alignLog2 uint16

	/* If INCLUDE_ISA_BIT, the ISA bit of the original base PC is then
	reinstated.  This is true for jumps and branches and false for
	PC-relative data instructions.  */
	includeIsa bool

	/* If FLIP_ISA_BIT, the ISA bit of the result is inverted.
	This is true for JALX and false otherwise.  */
	flipIsa bool
}

/* Return true if the assembly syntax allows OPERAND to be omitted.  */
func mipsOptionalOperand(operand *mipsOperand) bool {
	return (operand.operandType == OP_OPTIONAL_REG || operand.operandType == OP_REPEAT_PREV_REG)
}

// Return a version of INSN in which the field specified by OPERAND has value UVAL.
func mipsInsertOperand(operand *mipsOperand, insn uint32, uval uint32) uint32 {
	var mask uint32
	mask = (1 << operand.size) - 1
	// insn &= ~(mask << operand.lsb) // TODO
	insn |= (uval & mask) << operand.lsb
	return insn
}

// Extract OPERAND from instruction INSN
func mipsExtractOperand(operand *mipsOperand, insn uint) uint {
	return (insn >> operand.lsb) & ((1 << operand.size) - 1)
}

// UVAL is the value encoded by OPERAND.  Return it in signed form.
func mipsSignedOperand(operand *mipsOperand, uval uint) int {
	var sign_bit, mask uint
	mask = (1 << operand.size) - 1
	sign_bit = 1 << (operand.size - 1)
	return int(((uval + sign_bit) & mask) - sign_bit)
}

// Return the integer that OPERAND encodes as UVAL
func mipsDecodeIntOperand(operand *mipsIntOperand, uval uint16) uint16 {
	uval |= (operand.maxVal - uval) & -(1 << operand.size)
	uval += uint16(operand.bias)
	uval <<= operand.shift
	return uval
}

// Return the maximum value that can be encoded by OPERAND.
func mipsIntOperandMax(operand *mipsIntOperand) int {
	return int((operand.maxVal + operand.bias) << operand.shift)
}

// Return the minimum value that can be encoded by OPERAND.
func mipsIntOperandMin(operand *mipsIntOperand) int {
	var mask uint
	mask = (1 << operand.size) - 1
	return mipsIntOperandMax(operand) - int(mask<<operand.shift)
}

// Return the register that OPERAND encodes as UVAL
func mipsDecodeRegOperand(operand *mipsRegOperand, uval uint) uint {
	if operand.regMap != nil {
		uval = uint(operand.regMap[uval])
	}
	return uval
}

// PC-relative operand OPERAND has value UVAL and is relative to BASE_PC Return the address that it encodes
func mipsDecodePcrelOperand(operand *mipsPcrelOperand, base_pc uint32, uval uint32) uint32 {
	var addr uint32
	/*

		addr = base_pc & -(1 << operand.alignLog2)
		addr += uint32(mipsDecodeIntOperand(operand, uval))
		if operand.includeIsa {
			addr |= base_pc & 1
		}
		if operand.flipIsa {
			addr ^= 1
		}
	*/ // TODO
	return addr
}

func xIntBias(size, lsb, maxVal, bial, shift uint16, printHex bool) *mipsOperand {
	return &mipsIntOperand{
		mipsOperand{OP_INT, size, lsb}, maxVal, bial, shift, printHex,
	}
}

func xIntAdj(size, lsb, maxVal, shift int, printHex bool) mipsOperand {
	return xIntBias(size, lsb, maxVal, 0, shift, printHex)
}

func xUint(size, lsb int) mipsOperand {
	return xIntAdj(size, lsb, (1<<(size))-1, 0, false)
}

func xSint(size, lsb int) mipsOperand {
	return xIntAdj(size, lsb, (1<<((size)-1))-1, 0, false)
}

func xHint(size, lsb int) mipsOperand {
	return xIntAdj(size, lsb, (1<<(size))-1, 0, true)
}

func xBit(size, lsb, bias int) mipsOperand {
	return mipsIntOperand{
		{OP_INT, size, lsb}, (1 << (size)) - 1, bias, 0, true,
	}
}

func xMappedInt(size, lsb, maps int, printHex bool) mipsOperand {
	return mipsMappedIntOperand{
		{OP_MAPPED_INT, size, lsb}, maps, printHex,
	}
}

func xMsb(size, lsb, bias, addLsb, opSize int) mipsOperand {
	return mipsMsbOperand{
		{OP_MSB, size, lsb}, bias, addLsb, opSize,
	}
}

func xReg(size, lsb, bank int) mipsOperand {
	switch bank {
	case GP:
		op = OP_REG_GP
	default:
	}
	return mipsRegOperand{
		{OP_REG, size, lsb}, op, 0,
	}
}

func xOptionalReg2(size, lsb, bank int) mipsOperand {
	switch bank {
	case GP:
		op = OP_REG_GP
	default:
	}
	return mipsRegOperand{
		{OP_OPTIONAL_REG, size, lsb}, op, 0,
	}
}

func xMappedReg(size, lsb, bank, maps int) mipsOperand {
	switch bank {
	case GP:
		op = OP_REG_GP
	default:
	}
	return mipsRegOperand{
		{OP_REG, size, lsb}, op, maps,
	}
}

func xOptionalReg(size, lsb, bank, maps int) mipsOperand {
	switch bank {
	case GP:
		op = OP_REG_GP
	default:
	}
	return mipsRegOperand{
		{OP_OPTIONAL_REG, size, lsb}, op, maps,
	}
}

func xRegPair(szie, lsb, bank, maps int) mipsOperand {
	switch bank {
	case GP:
		op = OP_REG_GP
	default:
	}
	return mipsRegPairOperand{
		{OP_REG_PAIR, size, lsb}, op, MAP1, MAP2,
	}
}

func xPcrel(size, lsb, isSigned, shift, alignLog2 int, includeIsaBit, flipIsaBit bool) mipsOperand {
	return mipsPcrelOperand{
		{{OP_PCREL, size, lsb},
			(1 << ((size) - (isSigned))) - 1, 0, shift, true},
		alignLog2, includeIsaBit, flipIsaBit,
	}
}

func xJump(size, lsb, shift int) mipsOperand {
	return xPcrel(size, lsb, false, shift, size+shift, true, false)
}

func xJalx(size, lsb, shift int) mipsOperand {
	return xPcrel(size, lsb, false, shift, size+shift, true, true)
}

func xBranch(size, lsb, shift int) mipsOperand {
	return xPcrel(size, lsb, true, shift, 0, true, false)
}

func xSpecial(size, lsb, tp int) mipsOperand {
	var op int
	switch tp {
	case OP_TYPE1:
		op = OP_TYPE
	default:
	}

	return mipsOperand{op, size, lsb}
}

func xPrevCheck(size, lsb int, gtOK, ltOK, eqOK, zeroOK bool) mipsOperand {
	return mipsCheckPrevOperand{
		{OP_CHECK_PREV, size, lsb}, gtOK, ltOK, eqOK, zeroOK,
	}
}

var (
	// The 4-bit XYZW mask used in some VU0 instructions.
	mips_vu0_channel_mask = mipsOperand{OP_VU0_SUFFIX, 4, 21}
	reg_0_map             = []uint8{0}
)

// Return the mips_operand structure for the operand at the beginning of P
func decodeMipsOperand(p []byte) *mipsOperand {
	switch p[0] {
	case '-':
		switch p[1] {
		case 'a':
			return xIntAdj(19, 0, 262143, 2, false)
		case 'b':
			return xIntAdj(18, 0, 131071, 3, false)
		case 'd':
			return xSpecial(0, 0, REPEAT_DEST_REG)
		case 'm':
			return xSpecial(20, 6, SAVE_RESTORE_LIST)
		case 's':
			return xSpecial(5, 21, NON_ZERO_REG)
		case 't':
			return xSpecial(5, 16, NON_ZERO_REG)
		case 'u':
			return xPrevCheck(5, 16, true, false, false, false)
		case 'v':
			return xPrevCheck(5, 16, true, true, false, false)
		case 'w':
			return xPrevCheck(5, 16, false, true, true, true)
		case 'x':
			return xPrevCheck(5, 21, true, false, false, true)
		case 'y':
			return xPrevCheck(5, 21, false, true, false, false)
		case 'A':
			return xPcrel(19, 0, true, 2, 2, false, false)
		case 'B':
			return xPcrel(18, 0, true, 3, 3, false, false)
		}
		break

	case '+':
		switch p[1] {
		case '1':
			return xHint(5, 6)
		case '2':
			return xHint(10, 6)
		case '3':
			return xHint(15, 6)
		case '4':
			return xHint(20, 6)
		case '5':
			return xReg(5, 6, VF)
		case '6':
			return xReg(5, 11, VF)
		case '7':
			return xReg(5, 16, VF)
		case '8':
			return xReg(5, 6, VI)
		case '9':
			return xReg(5, 11, VI)
		case '0':
			return xReg(5, 16, VI)

		case 'A':
			return xBit(5, 6, 0) /* (0 .. 31) */
		case 'B':
			return xMsb(5, 11, 1, true, 32) /* (1 .. 32), 32-bit op */
		case 'C':
			return xMsb(5, 11, 1, false, 32) /* (1 .. 32), 32-bit op */
		case 'E':
			return xBit(5, 6, 32) /* (32 .. 63) */
		case 'F':
			return xMsb(5, 11, 33, true, 64) /* (33 .. 64), 64-bit op */
		case 'G':
			return xMsb(5, 11, 33, false, 64) /* (33 .. 64), 64-bit op */
		case 'H':
			return xMsb(5, 11, 1, false, 64) /* (1 .. 32), 64-bit op */
		case 'I':
			return xUint(2, 6)
		case 'J':
			return xHint(10, 11)
		case 'K':
			return xSpecial(4, 21, VU0_MATCH_SUFFIX)
		case 'L':
			return xSpecial(2, 21, VU0_SUFFIX)
		case 'M':
			return xSpecial(2, 23, VU0_SUFFIX)
		case 'N':
			return xSpecial(2, 0, VU0_MATCH_SUFFIX)
		case 'O':
			return xUint(3, 6)
		case 'P':
			return xBit(5, 6, 32) /* (32 .. 63) */
		case 'Q':
			return xSint(10, 6)
		case 'R':
			return xSpecial(0, 0, PC)
		case 'S':
			return xMsb(5, 11, 0, false, 63) /* (0 .. 31), 64-bit op */
		case 'T':
			return xIntAdj(10, 16, 511, 0, false) /* (-512 .. 511) << 0 */
		case 'U':
			return xIntAdj(10, 16, 511, 1, false) /* (-512 .. 511) << 1 */
		case 'V':
			return xIntAdj(10, 16, 511, 2, false) /* (-512 .. 511) << 2 */
		case 'W':
			return xIntAdj(10, 16, 511, 3, false) /* (-512 .. 511) << 3 */
		case 'X':
			return xBit(5, 16, 32) /* (32 .. 63) */
		case 'Z':
			return xReg(5, 0, FP)

		case 'a':
			return xSint(8, 6)
		case 'b':
			return xSint(8, 3)
		case 'c':
			return xIntAdj(9, 6, 255, 4, false) /* (-256 .. 255) << 4 */
		case 'd':
			return xReg(5, 6, MSA)
		case 'e':
			return xReg(5, 11, MSA)
		case 'f':
			return xIntAdj(15, 6, 32767, 3, true)
		case 'g':
			return xSint(5, 6)
		case 'h':
			return xReg(5, 16, MSA)
		case 'i':
			return xJalx(26, 0, 2)
		case 'j':
			return xSint(9, 7)
		case 'k':
			return xReg(5, 6, GP)
		case 'l':
			return xReg(5, 6, MSA_CTRL)
		case 'm':
			return xReg(0, 0, R5900_ACC)
		case 'n':
			return xReg(5, 11, MSA_CTRL)
		case 'o':
			return xSpecial(4, 16, IMM_INDEX)
		case 'p':
			return xBit(5, 6, 0) /* (0 .. 31), 32-bit op */
		case 'q':
			return xReg(0, 0, R5900_Q)
		case 'r':
			return xReg(0, 0, R5900_R)
		case 's':
			return xMsb(5, 11, 0, false, 31) /* (0 .. 31) */
		case 't':
			return xReg(5, 16, COPRO)
		case 'u':
			return xSpecial(3, 16, IMM_INDEX)
		case 'v':
			return xSpecial(2, 16, IMM_INDEX)
		case 'w':
			return xSpecial(1, 16, IMM_INDEX)
		case 'x':
			return xBit(5, 16, 0) /* (0 .. 31) */
		case 'y':
			return xReg(0, 0, R5900_I)
		case 'z':
			return xReg(5, 0, GP)

		case '~':
			return xBit(2, 6, 1) /* (1 .. 4) */
		case '!':
			return xBit(3, 16, 0) /* (0 .. 7) */
		case '@':
			return xBit(4, 16, 0) /* (0 .. 15) */
		case '#':
			return xBit(6, 16, 0) /* (0 .. 63) */
		case '$':
			return xUint(5, 16) /* (0 .. 31) */
		case '%':
			return xSint(5, 16) /* (-16 .. 15) */
		case '^':
			return xSint(10, 11) /* (-512 .. 511) */
		case '&':
			return xSpecial(0, 0, IMM_INDEX)
		case '*':
			return xSpecial(5, 16, REG_INDEX)
		case '|':
			return xBit(8, 16, 0) /* (0 .. 255) */
		case ':':
			return xSint(11, 0)
		case '\'':
			return xBranch(26, 0, 2)
		case '"':
			return xBranch(21, 0, 2)
		case ' ':
			return xSpecial(10, 16, SAME_RS_RT) // TODO
		case '\\':
			return xBit(2, 8, 0) /* (0 .. 3) */
		}
		break

	case '<':
		return xBit(5, 6, 0) /* (0 .. 31) */
	case '>':
		return xBit(5, 6, 32) /* (32 .. 63) */
	case '%':
		return xUint(3, 21)
	case ':':
		return xSint(7, 19)
	case '\'':
		return xHint(6, 16)
	case '@':
		return xSint(10, 16)
	case '!':
		return xUint(1, 5)
	case '$':
		return xUint(1, 4)
	case '*':
		return xReg(2, 18, ACC)
	case '&':
		return xReg(2, 13, ACC)
	case '~':
		return xSint(12, 0)
	case '\\':
		return xBit(3, 12, 0) /* (0 .. 7) */

	case '0':
		return xSint(6, 20)
	case '1':
		return xHint(5, 6)
	case '2':
		return xHint(2, 11)
	case '3':
		return xHint(3, 21)
	case '4':
		return xHint(4, 21)
	case '5':
		return xHint(8, 16)
	case '6':
		return xHint(5, 21)
	case '7':
		return xReg(2, 11, ACC)
	case '8':
		return xHint(6, 11)
	case '9':
		return xReg(2, 21, ACC)

	case 'B':
		return xHint(20, 6)
	case 'C':
		return xHint(25, 0)
	case 'D':
		return xReg(5, 6, FP)
	case 'E':
		return xReg(5, 16, COPRO)
	case 'G':
		return xReg(5, 11, COPRO)
	case 'H':
		return xUint(3, 0)
	case 'J':
		return xHint(19, 6)
	case 'K':
		return xReg(5, 11, HW)
	case 'M':
		return xReg(3, 8, CCC)
	case 'N':
		return xReg(3, 18, CCC)
	case 'O':
		return xUint(3, 21)
	case 'P':
		return xSpecial(5, 1, PERF_REG)
	case 'Q':
		return xSpecial(10, 16, MDMX_IMM_REG)
	case 'R':
		return xReg(5, 21, FP)
	case 'S':
		return xReg(5, 11, FP)
	case 'T':
		return xReg(5, 16, FP)
	case 'U':
		return xSpecial(10, 11, CLO_CLZ_DEST)
	case 'V':
		return xOptionalReg(5, 11, FP)
	case 'W':
		return xOptionalReg(5, 16, FP)
	case 'X':
		return xReg(5, 6, VEC)
	case 'Y':
		return xReg(5, 11, VEC)
	case 'Z':
		return xReg(5, 16, VEC)

	case 'a':
		return xJump(26, 0, 2)
	case 'b':
		return xReg(5, 21, GP)
	case 'c':
		return xHint(10, 16)
	case 'd':
		return xReg(5, 11, GP)
	case 'e':
		return xUint(3, 22)
	case 'g':
		return xReg(5, 11, COPRO)
	case 'h':
		return xHint(5, 11)
	case 'i':
		return xHint(16, 0)
	case 'j':
		return xSint(16, 0)
	case 'k':
		return xHint(5, 16)
	case 'o':
		return xSint(16, 0)
	case 'p':
		return xBranch(16, 0, 2)
	case 'q':
		return xHint(10, 6)
	case 'r':
		return xOptionalReg(5, 21, GP)
	case 's':
		return xReg(5, 21, GP)
	case 't':
		return xReg(5, 16, GP)
	case 'u':
		return xHint(16, 0)
	case 'v':
		return xOptionalReg(5, 21, GP)
	case 'w':
		return xOptionalReg(5, 16, GP)
	case 'x':
		return xReg(0, 0, GP)
	case 'z':
		return xMappedReg(0, 0, GP, reg_0_map)
	}
	return nil
}
