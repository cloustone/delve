package mips64asm

// MipsCpuType defines detail cpu famliy
type MipsCpuType uint32

const (
	CPU_MIPS64R2 CpuType = 65
	CPU_MIPS64R3         = 66
	CPU_MIPS64R4         = 68
)

type MipsIsaType uint32

const (
	ISA_MIPS1 MipsIsaType = itoa
	ISA_MIPS2
	ISA_MIPS3
	ISA_MIPS4
	ISA_MIPS5
	ISA_MIPS64R2
	ISA_MIPS64R3
	ISA_MIPS64R4
)

const (
	ASE_DSP         = 0x00000001 // DSP ASE
	ASE_DSP64       = 0x00000002 // DSP ASE
	ASE_DSPR2       = 0x00000004 // DSP R2 ASE
	ASE_EVA         = 0x00000008 // Enhanced VA Scheme
	ASE_MCU         = 0x00000010 // MCU (MicroController) ASE
	ASE_MDMX        = 0x00000020 // MDMX ASE
	ASE_MIPS3D      = 0x00000040 // MIPS-3D ASE
	ASE_MT          = 0x00000080 // MT ASE
	ASE_SMARTMIPS   = 0x00000100 // SmartMIPS ASE
	ASE_VIRT        = 0x00000200 // Virtualization ASE
	ASE_VIRT64      = 0x00000400 // Virtualization ASE
	ASE_MSA         = 0x00000800 // MSA Extension
	ASE_MSA64       = 0x00001000 // MSA Extension
	ASE_XPA         = 0x00002000 // eXtended Physical Address (XPA) Extension
	ASE_DSPR3       = 0x00004000 // DSP R3 Module
	ASE_MIPS16E2    = 0x00008000 // MIPS16e2 ASE
	ASE_MIPS16E2_MT = 0x00010000 // MIPS16e2 MT ASE instructions
	// The Virtualization ASE has eXtended Physical Addressing (XPA)
	// instructions which are only valid when both ASEs are enabled.
	ASE_XPA_VIRT      = 0x00020000
	ASE_CRC           = 0x00040000 // Cyclic redundancy check (CRC) ASE
	ASE_CRC64         = 0x00080000 // Cyclic redundancy check (CRC) ASE
	ASE_GINV          = 0x00100000 // Global INValidate Extension
	ASE_LOONGSON_MMI  = 0x00200000 // Loongson MultiMedia extensions Instructions (MMI)
	ASE_LOONGSON_CAM  = 0x00400000 // Loongson Content Address Memory (CAM)
	ASE_LOONGSON_EXT  = 0x00800000 // Loongson EXTensions (EXT) instructions
	ASE_LOONGSON_EXT2 = 0x01000000 // Loongson EXTensions R2 (EXT2) instructions
	// The Enhanced VA Scheme (EVA) extension has instructions which are only valid for the R6 ISA.
	ASE_EVA_R6 = 0x02000000
)

type mipsAbiOption struct {
	name     string
	gprNames []string
	fprNames []string
}

type mipsCp0SelName struct {
	cp0reg int
	sel    int
	name   string
}

type mipsArchOption struct {
	name        string
	processor   int
	isa         int
	ase         int
	cp0Names    []string
	cp0SelNames []mipsCp0SelName
	cp1Names    []string
	hwrNames    []string
}

const mipsCp0Names_mips3264r2 = [32]string{
	"c0_index", "c0_random", "c0_entrylo0", "c0_entrylo1", "c0_context", "c0_pagemask", "c0_wired", "c0_hwrena",
	"c0_badvaddr", "c0_count", "c0_entryhi", "c0_compare", "c0_status", "c0_cause", "c0_epc", "c0_prid",
	"c0_config", "c0_lladdr", "c0_watchlo", "c0_watchhi", "c0_xcontext", "$21", "$22", "c0_debug",
	"c0_depc", "c0_perfcnt", "c0_errctl", "c0_cacheerr", "c0_taglo", "c0_taghi", "c0_errorepc", "c0_desave",
}

const mipsCp1Names_mips3264 = [32]string{
	"c1_fir", "c1_ufr", "$2", "$3", "c1_unfr", "$5", "$6", "$7",
	"$8", "$9", "$10", "$11", "$12", "$13", "$14", "$15",
	"$16", "$17", "$18", "$19", "$20", "$21", "$22", "$23",
	"$24", "c1_fccr", "c1_fexr", "$27", "c1_fenr", "$29", "$30", "c1_fcsr",
}

const mipsCp0selNames_mips3264 = []mips_cp0sel_name{
	{16, 1, "c0_config1"},
	{16, 2, "c0_config2"},
	{16, 3, "c0_config3"},
	{18, 1, "c0_watchlo,1"},
	{18, 2, "c0_watchlo,2"},
	{18, 3, "c0_watchlo,3"},
	{18, 4, "c0_watchlo,4"},
	{18, 5, "c0_watchlo,5"},
	{18, 6, "c0_watchlo,6"},
	{18, 7, "c0_watchlo,7"},
	{19, 1, "c0_watchhi,1"},
	{19, 2, "c0_watchhi,2"},
	{19, 3, "c0_watchhi,3"},
	{19, 4, "c0_watchhi,4"},
	{19, 5, "c0_watchhi,5"},
	{19, 6, "c0_watchhi,6"},
	{19, 7, "c0_watchhi,7"},
	{25, 1, "c0_perfcnt,1"},
	{25, 2, "c0_perfcnt,2"},
	{25, 3, "c0_perfcnt,3"},
	{25, 4, "c0_perfcnt,4"},
	{25, 5, "c0_perfcnt,5"},
	{25, 6, "c0_perfcnt,6"},
	{25, 7, "c0_perfcnt,7"},
	{27, 1, "c0_cacheerr,1"},
	{27, 2, "c0_cacheerr,2"},
	{27, 3, "c0_cacheerr,3"},
	{28, 1, "c0_datalo"},
	{29, 1, "c0_datahi"},
}

const mipsHwrNames_mips3264r2 = [32]string{
	"hwr_cpunum", "hwr_synci_step", "hwr_cc", "hwr_ccres",
	"$4", "$5", "$6", "$7",
	"$8", "$9", "$10", "$11", "$12", "$13", "$14", "$15",
	"$16", "$17", "$18", "$19", "$20", "$21", "$22", "$23",
	"$24", "$25", "$26", "$27", "$28", "$29", "$30", "$31",
}

var mipsArchOptions = []mipsArchOption{
	{
		"mips64r2", CPU_MIPS64R2, ISA_MIPS64R2,
		ASE_MIPS3D | ASE_DSP | ASE_SSPR2 | ASE_DSP64 | ASE_EVA | ASE_MT | ASE_MCU | ASE_VIRT | ASE_VIRT64 | ASE_MSA | ASE_MSA64 | ASE_XPA,
		mipsCp0Names_mips3264r2, mipsCp0SelNames_mips3264r2, mipsCp1Names_mips3264r2, mipsHwrNames_mips3264R2,
	},
}
