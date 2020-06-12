package mips64asm

/* These are bit masks and shift counts to use to access the various
   fields of an instruction.  To retrieve the X field of an
   instruction, use the expression
	(i >> OP_SH_X) & OP_MASK_X
   To set the same field (to j), use
	i = (i &~ (OP_MASK_X << OP_SH_X)) | (j << OP_SH_X)

   Make sure you use fields that are appropriate for the instruction,
   of course.

   The 'i' format uses OP, RS, RT and IMMEDIATE.

   The 'j' format uses OP and TARGET.

   The 'r' format uses OP, RS, RT, RD, SHAMT and FUNCT.

   The 'b' format uses OP, RS, RT and DELTA.

   The floating point 'i' format uses OP, RS, RT and IMMEDIATE.

   The floating point 'r' format uses OP, FMT, FT, FS, FD and FUNCT.

   A breakpoint instruction uses OP, CODE and SPEC (10 bits of the
   breakpoint instruction are not defined; Kane says the breakpoint
   code field in BREAK is 20 bits; yet MIPS assemblers and debuggers
   only use ten bits).  An optional two-operand form of break/sdbbp
   allows the lower ten bits to be set too, and MIPS32 and later
   architectures allow 20 bits to be set with a single operand for
   the sdbbp instruction (using CODE20).

   The syscall instruction uses CODE20.

   The general coprocessor instructions use COPZ.  */
