// Package spsa is an implementation of simultaneous perturbation stochastic approximation

// https://en.wikipedia.org/wiki/Simultaneous_perturbation_stochastic_approximation

package spsa

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/bbirchnz/gomultivariate/optimise"
	"github.com/bbirchnz/gomultivariate/vector"
)

func TestNewOptimiser(t *testing.T) {
	type args struct {
		aInit      float64
		cInit      float64
		vectorSize int
		alpha      float64
		gamma      float64
		vectorMins vector.Vector32
		vectorMaxs vector.Vector32
	}
	tests := []struct {
		name string
		args args
		want *Optimiser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOptimiser(tt.args.aInit, tt.args.cInit, tt.args.vectorSize, tt.args.alpha, tt.args.gamma, tt.args.vectorMins, tt.args.vectorMaxs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOptimiser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptimiser_Run(t *testing.T) {
	type args struct {
		costFunction  optimise.CostFunction
		maxIterations int
		targetCost    float32
	}
	tests := []struct {
		name           string
		o              *Optimiser
		args           args
		wantBestMax    vector.Vector32
		wantBestMin    vector.Vector32
		wantLowestCost float32
	}{
		{
			name: "Solve X^2 + Y^2",
			o:    NewOptimiser(.75, .75, 2, .6, .4, vector.Vector32{-2, -2}, vector.Vector32{2, 2}),
			args: args{
				costFunction: func(v *vector.Vector32) float32 {
					return float32(math.Pow(float64((*v)[0]), 2) + math.Pow(float64((*v)[1]), 2))
				},
				maxIterations: 100,
				targetCost:    0.01,
			},
			wantBestMax:    vector.Vector32{0.1, 0.1},
			wantBestMin:    vector.Vector32{-0.1, -0.1},
			wantLowestCost: 0.01,
		},
	}
	rand.Seed(1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBest, gotLowestCost := tt.o.Run(tt.args.costFunction, tt.args.maxIterations, tt.args.targetCost)
			for i, value := range gotBest {
				if value > tt.wantBestMax[i] || value < tt.wantBestMin[i] {
					t.Errorf("O.Run() element %d gotBest = %v, want < %v, > %v", i, value, tt.wantBestMax[i], tt.wantBestMin[i])
				}
			}

			if gotLowestCost > tt.wantLowestCost {
				t.Errorf("O.Run() gotLowestCost = %v, want %v", gotLowestCost, tt.wantLowestCost)
			}
		})
	}
}
