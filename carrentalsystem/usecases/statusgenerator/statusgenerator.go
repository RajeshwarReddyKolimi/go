package statusgenerator

import "math/rand"

type Status struct {
}
type StatusGenerator interface {
	Generate() bool
}

func NewStatusGenerator() StatusGenerator {
	return &Status{}
}

func (rs *Status) Generate() bool {
	return rand.Int()%2 == 0
}
