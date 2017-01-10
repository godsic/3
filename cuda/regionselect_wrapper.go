package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/godsic/3/cuda/cu"
	"github.com/godsic/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for regionselect kernel
var regionselect_code cu.Function

// Stores the arguments for regionselect kernel invocation
type regionselect_args_t struct {
	arg_dst     unsafe.Pointer
	arg_src     unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_region  byte
	arg_N       int
	argptr      [5]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for regionselect kernel invocation
var regionselect_args regionselect_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	regionselect_args.argptr[0] = unsafe.Pointer(&regionselect_args.arg_dst)
	regionselect_args.argptr[1] = unsafe.Pointer(&regionselect_args.arg_src)
	regionselect_args.argptr[2] = unsafe.Pointer(&regionselect_args.arg_regions)
	regionselect_args.argptr[3] = unsafe.Pointer(&regionselect_args.arg_region)
	regionselect_args.argptr[4] = unsafe.Pointer(&regionselect_args.arg_N)
}

// Wrapper for regionselect CUDA kernel, asynchronous.
func k_regionselect_async(dst unsafe.Pointer, src unsafe.Pointer, regions unsafe.Pointer, region byte, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("regionselect")
	}

	regionselect_args.Lock()
	defer regionselect_args.Unlock()

	if regionselect_code == 0 {
		regionselect_code = fatbinLoad(regionselect_map, "regionselect")
	}

	regionselect_args.arg_dst = dst
	regionselect_args.arg_src = src
	regionselect_args.arg_regions = regions
	regionselect_args.arg_region = region
	regionselect_args.arg_N = N

	args := regionselect_args.argptr[:]
	cu.LaunchKernel(regionselect_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("regionselect")
	}
}

// maps compute capability on PTX code for regionselect kernel.
var regionselect_map = map[int]string{0: "",
	20: regionselect_ptx_20,
	30: regionselect_ptx_30,
	35: regionselect_ptx_35,
	50: regionselect_ptx_50,
	52: regionselect_ptx_52,
	53: regionselect_ptx_53}

