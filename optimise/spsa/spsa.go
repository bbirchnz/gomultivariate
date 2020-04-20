// Package spsa is an implementation of simultaneous perturbation stochastic approximation
// https://en.wikipedia.org/wiki/Simultaneous_perturbation_stochastic_approximation
package spsa

import (
	"math"

	"github.com/bbirchnz/gomultivariate/optimise"
	"github.com/bbirchnz/gomultivariate/vector"
)

// Optimiser defines a SPSA optimiser
type Optimiser struct {
	aInit      float64
	cInit      float64
	alpha      float64
	gamma      float64
	vectorSize int
	vectorMins vector.Vector32
	vectorMaxs vector.Vector32
}

// NewOptimiser initialises and returns a pointer to an Optimiser
func NewOptimiser(
	aInit float64,
	cInit float64,
	vectorSize int,
	alpha float64,
	gamma float64,
	vectorMins vector.Vector32,
	vectorMaxs vector.Vector32,
) *Optimiser {
	return &Optimiser{
		aInit:      aInit,
		cInit:      cInit,
		alpha:      alpha,
		gamma:      gamma,
		vectorSize: vectorSize,
		vectorMins: vectorMins,
		vectorMaxs: vectorMaxs,
	}
}

// Run executes an optimisation pass, iterating until maxIterations or targetCost is achieved
func (o *Optimiser) Run(costFunction optimise.CostFunction, maxIterations int, targetCost float32) (best vector.Vector32, lowestCost float32) {
	// initial random position:
	v := vector.NewRandomVector32(o.vectorSize, o.vectorMins, o.vectorMaxs)
	lowestCost = costFunction(&v)
	A := 10.0
	for i := 0; i < maxIterations; i++ {
		// escape if found target:
		lowestCost = costFunction(&v)
		if lowestCost <= targetCost {
			break
		}

		ai := math.Pow(o.aInit/(float64(i)+1+A), o.alpha)
		ci := math.Pow(o.cInit/(float64(i)+1), o.gamma)
		delta := vector.NewRandomVector32(o.vectorSize, vector.NewVector32(o.vectorSize, -1), vector.NewVector32(o.vectorSize, 1))

		deltaCi := delta.CMul(float32(ci))

		vPlus := v.EAdd(deltaCi)
		vMinus := v.ESub(deltaCi)
		lossPlus := costFunction(&vPlus)
		lossMinus := costFunction(&vMinus)

		ghat := deltaCi.CPow(-1).CMul(float32((lossPlus - lossMinus) / (2 * float32(ci))))
		v = v.ESub(ghat.CMul(float32(ai)))
		// keep positions within constraints:
		v = v.EMin(o.vectorMaxs).EMax(o.vectorMins)
	}
	return v, costFunction(&v)
}
