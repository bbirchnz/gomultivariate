package pso

import (
	"math"
	"math/rand"

	"github.com/bbirchnz/gomultivariate/optimise"
	"github.com/bbirchnz/gomultivariate/vector"
)

// PSO defines a Particle Swarm Optimiser
type PSO struct {
	// minimum allowed value for each element
	vectorMins vector.Vector32
	// maximum allowed value for each element
	vectorMaxs vector.Vector32
	// cost function to apply to candidate solutions
	costFunction optimise.CostFunction
	// number of particles in swarm
	particleCount int
	// current positions of each particle
	positions []vector.Vector32
	// current velocities of each particle
	velocities []vector.Vector32
	// lowest cost achieved by each particle
	lowestCosts []float32
	// position of lowest cost achieved by each particle
	lowestCostPositions []vector.Vector32
	// lowest cost seen in swarm
	globalLowestCost float32
	// position of lowest cost achieved by swarm
	globalLowestCostPosition vector.Vector32
	// size of vector
	vectorSize       int
	inertiaFactor    float32
	globalBestFactor float32
	localBestFactor  float32
}

// NewPSO initialises and returns a PSO
func NewPSO(
	vectorSize int,
	particleCount int,
	inertiaFactor float32,
	globalBestFactor float32,
	localBestFactor float32,
	vectorMins vector.Vector32,
	vectorMaxs vector.Vector32) *PSO {

	return &PSO{
		positions:                make([]vector.Vector32, particleCount),
		velocities:               make([]vector.Vector32, particleCount),
		lowestCostPositions:      make([]vector.Vector32, particleCount),
		vectorMaxs:               vectorMaxs,
		vectorMins:               vectorMins,
		particleCount:            particleCount,
		globalLowestCost:         math.MaxFloat32,
		globalLowestCostPosition: vector.NewVector32(vectorSize, 0),
		inertiaFactor:            inertiaFactor,
		localBestFactor:          localBestFactor,
		globalBestFactor:         globalBestFactor,
		lowestCosts:              vector.NewVector32(particleCount, math.MaxFloat32),
		vectorSize:               vectorSize,
	}
}

// Run executes an optimisation pass, iterating until maxIterations or targetCost is achieved
func (pso *PSO) Run(costFunction optimise.CostFunction, maxIterations int, targetCost float32) (best vector.Vector32, lowestCost float32) {

	pso.initialiseParticles()

	// setup channels for parallel cost calculating (create once per run)
	chanCostJobs := make(chan int, pso.particleCount)

	// calculate initial costs and global bests
	pso.calculateCosts(&costFunction, &chanCostJobs)

	// start iteratively improving them:
	for i := 0; i < maxIterations; i++ { // swarm update iteration
		if pso.globalLowestCost < targetCost {
			break
		}
		for p := 0; p < pso.particleCount; p++ { // particle
			for e := 0; e < pso.vectorSize; e++ { // element
				// if on the limits, reverse velocity
				if pso.positions[p][e] == pso.vectorMaxs[e] || pso.positions[p][e] == pso.vectorMins[e] {
					pso.velocities[p][e] = -pso.velocities[p][e]
				}

				// update velocity:
				pso.velocities[p][e] = pso.inertiaFactor*pso.velocities[p][e] +
					rand.Float32()*pso.localBestFactor*(pso.lowestCostPositions[p][e]-pso.positions[p][e]) +
					rand.Float32()*pso.globalBestFactor*(pso.globalLowestCostPosition[e]-pso.positions[p][e])
				// update position:
				pso.positions[p][e] = pso.positions[p][e] + pso.velocities[p][e]
				// cap on min/max
				if pso.positions[p][e] > pso.vectorMaxs[e] {
					pso.positions[p][e] = pso.vectorMaxs[e]
				}
				if pso.positions[p][e] < pso.vectorMins[e] {
					pso.positions[p][e] = pso.vectorMins[e]
				}
			}
		}
		// now calculate all particle costs and update bests:
		pso.calculateCosts(&costFunction, &chanCostJobs)
	}
	return pso.globalLowestCostPosition, pso.globalLowestCost
}

// initialiseParticles generates initial particle positions and velocities
func (pso *PSO) initialiseParticles() {
	for p := 0; p < pso.particleCount; p++ {
		position := vector.NewVector32(pso.vectorSize, 0)
		velocity := vector.NewVector32(pso.vectorSize, 0)

		for e := 0; e < pso.vectorSize; e++ {
			// random position between vectorMins and vectorMaxs
			position[e] = pso.vectorMins[e] + rand.Float32()*(pso.vectorMaxs[e]-pso.vectorMins[e])
			// random velocity where each element is random 0 -> +/- 10% of element range
			velocity[e] = (rand.Float32() - 0.5) * 2 * (pso.vectorMaxs[e] - pso.vectorMins[e]) / 10
		}
		pso.positions[p] = position
		pso.lowestCostPositions[p] = position
		pso.velocities[p] = velocity
	}
}

// calculateCosts calculates costs for particles in parallel based on current position
func (pso *PSO) calculateCosts(costFunction *optimise.CostFunction, costChannel *chan int) {
	// done := make(chan int, pso.particleCount)
	for p := 0; p < pso.particleCount; p++ {
		go func(p int, pso *PSO) {
			cost := (*costFunction)(&pso.positions[p])
			if cost < pso.lowestCosts[p] {
				pso.lowestCosts[p] = cost
				pso.lowestCostPositions[p] = pso.positions[p]
			}
			*costChannel <- 1
		}(p, pso)
	}
	// wait for all:
	for p := 0; p < pso.particleCount; p++ {
		<-*costChannel
	}
	// update global cost and best position
	for p := 0; p < pso.particleCount; p++ {
		if pso.lowestCosts[p] < pso.globalLowestCost {
			pso.globalLowestCost = pso.lowestCosts[p]
			pso.globalLowestCostPosition = pso.lowestCostPositions[p]
		}
	}
}