const (
	OP_MASK_OP        = 0x3f
	OP_SH_OP          = 26
	OP_MASK_RS        = 0x1f
	OP_SH_RS          = 21
	OP_MASK_FR        = 0x1f
	OP_SH_FR          = 21
	OP_MASK_FMT       = 0x1f
	OP_SH_FMT         = 21
	OP_MASK_BCC       = 0x7
	OP_SH_BCC         = 18
	OP_MASK_CODE      = 0x3ff
	OP_SH_CODE        = 16
	OP_MASK_CODE2     = 0x3ff
	OP_SH_CODE2       = 6
	OP_MASK_RT        = 0x1f
	OP_SH_RT          = 16
	OP_MASK_FT        = 0x1f
	OP_SH_FT          = 16
	OP_MASK_CACHE     = 0x1f
	OP_SH_CACHE       = 16
	OP_MASK_RD        = 0x1f
	OP_SH_RD          = 11
	OP_MASK_FS        = 0x1f
	OP_SH_FS          = 11
	OP_MASK_PREFX     = 0x1f
	OP_SH_PREFX       = 11
	OP_MASK_CCC       = 0x7
	OP_SH_CCC         = 8
	OP_MASK_CODE20    = 0xfffff /* 20 bit syscall/breakpoint code.  */
	OP_SH_CODE20      = 6
	OP_MASK_SHAMT     = 0x1f
	OP_SH_SHAMT       = 6
	OP_MASK_EXTLSB    = OP_MASK_SHAMT
	OP_SH_EXTLSB      = OP_SH_SHAMT
	OP_MASK_STYPE     = OP_MASK_SHAMT
	OP_SH_STYPE       = OP_SH_SHAMT
	OP_MASK_FD        = 0x1f
	OP_SH_FD          = 6
	OP_MASK_TARGET    = 0x3ffffff
	OP_SH_TARGET      = 0
	OP_MASK_COPZ      = 0x1ffffff
	OP_SH_COPZ        = 0
	OP_MASK_IMMEDIATE = 0xffff
	OP_SH_IMMEDIATE   = 0
	OP_MASK_DELTA     = 0xffff
	OP_SH_DELTA       = 0
	OP_MASK_FUNCT     = 0x3f
	OP_SH_FUNCT       = 0
	OP_MASK_SPEC      = 0x3f
	OP_SH_SPEC        = 0
	OP_SH_LOCC        = 8  /* FP condition code.  */
	OP_SH_HICC        = 18 /* FP condition code.  */
	OP_MASK_CC        = 0x7
	OP_SH_COP1NORM    = 25  /* Normal COP1 encoding.  */
	OP_MASK_COP1NORM  = 0x1 /* a single bit.  */
	OP_SH_COP1SPEC    = 21  /* COP1 encodings.  */
	OP_MASK_COP1SPEC  = 0xf
	OP_MASK_COP1SCLR  = 0x4
	OP_MASK_COP1CMP   = 0x3
	OP_SH_COP1CMP     = 4
	OP_SH_FORMAT      = 21 /* FP short format field.  */
	OP_MASK_FORMAT    = 0x7
	OP_SH_TRUE        = 16
	OP_MASK_TRUE      = 0x1
	OP_SH_GE          = 17
	OP_MASK_GE        = 0x01
	OP_SH_UNSIGNED    = 16
	OP_MASK_UNSIGNED  = 0x1
	OP_SH_HINT        = 16
	OP_MASK_HINT      = 0x1f
	OP_SH_MMI         = 0 /* Multimedia (parallel) op.  */
	OP_MASK_MMI       = 0x3f
	OP_SH_MMISUB      = 6
	OP_MASK_MMISUB    = 0x1f
	OP_MASK_PERFREG   = 0x1f /* Performance monitoring.  */
	OP_SH_PERFREG     = 1
	OP_SH_SEL         = 0   /* Coprocessor select field.  */
	OP_MASK_SEL       = 0x7 /* The sel field of mfcZ and mtcZ.  */
	OP_SH_CODE19      = 6   /* 19 bit wait code.  */
	OP_MASK_CODE19    = 0x7ffff
	OP_SH_ALN         = 21
	OP_MASK_ALN       = 0x7
	OP_SH_VSEL        = 21
	OP_MASK_VSEL      = 0x1f
	OP_MASK_VECBYTE   = 0x7 /* Selector field is really 4 bits,
	   but 0x8-0xf don't select bytes.  */
	OP_SH_VECBYTE    = 22
	OP_MASK_VECALIGN = 0x7 /* Vector byte-align (alni.ob) op.  */
	OP_SH_VECALIGN   = 21
	OP_MASK_INSMSB   = 0x1f /* "ins" MSB.  */
	OP_SH_INSMSB     = 11
	OP_MASK_EXTMSBD  = 0x1f /* "ext" MSBD.  */
	OP_SH_EXTMSBD    = 11

	/* MIPS DSP ASE */
	OP_SH_DSPACC     = 11
	OP_MASK_DSPACC   = 0x3
	OP_SH_DSPACC_S   = 21
	OP_MASK_DSPACC_S = 0x3
	OP_SH_DSPSFT     = 20
	OP_MASK_DSPSFT   = 0x3f
	OP_SH_DSPSFT_7   = 19
	OP_MASK_DSPSFT_7 = 0x7f
	OP_SH_SA3        = 21
	OP_MASK_SA3      = 0x7
	OP_SH_SA4        = 21
	OP_MASK_SA4      = 0xf
	OP_SH_IMM8       = 16
	OP_MASK_IMM8     = 0xff
	OP_SH_IMM10      = 16
	OP_MASK_IMM10    = 0x3ff
	OP_SH_WRDSP      = 11
	OP_MASK_WRDSP    = 0x3f
	OP_SH_RDDSP      = 16
	OP_MASK_RDDSP    = 0x3f
	OP_SH_BP         = 11
	OP_MASK_BP       = 0x3

	/* MIPS MT ASE */
	OP_SH_MT_U      = 5
	OP_MASK_MT_U    = 0x1
	OP_SH_MT_H      = 4
	OP_MASK_MT_H    = 0x1
	OP_SH_MTACC_T   = 18
	OP_MASK_MTACC_T = 0x3
	OP_SH_MTACC_D   = 13
	OP_MASK_MTACC_D = 0x3

	/* MIPS MCU ASE */
	OP_MASK_3BITPOS  = 0x7
	OP_SH_3BITPOS    = 12
	OP_MASK_OFFSET12 = 0xfff
	OP_SH_OFFSET12   = 0

	OP_OP_COP0 = 0x10
	OP_OP_COP1 = 0x11
	OP_OP_COP2 = 0x12
	OP_OP_COP3 = 0x13
	OP_OP_LWC1 = 0x31
	OP_OP_LWC2 = 0x32
	OP_OP_LWC3 = 0x33 /* a.k.a. pref */
	OP_OP_LDC1 = 0x35
	OP_OP_LDC2 = 0x36
	OP_OP_LDC3 = 0x37 /* a.k.a. ld */
	OP_OP_SWC1 = 0x39
	OP_OP_SWC2 = 0x3a
	OP_OP_SWC3 = 0x3b
	OP_OP_SDC1 = 0x3d
	OP_OP_SDC2 = 0x3e
	OP_OP_SDC3 = 0x3f /* a.k.a. sd */

	/* MIPS VIRT ASE */
	OP_MASK_CODE10 = 0x3ff
	OP_SH_CODE10   = 11

	/* Values in the 'VSEL' field.  */
	MDMX_FMTSEL_IMM_QH = 0x1d
	MDMX_FMTSEL_IMM_OB = 0x1e
	MDMX_FMTSEL_VEC_QH = 0x15
	MDMX_FMTSEL_VEC_OB = 0x16

	/* UDI */
	OP_SH_UDI1   = 6
	OP_MASK_UDI1 = 0x1f
	OP_SH_UDI2   = 6
	OP_MASK_UDI2 = 0x3ff
	OP_SH_UDI3   = 6
	OP_MASK_UDI3 = 0x7fff
	OP_SH_UDI4   = 6
	OP_MASK_UDI4 = 0xfffff

	/* Octeon */
	OP_SH_BBITIND   = 16
	OP_MASK_BBITIND = 0x1f
	OP_SH_CINSPOS   = 6
	OP_MASK_CINSPOS = 0x1f
	OP_SH_CINSLM1   = 11
	OP_MASK_CINSLM1 = 0x1f
	OP_SH_SEQI      = 6
	OP_MASK_SEQI    = 0x3ff

	/* Loongson */
	OP_SH_OFFSET_A   = 6
	OP_MASK_OFFSET_A = 0xff
	OP_SH_OFFSET_B   = 3
	OP_MASK_OFFSET_B = 0xff
	OP_SH_OFFSET_C   = 6
	OP_MASK_OFFSET_C = 0x1ff
	OP_SH_RZ         = 0
	OP_MASK_RZ       = 0x1f
	OP_SH_FZ         = 0
	OP_MASK_FZ       = 0x1f

	/* Every MICROMIPSOP_X definition requires a corresponding OP_X
	   definition, and vice versa.  This simplifies various parts
	   of the operand handling in GAS.  The fields below only exist
	   in the microMIPS encoding, so define each one to have an empty
	   range.  */
	OP_MASK_TRAP     = 0
	OP_SH_TRAP       = 0
	OP_MASK_OFFSET10 = 0
	OP_SH_OFFSET10   = 0
	OP_MASK_RS3      = 0
	OP_SH_RS3        = 0
	OP_MASK_MB       = 0
	OP_SH_MB         = 0
	OP_MASK_MC       = 0
	OP_SH_MC         = 0
	OP_MASK_MD       = 0
	OP_SH_MD         = 0
	OP_MASK_ME       = 0
	OP_SH_ME         = 0
	OP_MASK_MF       = 0
	OP_SH_MF         = 0
	OP_MASK_MG       = 0
	OP_SH_MG         = 0
	OP_MASK_MH       = 0
	OP_SH_MH         = 0
	OP_MASK_MJ       = 0
	OP_SH_MJ         = 0
	OP_MASK_ML       = 0
	OP_SH_ML         = 0
	OP_MASK_MM       = 0
	OP_SH_MM         = 0
	OP_MASK_MN       = 0
	OP_SH_MN         = 0
	OP_MASK_MP       = 0
	OP_SH_MP         = 0
	OP_MASK_MQ       = 0
	OP_SH_MQ         = 0
	OP_MASK_IMMA     = 0
	OP_SH_IMMA       = 0
	OP_MASK_IMMB     = 0
	OP_SH_IMMB       = 0
	OP_MASK_IMMC     = 0
	OP_SH_IMMC       = 0
	OP_MASK_IMMF     = 0
	OP_SH_IMMF       = 0
	OP_MASK_IMMG     = 0
	OP_SH_IMMG       = 0
	OP_MASK_IMMH     = 0
	OP_SH_IMMH       = 0
	OP_MASK_IMMI     = 0
	OP_SH_IMMI       = 0
	OP_MASK_IMMJ     = 0
	OP_SH_IMMJ       = 0
	OP_MASK_IMML     = 0
	OP_SH_IMML       = 0
	OP_MASK_IMMM     = 0
	OP_SH_IMMM       = 0
	OP_MASK_IMMN     = 0
	OP_SH_IMMN       = 0
	OP_MASK_IMMO     = 0
	OP_SH_IMMO       = 0
	OP_MASK_IMMP     = 0
	OP_SH_IMMP       = 0
	OP_MASK_IMMQ     = 0
	OP_SH_IMMQ       = 0
	OP_MASK_IMMU     = 0
	OP_SH_IMMU       = 0
	OP_MASK_IMMW     = 0
	OP_SH_IMMW       = 0
	OP_MASK_IMMX     = 0
	OP_SH_IMMX       = 0
	OP_MASK_IMMY     = 0
	OP_SH_IMMY       = 0
)

