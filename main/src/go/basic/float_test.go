package main

import (
	"math"
	"testing"
)

func Test_minMax(t *testing.T) {
	t.Log(math.MaxFloat32)             // 3.40282346638528859811704183484516925440e+38
	t.Log(math.SmallestNonzeroFloat32) // 1.401298464324817070923729583289916131280e-45
	t.Log(math.MaxFloat64)             // 1.79769313486231570814527423731704356798070e+308
	t.Log(math.SmallestNonzeroFloat64) // 4.9406564584124654417656879286822137236505980e-324
}

func Test_outputFEEE754float32(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0.15625", args{float32(0.15625)}, "0-01111100-01000000000000000000000"},
		{"-0.15625", args{float32(-0.15625)}, "1-01111100-01000000000000000000000"},
		{"2.75", args{float32(2.75)}, "0-10000000-01100000000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outputFEEE754float32(tt.args.f); got != tt.want {
				t.Errorf("outputFEEE754float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_outputFEEE754float64(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0.15625", args{float64(0.15625)}, "0-01111111100-0100000000000000000000000000000000000000000000000000"},
		{"-0.15625", args{float64(-0.15625)}, "1-01111111100-0100000000000000000000000000000000000000000000000000"},
		{"2.75", args{float64(2.75)}, "0-10000000000-0110000000000000000000000000000000000000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outputFEEE754float64(tt.args.f); got != tt.want {
				t.Errorf("outputFEEE754float64() = %v, want %v", got, tt.want)
			}
		})
	}
}
