package linutil

import (
	"fmt"

	"github.com/go-delve/delve/pkg/proc"
)

// Regs is a wrapper for sys.PtraceRegs.
type Mips64Registers struct {
	Regs     *Mips64PtraceRegs // general-purpose registers
	Fpregs   []proc.Register   // Formatted floating point registers
	Fpregset []byte            // holding all floating point register values

	loadFpRegs func(*Mips64Registers) error
}

func NewMips64Registers(regs *Mips64PtraceRegs, loadFpRegs func(*Mips64Registers) error) *Mips64Registers {
	return &Mips64Registers{Regs: regs, loadFpRegs: loadFpRegs}
}

// Mips64PtraceRegs is the struct used by the linux kernel to return the
// general purpose registers for Mips64 CPUs.
// copy from sys/unix/ztypes_linux_mips64.go:735
type Mips64PtraceRegs struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
}

// Slice returns the registers as a list of (name, value) pairs.
func (r *Mips64Registers) Slice(floatingPoint bool) ([]proc.Register, error) {
	var regs64 = []struct {
		k string
		v uint64
	}{
		{"$0", r.Regs.Regs[0]},
		{"$1", r.Regs.Regs[1]},
		{"$2", r.Regs.Regs[2]},
		{"$3", r.Regs.Regs[3]},
		{"$4", r.Regs.Regs[4]},
		{"$5", r.Regs.Regs[5]},
		{"$6", r.Regs.Regs[6]},
		{"$7", r.Regs.Regs[7]},
		{"$8", r.Regs.Regs[8]},
		{"$9", r.Regs.Regs[9]},
		{"$10", r.Regs.Regs[10]},
		{"$11", r.Regs.Regs[11]},
		{"$12", r.Regs.Regs[12]},
		{"$13", r.Regs.Regs[13]},
		{"$14", r.Regs.Regs[14]},
		{"$15", r.Regs.Regs[15]},
		{"$16", r.Regs.Regs[16]},
		{"$17", r.Regs.Regs[17]},
		{"$18", r.Regs.Regs[18]},
		{"$19", r.Regs.Regs[19]},
		{"$20", r.Regs.Regs[20]},
		{"$21", r.Regs.Regs[21]},
		{"$22", r.Regs.Regs[22]},
		{"$23", r.Regs.Regs[23]},
		{"$24", r.Regs.Regs[24]},
		{"$25", r.Regs.Regs[25]},
		{"$26", r.Regs.Regs[26]},
		{"$27", r.Regs.Regs[27]},
		{"$28", r.Regs.Regs[28]},
		{"$29", r.Regs.Regs[29]},
		{"$30", r.Regs.Regs[30]},
		{"$31", r.Regs.Regs[31]},
		{"Lo", r.Regs.Lo},
		{"Hi", r.Regs.Hi},
		{"Epc", r.Regs.Epc},
		{"Badvaddr", r.Regs.Badvaddr},
		{"Status", r.Regs.Status},
		{"Cause", r.Regs.Cause},
	}
	out := make([]proc.Register, 0, len(regs64)+len(r.Fpregs))
	for _, reg := range regs64 {
		out = proc.AppendUint64Register(out, reg.k, reg.v)
	}
	var floatLoadError error
	if floatingPoint {
		if r.loadFpRegs != nil {
			floatLoadError = r.loadFpRegs(r)
			r.loadFpRegs = nil
		}
		out = append(out, r.Fpregs...)
	}
	return out, floatLoadError
}

// PC returns the value of RIP register.
func (r *Mips64Registers) PC() uint64 {
	return r.Regs.Epc
}

// SP returns the value of RSP register.
func (r *Mips64Registers) SP() uint64 {
	return r.Regs.Regs[29]
}

func (r *Mips64Registers) BP() uint64 {
	return r.Regs.Regs[30]
}

// TLS returns the address of the thread local storage memory segment.
func (r *Mips64Registers) TLS() uint64 {
	return 0
}

// GAddr returns the address of the G variable if it is known, 0 and false
// otherwise.
func (r *Mips64Registers) GAddr() (uint64, bool) {
	return r.Regs.Regs[28], true
}

// Get returns the value of the n-th register (in mips64asm order).
func (r *Mips64Registers) Get(n int) (uint64, error) {
	if n >= 0 && n <= 32 {
		return r.Regs.Regs[n], nil
	}

	return 0, proc.ErrUnknownRegister
}

// Copy returns a copy of these registers that is guarenteed not to change.
func (r *Mips64Registers) Copy() (proc.Registers, error) {
	if r.loadFpRegs != nil {
		err := r.loadFpRegs(r)
		r.loadFpRegs = nil
		if err != nil {
			return nil, err
		}
	}
	var rr Mips64Registers
	rr.Regs = &Mips64PtraceRegs{}
	*(rr.Regs) = *(r.Regs)
	if r.Fpregs != nil {
		rr.Fpregs = make([]proc.Register, len(r.Fpregs))
		copy(rr.Fpregs, r.Fpregs)
	}
	if r.Fpregset != nil {
		rr.Fpregset = make([]byte, len(r.Fpregset))
		copy(rr.Fpregset, r.Fpregset)
	}
	return &rr, nil
}

type Mips64PtraceFpRegs struct {
	Regs [32]uint64
	Fpsr uint32 // TODO
	Fpcr uint32 // TODO
}

func (fpregs *Mips64PtraceFpRegs) Decode() (regs []proc.Register) {
	for i := 0; i < len(fpregs.Regs); i++ {
		regs = proc.AppendUint64Register(regs, fmt.Sprintf("f%d", i), fpregs.Regs[i])
	}
	return
}

func (fpregs *Mips64PtraceFpRegs) Byte() []byte {
	//	fpregs.Regs = make([]uint64, 32)
	//	return fpregs.Regs[:]
	return nil
}
