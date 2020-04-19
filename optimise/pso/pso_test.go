package pso

import (
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/bbirchnz/gomultivariate/optimise"
	"github.com/bbirchnz/gomultivariate/vector"
)

func TestNewPSO(t *testing.T) {
	type args struct {
		vectorSize       int
		particleCount    int
		inertiaFactor    float32
		globalBestFactor float32
		localBestFactor  float32
		vectorMins       vector.Vector32
		vectorMaxs       vector.Vector32
	}
	tests := []struct {
		name string
		args args
		want *PSO
	}{
		{
			args: args{5, 10, 1, 1, 1, vector.Vector32{0, 0, 0, 0, 0}, vector.Vector32{1, 1, 1, 1, 1}},
			want: &PSO{
				vectorMaxs:               vector.Vector32{1, 1, 1, 1, 1},
				vectorMins:               vector.Vector32{0, 0, 0, 0, 0},
				particleCount:            10,
				vectorSize:               5,
				positions:                make([]vector.Vector32, 10),
				globalLowestCostPosition: vector.Vector32{0, 0, 0, 0, 0},
				globalLowestCost:         math.MaxFloat32,
				lowestCostPositions:      make([]vector.Vector32, 10),
				lowestCosts:              []float32{math.MaxFloat32, math.MaxFloat32, math.MaxFloat32, math.MaxFloat32, math.MaxFloat32, math.MaxFloat32, math.MaxFloat32, math.MaxFloat32, math.MaxFloat32, math.MaxFloat32},
				globalBestFactor:         1,
				inertiaFactor:            1,
				localBestFactor:          1,
				velocities:               make([]vector.Vector32, 10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPSO(tt.args.vectorSize, tt.args.particleCount, tt.args.inertiaFactor, tt.args.globalBestFactor, tt.args.localBestFactor, tt.args.vectorMins, tt.args.vectorMaxs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPSO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPSO_Run(t *testing.T) {
	type args struct {
		costFunction  optimise.CostFunction
		maxIterations int
		targetCost    float32
	}
	tests := []struct {
		name           string
		pso            *PSO
		args           args
		wantBestMax    vector.Vector32
		wantBestMin    vector.Vector32
		wantLowestCost float32
	}{
		{
			name: "Solve X^2 + Y^2",
			pso:  NewPSO(2, 150, 0.4, 2, 1.5, vector.Vector32{-1, -1}, vector.Vector32{1, 1}),
			args: args{
				costFunction: func(v vector.Vector32) float32 {
					return float32(math.Pow(float64(v[0]), 2) + math.Pow(float64(v[1]), 2))
				},
				maxIterations: 100,
				targetCost:    0.1,
			},
			wantBestMax:    vector.Vector32{0.1, 0.1},
			wantBestMin:    vector.Vector32{-0.1, -0.1},
			wantLowestCost: 0.01,
		},
	}
	rand.Seed(1)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBest, gotLowestCost := tt.pso.Run(tt.args.costFunction, tt.args.maxIterations, tt.args.targetCost)
			for i, value := range gotBest {
				if value > tt.wantBestMax[i] || value < tt.wantBestMin[i] {
					t.Errorf("PSO.Run() element %d gotBest = %v, want < %v, > %v", i, value, tt.wantBestMax[i], tt.wantBestMin[i])
				}
			}

			if gotLowestCost > tt.wantLowestCost {
				t.Errorf("PSO.Run() gotLowestCost = %v, want %v", gotLowestCost, tt.wantLowestCost)
			}
		})
	}
}
