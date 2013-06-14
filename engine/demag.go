package engine

import (
	"code.google.com/p/mx3/cuda"
	"code.google.com/p/mx3/data"
)

var (
	Msat        ScalarParam                            // Saturation magnetization in A/m
	bsat        = scalarParam("Bsat", "T", nil)        // automatically derived from Msat, never zero
	B_demag     setterQuant                            // demag field in Tesla
	FFTM        fftm                                   // FFT of m
	EnableDemag bool                            = true // enable/disable demag field
)

var demag_ *cuda.DemagConvolution // does the heavy lifting and provides FFTM

func init() {
	Msat = scalarParam("Msat", "A/m", func(r int) {
		msat := Msat.GetRegion(r)
		bsat.setRegion(r, msat*Mu0)
		ku1_red.setRegion(r, safediv(Ku1.GetRegion(r), msat))
		dmi_red.setRegion(r, safediv(DMI.GetRegion(r), msat))
	})

	world.Var("EnableDemag", &EnableDemag)
	fftm_ := &FFTM
	world.ROnly("mFFT", &fftm_)
	B_demag_ := &B_demag
	world.ROnly("B_demag", &B_demag_)
	world.LValue("Msat", &Msat)
}

func initDemag() {
	demag_ = cuda.NewDemag(Mesh())
	B_demag = setter(3, Mesh(), "B_demag", "T", func(b *data.Slice, cansave bool) {
		if EnableDemag {
			demag_.Exec(b, M.buffer, bsat.Gpu(), regions.Gpu())
		} else {
			cuda.Zero(b)
		}
	})
	Quants["B_demag"] = &B_demag
}