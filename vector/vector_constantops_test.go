package vector

import (
	"reflect"
	"testing"
)

func TestVector32_CAdd(t *testing.T) {
	type args struct {
		c float32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{1, 2, 3}, args: args{2}, want: Vector32{3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.CAdd(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector32.CAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_CSub(t *testing.T) {
	type args struct {
		c float32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{1, 2, 3}, args: args{2}, want: Vector32{-1, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.CSub(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector32.CSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_CMul(t *testing.T) {
	type args struct {
		c float32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{1, 2, 3}, args: args{2}, want: Vector32{2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.CMul(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector32.CMul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector32_CPow(t *testing.T) {
	type args struct {
		c float32
	}
	tests := []struct {
		name string
		a    Vector32
		args args
		want Vector32
	}{
		{a: Vector32{1, 2, 3}, args: args{2}, want: Vector32{1, 4, 9}},
		{a: Vector32{1, 2, 3}, args: args{-1}, want: Vector32{1, 1.0 / 2, 1.0 / 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.CPow(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector32.CPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
