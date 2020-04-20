package vector

// EAdd does element wise addition of a + b
func (a Vector32) EAdd(b Vector32) Vector32 {
	out := make(Vector32, len(a))
	for i, value := range a {
		out[i] = value + b[i]
	}
	return out
}

// ESub does element wise subtraction of a - b
func (a Vector32) ESub(b Vector32) Vector32 {
	out := make(Vector32, len(a))
	for i, value := range a {
		out[i] = value - b[i]
	}
	return out
}

// EMul does element wise multiplication of a * b
func (a Vector32) EMul(b Vector32) Vector32 {
	out := make(Vector32, len(a))
	for i, value := range a {
		out[i] = value * b[i]
	}
	return out
}

// EMax returns a new vector where each element e[i] is max(a[i],b[i])
func (a Vector32) EMax(b Vector32) Vector32 {
	out := make(Vector32, len(a))
	for i := range a {
		if b[i] > a[i] {
			out[i] = b[i]
		} else {
			out[i] = a[i]
		}
	}
	return out
}

// EMin returns a new vector where each element e[i] is min(a[i],b[i])
func (a Vector32) EMin(b Vector32) Vector32 {
	out := make(Vector32, len(a))
	for i := range a {
		if b[i] < a[i] {
			out[i] = b[i]
		} else {
			out[i] = a[i]
		}
	}
	return out
}
