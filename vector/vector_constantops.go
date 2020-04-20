package vector

import "math"

// CAdd adds constant to each element and returns new vector
func (a Vector32) CAdd(c float32) Vector32 {
	out := make(Vector32, len(a))
	for i, value := range a {
		out[i] = value + c
	}
	return out
}

// CSub subtracts constant from each element and returns new vector
func (a Vector32) CSub(c float32) Vector32 {
	out := make(Vector32, len(a))
	for i, value := range a {
		out[i] = value - c
	}
	return out
}

// CMul multiplies by constant for each element and returns new vector
func (a Vector32) CMul(c float32) Vector32 {
	out := make(Vector32, len(a))
	for i, value := range a {
		out[i] = value * c
	}
	return out
}

// CPow take each element to power of c and returns new vector
func (a Vector32) CPow(c float32) Vector32 {
	out := make(Vector32, len(a))
	for i, value := range a {
		out[i] = float32(math.Pow(float64(value), float64(c)))
	}
	return out
}
