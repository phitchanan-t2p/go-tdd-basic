package captcha

import (
	"fmt"
	"math/rand"
)

// Randomizer interface defines a method for generating random numbers
type Randomizer interface {
	Random() int
	Captcha() string
}

// RandGenerator implements the Randomizer interface
type RandGenerator struct {
	Randomizer
}

// Random generates a random number between 100000 and 999999
func (r *RandGenerator) Random() int {
	return rand.Intn(999999-100000+1) + 100000
}

func NewRandGenerator(randomizer Randomizer) *RandGenerator {
	return &RandGenerator{
		Randomizer: randomizer,
	}
}

// Captcha returns string
func (c *RandGenerator) Captcha() string {
	randomNumber := c.Randomizer.Random()

	var captchaPrefix string
	switch {
	case randomNumber >= 100000 && randomNumber <= 299999:
		captchaPrefix = "T2P"
	case randomNumber >= 300000 && randomNumber <= 599999:
		captchaPrefix = "DPV"
	case randomNumber >= 600000 && randomNumber <= 999999:
		captchaPrefix = "LNV"
	}

	return fmt.Sprintf("%s%d", captchaPrefix, randomNumber)
}
