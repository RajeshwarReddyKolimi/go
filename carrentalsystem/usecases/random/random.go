package random

import "math/rand"

type RandomStatus struct {
}

type RandomStatusGenerator interface {
	Generate() bool
}

func NewRandomStatusGenerator() RandomStatusGenerator {
	return &RandomStatus{}
}

func (rs *RandomStatus) Generate() bool {
	return rand.Int()%2 == 0
}
