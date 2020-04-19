package vector

import (
	"reflect"
	"testing"
)

func TestVector32_Add(t *testing.T) {
	type args struct {
		b Vector32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{1, 2, 3}, args: args{b: Vector32{4, 5, 6}}, want: Vector32{5, 7, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Add(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_Sub(t *testing.T) {
	type args struct {
		b Vector32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{1, 2, 3}, args: args{b: Vector32{4, 5, 6}}, want: Vector32{-3, -3, -3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Sub(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_Mul(t *testing.T) {
	type args struct {
		b Vector32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{1, 2, 3}, args: args{b: Vector32{4, 5, 6}}, want: Vector32{4, 10, 18}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Mul(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_Max(t *testing.T) {
	tests := []struct {
		name string
		a    Vector32
		want float32
	}{
		{a: Vector32{1, 2, 3}, want: 3},
		{a: Vector32{3, 2, 1}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Max(); got != tt.want {
				t.Errorf("Vector.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_Min(t *testing.T) {
	tests := []struct {
		name string
		a    Vector32
		want float32
	}{
		{a: Vector32{1, 2, 3}, want: 1},
		{a: Vector32{3, 2, 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Min(); got != tt.want {
				t.Errorf("Vector.Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_String(t *testing.T) {
	tests := []struct {
		name string
		a    Vector32
		want string
	}{
		{a: Vector32{1, 2, 3}, want: "1.000,2.000,3.000,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Vector.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVector(t *testing.T) {
	type args struct {
		size  int
		value float32
	}
	tests := []struct {
		name string
		args args
		want Vector32
	}{
		{args: args{size: 5, value: 1}, want: Vector32{1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVector32(tt.args.size, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
