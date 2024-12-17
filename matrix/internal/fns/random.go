package fns

import (
	"math/rand"

	"matrix/internal/pb"
)

func rand_layout() pb.Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Layout_LAYOUT_AZERTY
	case 2:
		return pb.Layout_LAYOUT_QWERTY
	case 3:
		return pb.Layout_LAYOUT_QWERTZ
	default:
		return pb.Layout_LAYOUT_UNSPECIFIED
	}
}

func rand_bool() bool {
	return rand.Intn(2) == 1
}

func rand_string(s ...string) string {
	if len(s) == 0 {
		return ""
	}

	return s[rand.Intn(len(s))]
}

func rand_int32(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func rand_int64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func rand_uint64(min, max uint64) uint64 {
	return min + rand.Uint64()%(max-min+1)
}

func rand_float64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func rand_cpu_brand() string {
	return rand_string(CpuBrands...)
}

func rand_cpu_name(brand string) string {
	if brand == "Intel" {
		return rand_string(IntelCPUs...)
	}

	return rand_string(AmdCPUs...)
}

func rand_gpu_brand() string {
	return rand_string(GpuBrands...)
}

func rand_gpu_name(brand string) string {
	if brand == "Nvdia" {
		return rand_string(NvidiaGPUs...)
	}

	return rand_string(AmdGPUs...)
}
