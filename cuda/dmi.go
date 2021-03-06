package cuda

import (
	"unsafe"

	"github.com/godsic/3/data"
	"github.com/godsic/3/util"
)

// Add effective field of Dzyaloshinskii-Moriya interaction to Beff (Tesla).
// According to Bagdanov and Röβler, PRL 87, 3, 2001. eq.8 (out-of-plane symmetry breaking).
// See dmi.cu
func AddDMI(Beff *data.Slice, m *data.Slice, Aex_red, Dex_red SymmLUT, regions *Bytes, mesh *data.Mesh) {
	cellsize := mesh.CellSize()
	N := Beff.Size()
	util.Argument(m.Size() == N)
	cfg := make3DConf(N)

	k_adddmi_async(Beff.DevPtr(X), Beff.DevPtr(Y), Beff.DevPtr(Z),
		m.DevPtr(X), m.DevPtr(Y), m.DevPtr(Z),
		unsafe.Pointer(Aex_red), unsafe.Pointer(Dex_red), regions.Ptr,
		float32(cellsize[X]*1e9), float32(cellsize[Y]*1e9), float32(cellsize[Z]*1e9), N[X], N[Y], N[Z], mesh.PBC_code(), cfg)
}