/* This is a list of macro expanded instructions.

   _I appended means immediate
   _A appended means target address of a jump
   _AB appended means address with (possibly zero) base register
   _D appended means 64 bit floating point constant
   _S appended means 32 bit floating point constant.  */

const (
	M_ABS int = iota
	M_ACLR_AB
	M_ADD_I
	M_ADDU_I
	M_AND_I
	M_ASET_AB
	M_BALIGN
	M_BC1FL
	M_BC1TL
	M_BC2FL
	M_BC2TL
	M_BEQ
	M_BEQ_I
	M_BEQL
	M_BEQL_I
	M_BGE
	M_BGEL
	M_BGE_I
	M_BGEL_I
	M_BGEU
	M_BGEUL
	M_BGEU_I
	M_BGEUL_I
	M_BGEZ
	M_BGEZL
	M_BGEZALL
	M_BGT
	M_BGTL
	M_BGT_I
	M_BGTL_I
	M_BGTU
	M_BGTUL
	M_BGTU_I
	M_BGTUL_I
	M_BGTZ
	M_BGTZL
	M_BLE
	M_BLEL
	M_BLE_I
	M_BLEL_I
	M_BLEU
	M_BLEUL
	M_BLEU_I
	M_BLEUL_I
	M_BLEZ
	M_BLEZL
	M_BLT
	M_BLTL
	M_BLT_I
	M_BLTL_I
	M_BLTU
	M_BLTUL
	M_BLTU_I
	M_BLTUL_I
	M_BLTZ
	M_BLTZL
	M_BLTZALL
	M_BNE
	M_BNEL
	M_BNE_I
	M_BNEL_I
	M_CACHE_AB
	M_CACHEE_AB
	M_DABS
	M_DADD_I
	M_DADDU_I
	M_DDIV_3
	M_DDIV_3I
	M_DDIVU_3
	M_DDIVU_3I
	M_DIV_3
	M_DIV_3I
	M_DIVU_3
	M_DIVU_3I
	M_DLA_AB
	M_DLCA_AB
	M_DLI
	M_DMUL
	M_DMUL_I
	M_DMULO
	M_DMULO_I
	M_DMULOU
	M_DMULOU_I
	M_DREM_3
	M_DREM_3I
	M_DREMU_3
	M_DREMU_3I
	M_DSUB_I
	M_DSUBU_I
	M_DSUBU_I_2
	M_J_A
	M_JAL_1
	M_JAL_2
	M_JAL_A
	M_JALS_1
	M_JALS_2
	M_JALS_A
	M_JRADDIUSP
	M_JRC
	M_L_DAB
	M_LA_AB
	M_LB_AB
	M_LBE_AB
	M_LBU_AB
	M_LBUE_AB
	M_LCA_AB
	M_LD_AB
	M_LDC1_AB
	M_LDC2_AB
	M_LQC2_AB
	M_LDC3_AB
	M_LDL_AB
	M_LDM_AB
	M_LDP_AB
	M_LDR_AB
	M_LH_AB
	M_LHE_AB
	M_LHU_AB
	M_LHUE_AB
	M_LI
	M_LI_D
	M_LI_DD
	M_LI_S
	M_LI_SS
	M_LL_AB
	M_LLD_AB
	M_LLDP_AB
	M_LLE_AB
	M_LLWP_AB
	M_LLWPE_AB
	M_LQ_AB
	M_LW_AB
	M_LWE_AB
	M_LWC0_AB
	M_LWC1_AB
	M_LWC2_AB
	M_LWC3_AB
	M_LWL_AB
	M_LWLE_AB
	M_LWM_AB
	M_LWP_AB
	M_LWR_AB
	M_LWRE_AB
	M_LWU_AB
	M_MSGSND
	M_MSGLD
	M_MSGLD_T
	M_MSGWAIT
	M_MSGWAIT_T
	M_MOVE
	M_MOVEP
	M_MUL
	M_MUL_I
	M_MULO
	M_MULO_I
	M_MULOU
	M_MULOU_I
	M_NOR_I
	M_OR_I
	M_PREF_AB
	M_PREFE_AB
	M_REM_3
	M_REM_3I
	M_REMU_3
	M_REMU_3I
	M_DROL
	M_ROL
	M_DROL_I
	M_ROL_I
	M_DROR
	M_ROR
	M_DROR_I
	M_ROR_I
	M_S_DA
	M_S_DAB
	M_S_S
	M_SAA_AB
	M_SAAD_AB
	M_SC_AB
	M_SCD_AB
	M_SCDP_AB
	M_SCE_AB
	M_SCWP_AB
	M_SCWPE_AB
	M_SD_AB
	M_SDC1_AB
	M_SDC2_AB
	M_SQC2_AB
	M_SDC3_AB
	M_SDL_AB
	M_SDM_AB
	M_SDP_AB
	M_SDR_AB
	M_SEQ
	M_SEQ_I
	M_SGE
	M_SGE_I
	M_SGEU
	M_SGEU_I
	M_SGT
	M_SGT_I
	M_SGTU
	M_SGTU_I
	M_SLE
	M_SLE_I
	M_SLEU
	M_SLEU_I
	M_SLT_I
	M_SLTU_I
	M_SNE
	M_SNE_I
	M_SB_AB
	M_SBE_AB
	M_SH_AB
	M_SHE_AB
	M_SQ_AB
	M_SW_AB
	M_SWE_AB
	M_SWC0_AB
	M_SWC1_AB
	M_SWC2_AB
	M_SWC3_AB
	M_SWL_AB
	M_SWLE_AB
	M_SWM_AB
	M_SWP_AB
	M_SWR_AB
	M_SWRE_AB
	M_SUB_I
	M_SUBU_I
	M_SUBU_I_2
	M_TEQ_I
	M_TGE_I
	M_TGEU_I
	M_TLT_I
	M_TLTU_I
	M_TNE_I
	M_TRUNCWD
	M_TRUNCWS
	M_ULD_AB
	M_ULH_AB
	M_ULHU_AB
	M_ULW_AB
	M_USH_AB
	M_USW_AB
	M_USD_AB
	M_XOR_I
	M_COP0
	M_COP1
	M_COP2
	M_COP3
	M_NUM_MACROS
)

