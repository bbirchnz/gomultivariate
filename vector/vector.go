package vector

import "fmt"

// Vector32 represents a float32 vector
type Vector32 []float32

// NewVector32 returns a new vector with size {size} and all elements set to {value}
func NewVector32(size int, value float32) Vector32 {
	v := make([]float32, size)
	for i := range v {
		v[i] = value
	}
	return v
}

// Add does element wise addition of a + b
func (a Vector32) Add(b Vector32) Vector32 {
	for i, value := range a {
		a[i] = value + b[i]
	}
	return a
}

// Sub does element wise subtraction of a - b
func (a Vector32) Sub(b Vector32) Vector32 {
	for i, value := range a {
		a[i] = value - b[i]
	}
	return a
}

// Mul does element wise multiplication of a * b
func (a Vector32) Mul(b Vector32) Vector32 {
	for i, value := range a {
		a[i] = value * b[i]
	}
	return a
}

// Max returns maximum value in vector a
func (a Vector32) Max() float32 {
	current := a[0]
	for i := range a {
		if a[i] > current {
			current = a[i]
		}
	}
	return current
}

// Min returns minimum value in vector a
func (a Vector32) Min() float32 {
	current := a[0]
	for i := range a {
		if a[i] < current {
			current = a[i]
		}
	}
	return current
}

// String returns vector as a comma seperated string with numbers to 3 dp
func (a Vector32) String() string {
	out := ""
	for _, v := range a {
		out = out + fmt.Sprintf("%.3f,", v)
	}
	return out
}
