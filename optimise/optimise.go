package optimise

import "github.com/bbirchnz/gomultivariate/vector"

type Optimiser interface {
	Run(costFunction CostFunction, maxIterations int) vector.Vector32
}

type CostFunction func(vector.Vector32) float32