// regionselect PTX code for various compute capabilities.
const (
	regionselect_ptx_20 = `
.version 4.3
.target sm_20
.address_size 64

	// .globl	regionselect

.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<3>;
	.reg .f32 	%f<5>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd1, [regionselect_param_0];
	ld.param.u64 	%rd2, [regionselect_param_1];
	ld.param.u64 	%rd3, [regionselect_param_2];
	ld.param.u32 	%r2, [regionselect_param_4];
	ld.param.u8 	%rs1, [regionselect_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_4;

	cvta.to.global.u64 	%rd4, %rd3;
	cvt.s64.s32	%rd5, %r1;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.u8 	%rs2, [%rd6];
	mov.f32 	%f4, 0f00000000;
	setp.ne.s16	%p2, %rs2, %rs1;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd7, %rd2;
	mul.wide.s32 	%rd8, %r1, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.f32 	%f4, [%rd9];

BB0_3:
	cvta.to.global.u64 	%rd10, %rd1;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	st.global.f32 	[%rd12], %f4;

BB0_4:
	ret;
}


`
	regionselect_ptx_30 = `
.version 4.3
.target sm_30
.address_size 64

	// .globl	regionselect

.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<3>;
	.reg .f32 	%f<5>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd1, [regionselect_param_0];
	ld.param.u64 	%rd2, [regionselect_param_1];
	ld.param.u64 	%rd3, [regionselect_param_2];
	ld.param.u32 	%r2, [regionselect_param_4];
	ld.param.u8 	%rs1, [regionselect_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_4;

	cvta.to.global.u64 	%rd4, %rd3;
	cvt.s64.s32	%rd5, %r1;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.u8 	%rs2, [%rd6];
	mov.f32 	%f4, 0f00000000;
	setp.ne.s16	%p2, %rs2, %rs1;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd7, %rd2;
	mul.wide.s32 	%rd8, %r1, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.f32 	%f4, [%rd9];

BB0_3:
	cvta.to.global.u64 	%rd10, %rd1;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	st.global.f32 	[%rd12], %f4;

BB0_4:
	ret;
}


`
	regionselect_ptx_35 = `
.version 4.3
.target sm_35
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	regionselect
.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<4>;
	.reg .f32 	%f<5>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd1, [regionselect_param_0];
	ld.param.u64 	%rd2, [regionselect_param_1];
	ld.param.u64 	%rd3, [regionselect_param_2];
	ld.param.u32 	%r2, [regionselect_param_4];
	ld.param.u8 	%rs1, [regionselect_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB6_4;

	cvta.to.global.u64 	%rd4, %rd3;
	cvt.s64.s32	%rd5, %r1;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.u8 	%rs2, [%rd6];
	mov.f32 	%f4, 0f00000000;
	setp.ne.s16	%p2, %rs2, %rs1;
	@%p2 bra 	BB6_3;

	cvta.to.global.u64 	%rd7, %rd2;
	mul.wide.s32 	%rd8, %r1, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f4, [%rd9];

BB6_3:
	cvta.to.global.u64 	%rd10, %rd1;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	st.global.f32 	[%rd12], %f4;

BB6_4:
	ret;
}


`
	regionselect_ptx_50 = `
.version 4.3
.target sm_50
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	regionselect
.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<4>;
	.reg .f32 	%f<5>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd1, [regionselect_param_0];
	ld.param.u64 	%rd2, [regionselect_param_1];
	ld.param.u64 	%rd3, [regionselect_param_2];
	ld.param.u32 	%r2, [regionselect_param_4];
	ld.param.u8 	%rs1, [regionselect_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB6_4;

	cvta.to.global.u64 	%rd4, %rd3;
	cvt.s64.s32	%rd5, %r1;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.u8 	%rs2, [%rd6];
	mov.f32 	%f4, 0f00000000;
	setp.ne.s16	%p2, %rs2, %rs1;
	@%p2 bra 	BB6_3;

	cvta.to.global.u64 	%rd7, %rd2;
	mul.wide.s32 	%rd8, %r1, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f4, [%rd9];

BB6_3:
	cvta.to.global.u64 	%rd10, %rd1;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	st.global.f32 	[%rd12], %f4;

BB6_4:
	ret;
}


`
	regionselect_ptx_52 = `
.version 4.3
.target sm_52
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	regionselect
.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<4>;
	.reg .f32 	%f<5>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd1, [regionselect_param_0];
	ld.param.u64 	%rd2, [regionselect_param_1];
	ld.param.u64 	%rd3, [regionselect_param_2];
	ld.param.u32 	%r2, [regionselect_param_4];
	ld.param.u8 	%rs1, [regionselect_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB6_4;

	cvta.to.global.u64 	%rd4, %rd3;
	cvt.s64.s32	%rd5, %r1;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.u8 	%rs2, [%rd6];
	mov.f32 	%f4, 0f00000000;
	setp.ne.s16	%p2, %rs2, %rs1;
	@%p2 bra 	BB6_3;

	cvta.to.global.u64 	%rd7, %rd2;
	mul.wide.s32 	%rd8, %r1, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f4, [%rd9];

BB6_3:
	cvta.to.global.u64 	%rd10, %rd1;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	st.global.f32 	[%rd12], %f4;

BB6_4:
	ret;
}


`
	regionselect_ptx_53 = `
.version 4.3
.target sm_53
.address_size 64

	// .weak	cudaMalloc

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaFuncGetAttributes
.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaDeviceGetAttribute
.weak .func  (.param .b32 func_retval0) cudaDeviceGetAttribute(
	.param .b64 cudaDeviceGetAttribute_param_0,
	.param .b32 cudaDeviceGetAttribute_param_1,
	.param .b32 cudaDeviceGetAttribute_param_2
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaGetDevice
.weak .func  (.param .b32 func_retval0) cudaGetDevice(
	.param .b64 cudaGetDevice_param_0
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessor
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessor(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessor_param_3
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .weak	cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags
.weak .func  (.param .b32 func_retval0) cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags(
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_0,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_1,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_2,
	.param .b64 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_3,
	.param .b32 cudaOccupancyMaxActiveBlocksPerMultiprocessorWithFlags_param_4
)
{
	.reg .b32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	ret;
}

	// .globl	regionselect
.visible .entry regionselect(
	.param .u64 regionselect_param_0,
	.param .u64 regionselect_param_1,
	.param .u64 regionselect_param_2,
	.param .u8 regionselect_param_3,
	.param .u32 regionselect_param_4
)
{
	.reg .pred 	%p<3>;
	.reg .b16 	%rs<4>;
	.reg .f32 	%f<5>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<13>;


	ld.param.u64 	%rd1, [regionselect_param_0];
	ld.param.u64 	%rd2, [regionselect_param_1];
	ld.param.u64 	%rd3, [regionselect_param_2];
	ld.param.u32 	%r2, [regionselect_param_4];
	ld.param.u8 	%rs1, [regionselect_param_3];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB6_4;

	cvta.to.global.u64 	%rd4, %rd3;
	cvt.s64.s32	%rd5, %r1;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.u8 	%rs2, [%rd6];
	mov.f32 	%f4, 0f00000000;
	setp.ne.s16	%p2, %rs2, %rs1;
	@%p2 bra 	BB6_3;

	cvta.to.global.u64 	%rd7, %rd2;
	mul.wide.s32 	%rd8, %r1, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f4, [%rd9];

BB6_3:
	cvta.to.global.u64 	%rd10, %rd1;
	mul.wide.s32 	%rd11, %r1, 4;
	add.s64 	%rd12, %rd10, %rd11;
	st.global.f32 	[%rd12], %f4;

BB6_4:
	ret;
}


`
)
