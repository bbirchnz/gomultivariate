package optimise

import "github.com/bbirchnz/gomultivariate/vector"

//
type Optimiser interface {
	Run(costFunction CostFunction, maxIterations int) (best vector.Vector32, lowestCost float32)
}

// CostFunction is a function that takes a possible solution and returns its "cost" as a float32
type CostFunction func(vector.Vector32) float32
