package vector

import (
	"fmt"
	"math/rand"
)

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

// NewRandomVector32 returns a new vector with size {size} where each element is random
// between vectorMins and vectorMaxs
func NewRandomVector32(size int, vectorMins Vector32, vectorMaxs Vector32) Vector32 {
	v := make([]float32, size)
	for i := range v {
		v[i] = vectorMins[i] + rand.Float32()*(vectorMaxs[i]-vectorMins[i])
	}
	return v
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
