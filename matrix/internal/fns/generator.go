package fns

import (
	"matrix/internal/pb"
)

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:  rand_layout(),
		Backlit: rand_bool(),
	}
}

func NewCPU() *pb.CPU {
	b := rand_cpu_brand()
	cores := rand_int32(2, 8)
	min_ghz := rand_float64(2.0, 3.3)

	return &pb.CPU{
		Brand:         b,
		Name:          rand_cpu_name(b),
		NumberCores:   cores,
		NumberThreads: rand_int32(cores, 12),
		MinGhz:        min_ghz,
		MaxGhz:        rand_float64(min_ghz, 5.0),
	}
}

func NewGPU() *pb.GPU {
	b := rand_gpu_brand()
	min_ghz := rand_float64(1.0, 1.5)

	return &pb.GPU{
		Brand:  b,
		Name:   rand_gpu_name(b),
		MinGhz: min_ghz,
		MaxGhz: rand_float64(min_ghz, 2.0),
		Memory: &pb.Memory{
			Value: rand_uint64(2, 6),
			Unit:  pb.Unit_UNIT_GIGABYTE,
		},
	}
}

func NewRAM() *pb.Memory {
	return &pb.Memory{
		Value: rand_uint64(4, 64),
		Unit:  pb.Unit_UNIT_GIGABYTE,
	}
}
