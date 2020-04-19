package pso

import (
	"github.com/bbirchnz/gomultivariate/optimise"
	"github.com/bbirchnz/gomultivariate/vector"
)

// PSO defines a Particle Swarm Optimiser
type PSO struct {
	vectorMins               vector.Vector32
	vectorMaxs               vector.Vector32
	costFunction             optimise.CostFunction
	particleCount            int
	positionArray            []vector.Vector32
	velocityArray            []vector.Vector32
	lowestCostArray          []float32
	lowestCostPositionArray  []vector.Vector32
	globalLowestCost         float32
	globalLowestCostPosition vector.Vector32
	inertiaFactor            float32
	globalBestFactor         float32
	localBestFactor          float32
	iterationCount           int
}

func NewPSO(
	vectorSize int,
	particleCount int,
	inertiaFactor float32,
	globalBestFactor float32,
	localBestFactor float32,
	vectorMins vector.Vector32,
	vectorMaxs vector.Vector32,
	costFunction optimise.CostFunction) *PSO {

	pso := PSO{
		positionArray:           make([]vector.Vector32, particleCount),
		velocityArray:           make([]vector.Vector32, particleCount),
		lowestCostPositionArray: make([]vector.Vector32, particleCount),
		vectorMaxs:              vectorMaxs,
		vectorMins:              vectorMins,
		costFunction:            costFunction,
	}

	return &pso
}