/* The order of overloaded instructions matters.  Label arguments and
   register arguments look the same. Instructions that can have either
   for arguments must apear in the correct order in this table for the
   assembler to pick the right one. In other words, entries with
   immediate operands must apear after the same instruction with
   registers.

   Many instructions are short hand for other instructions (i.e., The
   jal <register> instruction is short for jalr <register>).  */

/* These are the characters which may appear in the args field of an
   instruction.  They appear in the order in which the fields appear
   when the instruction is used.  Commas and parentheses in the args
   string are ignored when assembling, and written into the output
   when disassembling.

   Each of these characters corresponds to a mask field defined above.

   "1" 5 bit sync type (OP_*_STYPE)
   "<" 5 bit shift amount (OP_*_SHAMT)
   ">" shift amount between 32 and 63, stored after subtracting 32 (OP_*_SHAMT)
   "a" 26 bit target address (OP_*_TARGET)
   "+i" likewise, but flips bit 0
   "b" 5 bit base register (OP_*_RS)
   "c" 10 bit breakpoint code (OP_*_CODE)
   "d" 5 bit destination register specifier (OP_*_RD)
   "h" 5 bit prefx hint (OP_*_PREFX)
   "i" 16 bit unsigned immediate (OP_*_IMMEDIATE)
   "j" 16 bit signed immediate (OP_*_DELTA)
   "k" 5 bit cache opcode in target register position (OP_*_CACHE)
   "o" 16 bit signed offset (OP_*_DELTA)
   "p" 16 bit PC relative branch target address (OP_*_DELTA)
   "q" 10 bit extra breakpoint code (OP_*_CODE2)
   "r" 5 bit same register used as both source and target (OP_*_RS)
   "s" 5 bit source register specifier (OP_*_RS)
   "t" 5 bit target register (OP_*_RT)
   "u" 16 bit upper 16 bits of address (OP_*_IMMEDIATE)
   "v" 5 bit same register used as both source and destination (OP_*_RS)
   "w" 5 bit same register used as both target and destination (OP_*_RT)
   "U" 5 bit same destination register in both OP_*_RD and OP_*_RT
       (used by clo and clz)
   "C" 25 bit coprocessor function code (OP_*_COPZ)
   "B" 20 bit syscall/breakpoint function code (OP_*_CODE20)
   "J" 19 bit wait function code (OP_*_CODE19)
   "x" accept and ignore register name
   "z" must be zero register
   "K" 5 bit Hardware Register (rdhwr instruction) (OP_*_RD)
   "+A" 5 bit ins/ext/dins/dext/dinsm/dextm position, which becomes
        LSB (OP_*_SHAMT; OP_*_EXTLSB or OP_*_STYPE may be used for
        microMIPS compatibility).
	Enforces: 0 <= pos < 32.
   "+B" 5 bit ins/dins size, which becomes MSB (OP_*_INSMSB).
	Requires that "+A" or "+E" occur first to set position.
	Enforces: 0 < (pos+size) <= 32.
   "+C" 5 bit ext/dext size, which becomes MSBD (OP_*_EXTMSBD).
	Requires that "+A" or "+E" occur first to set position.
	Enforces: 0 < (pos+size) <= 32.
	(Also used by "dext" w/ different limits, but limits for
	that are checked by the M_DEXT macro.)
   "+E" 5 bit dinsu/dextu position, which becomes LSB-32 (OP_*_SHAMT).
	Enforces: 32 <= pos < 64.
   "+F" 5 bit "dinsm/dinsu" size, which becomes MSB-32 (OP_*_INSMSB).
	Requires that "+A" or "+E" occur first to set position.
	Enforces: 32 < (pos+size) <= 64.
   "+G" 5 bit "dextm" size, which becomes MSBD-32 (OP_*_EXTMSBD).
	Requires that "+A" or "+E" occur first to set position.
	Enforces: 32 < (pos+size) <= 64.
   "+H" 5 bit "dextu" size, which becomes MSBD (OP_*_EXTMSBD).
	Requires that "+A" or "+E" occur first to set position.
	Enforces: 32 < (pos+size) <= 64.

   Floating point instructions:
   "D" 5 bit destination register (OP_*_FD)
   "M" 3 bit compare condition code (OP_*_CCC) (only used for mips4 and up)
   "N" 3 bit branch condition code (OP_*_BCC) (only used for mips4 and up)
   "S" 5 bit fs source 1 register (OP_*_FS)
   "T" 5 bit ft source 2 register (OP_*_FT)
   "R" 5 bit fr source 3 register (OP_*_FR)
   "V" 5 bit same register used as floating source and destination (OP_*_FS)
   "W" 5 bit same register used as floating target and destination (OP_*_FT)

   Coprocessor instructions:
   "E" 5 bit target register (OP_*_RT)
   "G" 5 bit destination register (OP_*_RD)
   "H" 3 bit sel field for (d)mtc* and (d)mfc* (OP_*_SEL)
   "P" 5 bit performance-monitor register (OP_*_PERFREG)
   "e" 5 bit vector register byte specifier (OP_*_VECBYTE)
   "%" 3 bit immediate vr5400 vector alignment operand (OP_*_VECALIGN)

   Macro instructions:
   "A" General 32 bit expression
   "I" 32 bit immediate (value placed in imm_expr).
   "F" 64 bit floating point constant in .rdata
   "L" 64 bit floating point constant in .lit8
   "f" 32 bit floating point constant
   "l" 32 bit floating point constant in .lit4

   MDMX and VR5400 instruction operands (note that while these use the
   FP register fields, the MDMX instructions accept both $fN and $vN names
   for the registers):
   "O"	alignment offset (OP_*_ALN)
   "Q"	vector/scalar/immediate source (OP_*_VSEL and OP_*_FT)
   "X"	destination register (OP_*_FD)
   "Y"	source register (OP_*_FS)
   "Z"	source register (OP_*_FT)

   R5900 VU0 Macromode instructions:
   "+5" 5 bit floating point register (FD)
   "+6" 5 bit floating point register (FS)
   "+7" 5 bit floating point register (FT)
   "+8" 5 bit integer register (FD)
   "+9" 5 bit integer register (FS)
   "+0" 5 bit integer register (FT)
   "+K" match an existing 4-bit channel mask starting at bit 21
   "+L" 2-bit channel index starting at bit 21
   "+M" 2-bit channel index starting at bit 23
   "+N" match an existing 2-bit channel index starting at bit 0
   "+f" 15 bit immediate for VCALLMS
   "+g" 5 bit signed immediate for VIADDI
   "+m" $ACC register (syntax only)
   "+q" $Q register (syntax only)
   "+r" $R register (syntax only)
   "+y" $I register (syntax only)
   "#+" "++" decorator in ($reg++) sequence
   "#-" "--" decorator in (--$reg) sequence

   DSP ASE usage:
   "2" 2 bit unsigned immediate for byte align (OP_*_BP)
   "3" 3 bit unsigned immediate (OP_*_SA3)
   "4" 4 bit unsigned immediate (OP_*_SA4)
   "5" 8 bit unsigned immediate (OP_*_IMM8)
   "6" 5 bit unsigned immediate (OP_*_RS)
   "7" 2 bit dsp accumulator register (OP_*_DSPACC)
   "8" 6 bit unsigned immediate (OP_*_WRDSP)
   "9" 2 bit dsp accumulator register (OP_*_DSPACC_S)
   "0" 6 bit signed immediate (OP_*_DSPSFT)
   ":" 7 bit signed immediate (OP_*_DSPSFT_7)
   "'" 6 bit unsigned immediate (OP_*_RDDSP)
   "@" 10 bit signed immediate (OP_*_IMM10)

   MT ASE usage:
   "!" 1 bit usermode flag (OP_*_MT_U)
   "$" 1 bit load high flag (OP_*_MT_H)
   "*" 2 bit dsp/smartmips accumulator register (OP_*_MTACC_T)
   "&" 2 bit dsp/smartmips accumulator register (OP_*_MTACC_D)
   "g" 5 bit coprocessor 1 and 2 destination register (OP_*_RD)
   "+t" 5 bit coprocessor 0 destination register (OP_*_RT)

   MCU ASE usage:
   "~" 12 bit offset (OP_*_OFFSET12)
   "\" 3 bit position for aset and aclr (OP_*_3BITPOS)

   VIRT ASE usage:
   "+J" 10-bit hypcall code (OP_*CODE10)

   UDI immediates:
   "+1" UDI immediate bits 6-10
   "+2" UDI immediate bits 6-15
   "+3" UDI immediate bits 6-20
   "+4" UDI immediate bits 6-25

   Octeon:
   "+x" Bit index field of bbit.  Enforces: 0 <= index < 32.
   "+X" Bit index field of bbit aliasing bbit32.  Matches if 32 <= index < 64,
	otherwise skips to next candidate.
   "+p" Position field of cins/cins32/exts/exts32. Enforces 0 <= pos < 32.
   "+P" Position field of cins/exts aliasing cins32/exts32.  Matches if
	32 <= pos < 64, otherwise skips to next candidate.
   "+Q" Immediate field of seqi/snei.  Enforces -512 <= imm < 512.
   "+s" Length-minus-one field of cins32/exts32.  Requires msb position
	of the field to be <= 31.
   "+S" Length-minus-one field of cins/exts.  Requires msb position
	of the field to be <= 63.

   Loongson-ext ASE:
   "+a" 8-bit signed offset in bit 6 (OP_*_OFFSET_A)
   "+b" 8-bit signed offset in bit 3 (OP_*_OFFSET_B)
   "+c" 9-bit signed offset in bit 6 (OP_*_OFFSET_C)
   "+z" 5-bit rz register (OP_*_RZ)
   "+Z" 5-bit fz register (OP_*_FZ)

   interAptiv MR2:
   "-m" register list for SAVE/RESTORE instruction

   Enhanced VA Scheme:
   "+j" 9-bit signed offset in bit 7 (OP_*_EVAOFFSET)

   MSA Extension:
   "+d" 5-bit MSA register (FD)
   "+e" 5-bit MSA register (FS)
   "+h" 5-bit MSA register (FT)
   "+k" 5-bit GPR at bit 6
   "+l" 5-bit MSA control register at bit 6
   "+n" 5-bit MSA control register at bit 11
   "+o" 4-bit vector element index at bit 16
   "+u" 3-bit vector element index at bit 16
   "+v" 2-bit vector element index at bit 16
   "+w" 1-bit vector element index at bit 16
   "+T" (-512 .. 511) << 0 at bit 16
   "+U" (-512 .. 511) << 1 at bit 16
   "+V" (-512 .. 511) << 2 at bit 16
   "+W" (-512 .. 511) << 3 at bit 16
   "+~" 2 bit LSA/DLSA shift amount from 1 to 4 at bit 6
   "+!" 3 bit unsigned bit position at bit 16
   "+@" 4 bit unsigned bit position at bit 16
   "+#" 6 bit unsigned bit position at bit 16
   "+$" 5 bit unsigned immediate at bit 16
   "+%" 5 bit signed immediate at bit 16
   "+^" 10 bit signed immediate at bit 11
   "+&" 0 vector element index
   "+*" 5-bit register vector element index at bit 16
   "+|" 8-bit mask at bit 16

   MIPS R6:
   "+:" 11-bit mask at bit 0
   "+'" 26 bit PC relative branch target address
   "+"" 21 bit PC relative branch target address
   "+;" 5 bit same register in both OP_*_RS and OP_*_RT
   "+I" 2bit unsigned bit position at bit 6
   "+O" 3bit unsigned bit position at bit 6
   "+R" must be program counter
   "-a" (-262144 .. 262143) << 2 at bit 0
   "-b" (-131072 .. 131071) << 3 at bit 0
   "-d" Same as destination register GP
   "-s" 5 bit source register specifier (OP_*_RS) not $0
   "-t" 5 bit source register specifier (OP_*_RT) not $0
   "-u" 5 bit source register specifier (OP_*_RT) greater than OP_*_RS
   "-v" 5 bit source register specifier (OP_*_RT) not $0 not OP_*_RS
   "-w" 5 bit source register specifier (OP_*_RT) less than or equal to OP_*_RS
   "-x" 5 bit source register specifier (OP_*_RT) greater than or
        equal to OP_*_RS
   "-y" 5 bit source register specifier (OP_*_RT) not $0 less than OP_*_RS
   "-A" symbolic offset (-262144 .. 262143) << 2 at bit 0
   "-B" symbolic offset (-131072 .. 131071) << 3 at bit 0

   GINV ASE usage:
   "+\" 2 bit Global TLB invalidate type at bit 8

   Other:
   "()" parens surrounding optional value
   ","  separates operands
   "+"  Start of extension sequence.

   Characters used so far, for quick reference when adding more:
   "1234567890"
   "%[]<>(),+-:'@!#$*&\~"
   "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
   "abcdefghijklopqrstuvwxz"

   Extension character sequences used so far ("+" followed by the
   following), for quick reference when adding more:
   "1234567890"
   "~!@#$%^&*|:'";\"
   "ABCEFGHIJKLMNOPQRSTUVWXZ"
   "abcdefghijklmnopqrstuvwxyz"

   Extension character sequences used so far ("-" followed by the
   following), for quick reference when adding more:
   "AB"
   "abdmstuvwxy"
*/

