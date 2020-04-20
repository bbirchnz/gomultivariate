package vector

import (
	"reflect"
	"testing"
)

func TestVector32_EAdd(t *testing.T) {
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
			if got := tt.a.EAdd(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Add() = %v, want %v", got, tt.want)
			}
			if got := tt.a.EAdd(tt.args.b); reflect.DeepEqual(got, tt.a) {
				t.Errorf("Vector.Add() = %v, should not equal a: %v", got, tt.a)
			}
		})
	}
}

func TestVector32_ESub(t *testing.T) {
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
			if got := tt.a.ESub(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_EMul(t *testing.T) {
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
			if got := tt.a.EMul(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_EMax(t *testing.T) {
	type args struct {
		b Vector32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{0, 0}, args: args{Vector32{-5, 5}}, want: Vector32{0, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.EMax(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector32.EMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_EMin(t *testing.T) {
	type args struct {
		b Vector32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{0, 0}, args: args{Vector32{-5, 5}}, want: Vector32{-5, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.EMin(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector32.EMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
