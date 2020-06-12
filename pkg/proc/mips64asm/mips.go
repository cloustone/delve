package mips64asm

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

var (
	mipsCp0Names_mips3264r2 = []string{
		"c0_index", "c0_random", "c0_entrylo0", "c0_entrylo1", "c0_context", "c0_pagemask", "c0_wired", "c0_hwrena",
		"c0_badvaddr", "c0_count", "c0_entryhi", "c0_compare", "c0_status", "c0_cause", "c0_epc", "c0_prid",
		"c0_config", "c0_lladdr", "c0_watchlo", "c0_watchhi", "c0_xcontext", "$21", "$22", "c0_debug",
		"c0_depc", "c0_perfcnt", "c0_errctl", "c0_cacheerr", "c0_taglo", "c0_taghi", "c0_errorepc", "c0_desave",
	}

	mipsCp1Names_mips3264r2 = []string{
		"c1_fir", "c1_ufr", "$2", "$3", "c1_unfr", "$5", "$6", "$7",
		"$8", "$9", "$10", "$11", "$12", "$13", "$14", "$15",
		"$16", "$17", "$18", "$19", "$20", "$21", "$22", "$23",
		"$24", "c1_fccr", "c1_fexr", "$27", "c1_fenr", "$29", "$30", "c1_fcsr",
	}

	mipsCp0SelNames_mips3264r2 = []mipsCp0SelName{
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

	mipsHwrNames_mips3264r2 = []string{
		"hwr_cpunum", "hwr_synci_step", "hwr_cc", "hwr_ccres",
		"$4", "$5", "$6", "$7",
		"$8", "$9", "$10", "$11", "$12", "$13", "$14", "$15",
		"$16", "$17", "$18", "$19", "$20", "$21", "$22", "$23",
		"$24", "$25", "$26", "$27", "$28", "$29", "$30", "$31",
	}

	mipsArchOptions = []mipsArchOption{
		{
			"mips64r2", CPU_MIPS64R2, ISA_MIPS64R2,
			ASE_MIPS3D | ASE_DSP | ASE_DSP64 | ASE_EVA | ASE_MT | ASE_MCU | ASE_VIRT | ASE_VIRT64 | ASE_MSA | ASE_MSA64 | ASE_XPA,
			mipsCp0Names_mips3264r2, mipsCp0SelNames_mips3264r2, mipsCp1Names_mips3264r2, mipsHwrNames_mips3264r2,
		},
	}
)