/* These are the bits which may be set in the pinfo field of an
   instructions, if it is not equal to INSN_MACRO.  */

/* Writes to operand number N.  */
const (
	INSN_WRITE_SHIFT = 0
	INSN_WRITE_1     = 0x00000001
	INSN_WRITE_2     = 0x00000002
	INSN_WRITE_ALL   = 0x00000003
	/* Reads from operand number N.  */
	INSN_READ_SHIFT = 2
	INSN_READ_1     = 0x00000004
	INSN_READ_2     = 0x00000008
	INSN_READ_3     = 0x00000010
	INSN_READ_4     = 0x00000020
	INSN_READ_ALL   = 0x0000003c
	/* Modifies general purpose register 31.  */
	INSN_WRITE_GPR_31 = 0x00000040
	/* Modifies coprocessor condition code.  */
	INSN_WRITE_COND_CODE = 0x00000080
	/* Reads coprocessor condition code.  */
	INSN_READ_COND_CODE = 0x00000100
	/* TLB operation.  */
	INSN_TLB = 0x00000200
	/* Reads coprocessor register other than floating point register.  */
	INSN_COP = 0x00000400
	/* Instruction loads value from memory.  */
	INSN_LOAD_MEMORY = 0x00000800
	/* Instruction loads value from coprocessor, (may require delay).  */
	INSN_LOAD_COPROC = 0x00001000
	/* Instruction has unconditional branch delay slot.  */
	INSN_UNCOND_BRANCH_DELAY = 0x00002000
	/* Instruction has conditional branch delay slot.  */
	INSN_COND_BRANCH_DELAY = 0x00004000
	/* Conditional branch likely: if branch not taken, insn nullified.  */
	INSN_COND_BRANCH_LIKELY = 0x00008000
	/* Moves to coprocessor register, (may require delay).  */
	INSN_COPROC_MOVE = 0x00010000
	/* Loads coprocessor register from memory, requiring delay.  */
	INSN_COPROC_MEMORY_DELAY = 0x00020000
	/* Reads the HI register.  */
	INSN_READ_HI = 0x00040000
	/* Reads the LO register.  */
	INSN_READ_LO = 0x00080000
	/* Modifies the HI register.  */
	INSN_WRITE_HI = 0x00100000
	/* Modifies the LO register.  */
	INSN_WRITE_LO = 0x00200000
	/* Not to be placed in a branch delay slot, either architecturally
	   or for ease of handling (such as with instructions that take a trap).  */
	INSN_NO_DELAY_SLOT = 0x00400000
	/* Instruction stores value into memory.  */
	INSN_STORE_MEMORY = 0x00800000
	/* Instruction uses single precision floating point.  */
	FP_S = 0x01000000
	/* Instruction uses double precision floating point.  */
	FP_D = 0x02000000
	/* Instruction is part of the tx39's integer multiply family.    */
	INSN_MULT = 0x04000000
	/* Reads general purpose register 24.  */
	INSN_READ_GPR_24 = 0x08000000
	/* Writes to general purpose register 24.  */
	INSN_WRITE_GPR_24 = 0x10000000
	/* A user-defined instruction.  */
	INSN_UDI = 0x20000000
	/* Instruction is actually a macro.  It should be ignored by the
	   disassembler, and requires special treatment by the assembler.  */
	INSN_MACRO = 0xffffffff

	/* These are the bits which may be set in the pinfo2 field of an
	   instruction. */

	/* Instruction is a simple alias (I.E. "move" for daddu/addu/or) */
	INSN2_ALIAS = 0x00000001
	/* Instruction reads MDMX accumulator. */
	INSN2_READ_MDMX_ACC = 0x00000002
	/* Instruction writes MDMX accumulator. */
	INSN2_WRITE_MDMX_ACC = 0x00000004
	/* Macro uses single-precision floating-point instructions.  This should
	   only be set for macros.  For instructions, FP_S in pinfo carries the
	   same information.  */
	INSN2_M_FP_S = 0x00000008
	/* Macro uses double-precision floating-point instructions.  This should
	   only be set for macros.  For instructions, FP_D in pinfo carries the
	   same information.  */
	INSN2_M_FP_D = 0x00000010
	/* Instruction has a branch delay slot that requires a 16-bit instruction.  */
	INSN2_BRANCH_DELAY_16BIT = 0x00000020
	/* Instruction has a branch delay slot that requires a 32-bit instruction.  */
	INSN2_BRANCH_DELAY_32BIT = 0x00000040
	/* Writes to the stack pointer ($29).  */
	INSN2_WRITE_SP = 0x00000080
	/* Reads from the stack pointer ($29).  */
	INSN2_READ_SP = 0x00000100
	/* Reads the RA ($31) register.  */
	INSN2_READ_GPR_31 = 0x00000200
	/* Reads the program counter ($pc).  */
	INSN2_READ_PC = 0x00000400
	/* Is an unconditional branch insn. */
	INSN2_UNCOND_BRANCH = 0x00000800
	/* Is a conditional branch insn. */
	INSN2_COND_BRANCH = 0x00001000
	/* Reads from $16.  This is true of the MIPS16 0x6500 nop.  */
	INSN2_READ_GPR_16 = 0x00002000
	/* Has an "\.x?y?z?w?" suffix based on mips_vu0_channel_mask.  */
	INSN2_VU0_CHANNEL_SUFFIX = 0x00004000
	/* Instruction has a forbidden slot.  */
	INSN2_FORBIDDEN_SLOT = 0x00008000
	/* Opcode table entry is for a short MIPS16 form only.  An extended
	   encoding may still exist, but with a separate opcode table entry
	   required.  In disassembly the presence of this flag in an otherwise
	   successful match against an extended instruction encoding inhibits
	   matching against any subsequent short table entry even if it does
	   not have this flag set.  A table entry matching the full extended
	   encoding is needed or otherwise the final EXTEND entry will apply,
	   for the disassembly of the prefix only.  */
	INSN2_SHORT_ONLY = 0x00010000

	/* Masks used to mark instructions to indicate which MIPS ISA level
	   they were introduced in.  INSN_ISA_MASK masks an enumeration that
	   specifies the base ISA level(s).  The remainder of a 32-bit
	   word constructed using these macros is a bitmask of the remaining
	   INSN_* values below.  */

	INSN_ISA_MASK = 0x0000001 // ful

	/* We cannot start at zero due to ISA_UNKNOWN below.  */
	INSN_ISA1    = 1
	INSN_ISA2    = 2
	INSN_ISA3    = 3
	INSN_ISA4    = 4
	INSN_ISA5    = 5
	INSN_ISA32   = 6
	INSN_ISA32R2 = 7
	INSN_ISA32R3 = 8
	INSN_ISA32R5 = 9
	INSN_ISA32R6 = 10
	INSN_ISA64   = 11
	INSN_ISA64R2 = 12
	INSN_ISA64R3 = 13
	INSN_ISA64R5 = 14
	INSN_ISA64R6 = 15
	/* Below this point the INSN_* values correspond to combinations of ISAs.
	   They are only for use in the opcodes table to indicate membership of
	   a combination of ISAs that cannot be expressed using the usual inclusion
	   ordering on the above INSN_* values.  */
	INSN_ISA3_32   = 16
	INSN_ISA3_32R2 = 17
	INSN_ISA4_32   = 18
	INSN_ISA4_32R2 = 19
	INSN_ISA5_32R2 = 20

	/* The R6 definitions shown below state that they support all previous ISAs.
	   This is not actually true as some instructions are removed in R6.
	   The problem is that the removed instructions in R6 come from different
	   ISAs.  One approach to solve this would be to describe in the membership
	   field of the opcode table the different ISAs an instruction belongs to.
	   This would require us to create a large amount of different ISA
	   combinations which is hard to manage.  A cleaner approach (which is
	   implemented here) is to say that R6 is an extension of R5 and then to
	   deal with the removed instructions by adding instruction exclusions
	   for R6 in the opcode table.  */

	/* Bit INSN_ISA<X> - 1 of INSN_UPTO<Y> is set if ISA Y includes ISA X.  */
	/*
	    ISAF(X) (1 << (INSN_ISA##X - 1))
	    INSN_UPTO1    ISAF(1)
	    INSN_UPTO2    INSN_UPTO1 | ISAF(2)
	    INSN_UPTO3    INSN_UPTO2 | ISAF(3) | ISAF(3_32) | ISAF(3_32R2)
	    INSN_UPTO4    INSN_UPTO3 | ISAF(4) | ISAF(4_32) | ISAF(4_32R2)
	    INSN_UPTO5    INSN_UPTO4 | ISAF(5) | ISAF(5_32R2)
	    INSN_UPTO32   INSN_UPTO2 | ISAF(32) | ISAF(3_32) | ISAF(4_32)
	    INSN_UPTO32R2 INSN_UPTO32 | ISAF(32R2) \
	   			| ISAF(3_32R2) | ISAF(4_32R2) | ISAF(5_32R2)
	    INSN_UPTO32R3 INSN_UPTO32R2 | ISAF(32R3)
	    INSN_UPTO32R5 INSN_UPTO32R3 | ISAF(32R5)
	    INSN_UPTO32R6 INSN_UPTO32R5 | ISAF(32R6)
	    INSN_UPTO64   INSN_UPTO5 | ISAF(64) | ISAF(32)
	    INSN_UPTO64R2 INSN_UPTO64 | ISAF(64R2) | ISAF(32R2)
	    INSN_UPTO64R3 INSN_UPTO64R2 | ISAF(64R3) | ISAF(32R3)
	    INSN_UPTO64R5 INSN_UPTO64R3 | ISAF(64R5) | ISAF(32R5)
	    INSN_UPTO64R6 INSN_UPTO64R5 | ISAF(64R6) | ISAF(32R6)

	   /* The same information in table form: bit INSN_ISA<X> - 1 of index
	      INSN_UPTO<Y> - 1 is set if ISA Y includes ISA X.  */
	/*
		static const unsigned int mips_isa_table[] = {
		  INSN_UPTO1,
		  INSN_UPTO2,
		  INSN_UPTO3,
		  INSN_UPTO4,
		  INSN_UPTO5,
		  INSN_UPTO32,
		  INSN_UPTO32R2,
		  INSN_UPTO32R3,
		  INSN_UPTO32R5,
		  INSN_UPTO32R6,
		  INSN_UPTO64,
		  INSN_UPTO64R2,
		  INSN_UPTO64R3,
		  INSN_UPTO64R5,
		  INSN_UPTO64R6
		};
		#undef ISAF
	*/
	/* Masks used for Chip specific instructions.  */
	INSN_CHIP_MASK = 0xc7ff4f60

	/* Cavium Networks Octeon instructions.  */
	INSN_OCTEON  = 0x00000800
	INSN_OCTEONP = 0x00000200
	INSN_OCTEON2 = 0x00000100
	INSN_OCTEON3 = 0x00000040

	/* MIPS R5900 instruction */
	INSN_5900 = 0x00004000

	/* MIPS R4650 instruction.  */
	INSN_4650 = 0x00010000
	/* LSI R4010 instruction.  */
	INSN_4010 = 0x00020000
	/* NEC VR4100 instruction.  */
	INSN_4100 = 0x00040000
	/* Toshiba R3900 instruction.  */
	INSN_3900 = 0x00080000
	/* MIPS R10000 instruction.  */
	INSN_10000 = 0x00100000
	/* Broadcom SB-1 instruction.  */
	INSN_SB1 = 0x00200000
	/* NEC VR4111/VR4181 instruction.  */
	INSN_4111 = 0x00400000
	/* NEC VR4120 instruction.  */
	INSN_4120 = 0x00800000
	/* NEC VR5400 instruction.  */
	INSN_5400 = 0x01000000
	/* NEC VR5500 instruction.  */
	INSN_5500 = 0x02000000

	/* ST Microelectronics Loongson 2E.  */
	INSN_LOONGSON_2E = 0x40000000
	/* ST Microelectronics Loongson 2F.  */
	INSN_LOONGSON_2F = 0x80000000
	/* RMI Xlr instruction */
	INSN_XLR = 0x00000020
	/* Imagination interAptiv MR2.  */
	INSN_INTERAPTIV_MR2 = 0x04000000

	/* DSP ASE */
	ASE_DSP   = 0x00000001
	ASE_DSP64 = 0x00000002
	/* DSP R2 ASE  */
	ASE_DSPR2 = 0x00000004
	/* Enhanced VA Scheme */
	ASE_EVA = 0x00000008
	/* MCU (MicroController) ASE */
	ASE_MCU = 0x00000010
	/* MDMX ASE */
	ASE_MDMX = 0x00000020
	/* MIPS-3D ASE */
	ASE_MIPS3D = 0x00000040
	/* MT ASE */
	ASE_MT = 0x00000080
	/* SmartMIPS ASE  */
	ASE_SMARTMIPS = 0x00000100
	/* Virtualization ASE */
	ASE_VIRT   = 0x00000200
	ASE_VIRT64 = 0x00000400
	/* MSA Extension  */
	ASE_MSA   = 0x00000800
	ASE_MSA64 = 0x00001000
	/* eXtended Physical Address (XPA) Extension.  */
	ASE_XPA = 0x00002000
	/* DSP R3 Module.  */
	ASE_DSPR3 = 0x00004000
	/* MIPS16e2 ASE.  */
	ASE_MIPS16E2 = 0x00008000
	/* MIPS16e2 MT ASE instructions.  */
	ASE_MIPS16E2_MT = 0x00010000
	/* The Virtualization ASE has eXtended Physical Addressing (XPA)
	   instructions which are only valid when both ASEs are enabled.  */
	ASE_XPA_VIRT = 0x00020000
	/* Cyclic redundancy check (CRC) ASE.  */
	ASE_CRC   = 0x00040000
	ASE_CRC64 = 0x00080000
	/* Global INValidate Extension.  */
	ASE_GINV = 0x00100000
	/* Loongson MultiMedia extensions Instructions (MMI).  */
	ASE_LOONGSON_MMI = 0x00200000
	/* Loongson Content Address Memory (CAM).  */
	ASE_LOONGSON_CAM = 0x00400000
	/* Loongson EXTensions (EXT) instructions.  */
	ASE_LOONGSON_EXT = 0x00800000
	/* Loongson EXTensions R2 (EXT2) instructions.  */
	ASE_LOONGSON_EXT2 = 0x01000000
	/* The Enhanced VA Scheme (EVA) extension has instructions which are
	   only valid for the R6 ISA.  */
	ASE_EVA_R6 = 0x02000000

	/* MIPS ISA defines, use instead of hardcoding ISA level.  */

	ISA_UNKNOWN = 0 /* Gas internal use.  */
	ISA_MIPS1   = INSN_ISA1
	ISA_MIPS2   = INSN_ISA2
	ISA_MIPS3   = INSN_ISA3
	ISA_MIPS4   = INSN_ISA4
	ISA_MIPS5   = INSN_ISA5

	ISA_MIPS32 = INSN_ISA32
	ISA_MIPS64 = INSN_ISA64

	ISA_MIPS32R2 = INSN_ISA32R2
	ISA_MIPS32R3 = INSN_ISA32R3
	ISA_MIPS32R5 = INSN_ISA32R5
	ISA_MIPS64R2 = INSN_ISA64R2
	ISA_MIPS64R3 = INSN_ISA64R3
	ISA_MIPS64R5 = INSN_ISA64R5

	ISA_MIPS32R6 = INSN_ISA32R6
	ISA_MIPS64R6 = INSN_ISA64R6

	/* CPU defines, use instead of hardcoding processor number. Keep this
	   in sync with bfd/archures.c in order for machine selection to work.  */
	CPU_UNKNOWN        = 0 /* Gas internal use.  */
	CPU_R3000          = 3000
	CPU_R3900          = 3900
	CPU_R4000          = 4000
	CPU_R4010          = 4010
	CPU_VR4100         = 4100
	CPU_R4111          = 4111
	CPU_VR4120         = 4120
	CPU_R4300          = 4300
	CPU_R4400          = 4400
	CPU_R4600          = 4600
	CPU_R4650          = 4650
	CPU_R5000          = 5000
	CPU_VR5400         = 5400
	CPU_VR5500         = 5500
	CPU_R5900          = 5900
	CPU_R6000          = 6000
	CPU_RM7000         = 7000
	CPU_R8000          = 8000
	CPU_RM9000         = 9000
	CPU_R10000         = 10000
	CPU_R12000         = 12000
	CPU_R14000         = 14000
	CPU_R16000         = 16000
	CPU_MIPS16         = 16
	CPU_MIPS32         = 32
	CPU_MIPS32R2       = 33
	CPU_MIPS32R3       = 34
	CPU_MIPS32R5       = 36
	CPU_MIPS32R6       = 37
	CPU_MIPS5          = 5
	CPU_MIPS64         = 64
	CPU_MIPS64R2       = 65
	CPU_MIPS64R3       = 66
	CPU_MIPS64R5       = 68
	CPU_MIPS64R6       = 69
	CPU_SB1            = 12310201 /* octal 'SB', 01.  */
	CPU_LOONGSON_2E    = 3001
	CPU_LOONGSON_2F    = 3002
	CPU_GS464          = 3003
	CPU_GS464E         = 3004
	CPU_GS264E         = 3005
	CPU_OCTEON         = 6501
	CPU_OCTEONP        = 6601
	CPU_OCTEON2        = 6502
	CPU_OCTEON3        = 6503
	CPU_XLR            = 887682 /* decimal 'XLR'   */
	CPU_INTERAPTIV_MR2 = 736550 /* decimal 'IA2'  */
)

/* Return true if the given CPU is included in INSN_* mask MASK.  */
