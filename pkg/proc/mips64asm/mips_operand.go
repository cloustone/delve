package mips64asm

type mispOperandType uint16

/* Enumerates the various types of MIPS operand.  */
const (
	/* Described by mips_int_operand.  */
	OP_INT mipsOperandType = itoa,

		/* Described by mips_mapped_int_operand.  */
		OP_MAPPED_INT,

		/* Described by mips_msb_operand.  */
		OP_MSB,

		/* Described by mips_reg_operand.  */
		OP_REG,

		/* Like OP_REG, but can be omitted if the register is the same as the
		   previous operand.  */
		OP_OPTIONAL_REG,

		/* Described by mips_reg_pair_operand.  */
		OP_REG_PAIR,

		/* Described by mips_pcrel_operand.  */
		OP_PCREL,

		/* A performance register.  The field is 5 bits in size, but the supported
		   values are much more restricted.  */
		OP_PERF_REG,

		/* The final operand in a microMIPS ADDIUSP instruction.  It mostly acts
		   as a normal 9-bit signed offset that is multiplied by four, but there
		   are four special cases:

		   -2 * 4 => -258 * 4
		   -1 * 4 => -257 * 4
			0 * 4 =>  256 * 4
			1 * 4 =>  257 * 4.  */
		OP_ADDIUSP_INT,

		/* The target of a (D)CLO or (D)CLZ instruction.  The operand spans two
		   5-bit register fields, both of which must be set to the destination
		   register.  */
		OP_CLO_CLZ_DEST,

		/* A register list for a microMIPS LWM or SWM instruction.  The operand
		   size determines whether the 16-bit or 32-bit encoding is required.  */
		OP_LWM_SWM_LIST,

		/* The register list for an emulated MIPS16 ENTRY or EXIT instruction.  */
		OP_ENTRY_EXIT_LIST,

		/* The register list and frame size for a MIPS16 SAVE or RESTORE
		   instruction.  */
		OP_SAVE_RESTORE_LIST,

		/* A 10-bit field VVVVVNNNNN used for octobyte and quadhalf instructions:

		   V      Meaning
		   -----  -------
		   0EEE0  8 copies of $vN[E], OB format
		   0EE01  4 copies of $vN[E], QH format
		   10110  all 8 elements of $vN, OB format
		   10101  all 4 elements of $vN, QH format
		   11110  8 copies of immediate N, OB format
		   11101  4 copies of immediate N, QH format.  */
		OP_MDMX_IMM_REG,

		/* A register operand that must match the destination register.  */
		OP_REPEAT_DEST_REG,

		/* A register operand that must match the previous register.  */
		OP_REPEAT_PREV_REG,

		/* $pc, which has no encoding in the architectural instruction.  */
		OP_PC,

		/* $28, which has no encoding in the MIPS16e architectural instruction.  */
		OP_REG28,

		/* A 4-bit XYZW channel mask or 2-bit XYZW index; the size determines
		   which.  */
		OP_VU0_SUFFIX,

		/* Like OP_VU0_SUFFIX, but used when the operand's value has already
		   been set.  Any suffix used here must match the previous value.  */
		OP_VU0_MATCH_SUFFIX,

		/* An index selected by an integer, e.g. [1].  */
		OP_IMM_INDEX,

		/* An index selected by a register, e.g. [$2].  */
		OP_REG_INDEX,

		/* The operand spans two 5-bit register fields, both of which must be set to
		   the source register.  */
		OP_SAME_RS_RT,

		/* Described by mips_prev_operand.  */
		OP_CHECK_PREV,

		/* A register operand that must not be zero.  */
		OP_NON_ZERO_REG
)

type mipsRegOperandType uint16

/* Enumerates the types of MIPS register.  */
const (
	/* General registers $0-$31.  Software names like $at can also be used.  */
	OP_REG_GP mipsRegOperandType = itoa,

		/* Floating-point registers $f0-$f31.  */
		OP_REG_FP,

		/* Coprocessor condition code registers $cc0-$cc7.  FPU condition codes
		   can also be written $fcc0-$fcc7.  */
		OP_REG_CCC,

		/* FPRs used in a vector capacity.  They can be written $f0-$f31
		   or $v0-$v31, although the latter form is not used for the VR5400
		   vector instructions.  */
		OP_REG_VEC,

		/* DSP accumulator registers $ac0-$ac3.  */
		OP_REG_ACC,

		/* Coprocessor registers $0-$31.  Mnemonic names like c0_cause can
		   also be used in some contexts.  */
		OP_REG_COPRO,

		/* Hardware registers $0-$31.  Mnemonic names like hwr_cpunum can
		   also be used in some contexts.  */
		OP_REG_HW,

		/* Floating-point registers $vf0-$vf31.  */
		OP_REG_VF,

		/* Integer registers $vi0-$vi31.  */
		OP_REG_VI,

		/* R5900 VU0 registers $I, $Q, $R and $ACC.  */
		OP_REG_R5900_I,
		OP_REG_R5900_Q,
		OP_REG_R5900_R,
		OP_REG_R5900_ACC,

		/* MSA registers $w0-$w31.  */
		OP_REG_MSA,

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
	bias     int
	shift    uint32
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
	mipsRegOperandType regType /* The type of register.  */
	regMap             []byte
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
func mipsInsertOperand(operand *mipsOperand, insn uint32, uval uint) uint32 {
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
	return ((uval + sign_bit) & mask) - sign_bit
}

// Return the integer that OPERAND encodes as UVAL
func mipsDecodeIntOperand(operand *mipsIntOperand, uval uint) {
	uval |= (operand.maxVal - uval) & -(1 << operand.size)
	uval += operand.bias
	uval <<= operand.shift
	return uval
}

// Return the maximum value that can be encoded by OPERAND.
func mipsIntOperandMax(operand *mipsIntOperand) int {
	return (operand.maxVal + operand.bias) << operand.shift
}

// Return the minimum value that can be encoded by OPERAND.
func mipsIntOperandMin(operand *mipsIntOperand) int {
	var mask uint
	mask = (1 << operand.size) - 1
	return mipsIntOperandMax(operand) - (mask << operand.shift)
}

// Return the register that OPERAND encodes as UVAL
func mipsDecodeRegOperand(operand *mipsRegOperand, uval uint) int {
	if operand.regMap {
		uval = operand.regRap[uval]
	}
	return uval
}

// PC-relative operand OPERAND has value UVAL and is relative to BASE_PC Return the address that it encodes
func mipsDecodePcrelOperand(operand *mipsPcrelOperand, base_pc uint32, uval uint32) uint32 {
	var addr uint32

	addr = base_pc & -(1 << operand.alignLog2)
	addr += mipsDecodeIntOperand(operand, uval)
	if operand.includeIsa {
		addr |= base_pc & 1
	}
	if operand.flipIsa {
		addr ^= 1
	}
	return addr
}
