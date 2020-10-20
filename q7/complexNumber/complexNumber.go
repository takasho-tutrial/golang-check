package complexnumber

import (
	"math"
)

type (
	// ComplexNumber is complex number
	ComplexNumber interface {
		GetReal() float64
		GetImaginary() float64
		GetNorm() float64
		Addition(real, imaginary float64) ComplexNumber
		GetAngle() float64
		Multiplication(c, d float64) ComplexNumber
	}
	complexNumber struct {
		real      float64
		imaginary float64
	}
)

// NewComplexNumber is a function return ComplexNumber interface
func NewComplexNumber(real, imaginary float64) (ComplexNumber, error) {
	return &complexNumber{
		real:      real,
		imaginary: imaginary,
	}, nil
}

func (cNum *complexNumber) GetReal() float64 {
	return cNum.real
}

func (cNum *complexNumber) GetImaginary() float64 {
	return cNum.imaginary
}

func (cNum *complexNumber) GetNorm() float64 {
	var (
		real      float64 = cNum.real
		imaginary float64 = cNum.imaginary
	)
	return math.Sqrt(real*real + imaginary*imaginary)
}

func (cNum *complexNumber) Addition(real, imaginary float64) ComplexNumber {
	return &complexNumber{
		real:      cNum.real + real,
		imaginary: cNum.imaginary + imaginary,
	}
}

func (cNum *complexNumber) GetAngle() float64 {
	var (
		r       float64 = cNum.real
		i       float64 = cNum.imaginary
		isPosiR bool    = r >= 0
		isPosiI bool    = i >= 0
	)
	angleRad := math.Atan(i / r)
	angleAns := angleRad / math.Pi * 180

	if isPosiR && isPosiI {
		return angleAns
	} else if !isPosiR && isPosiI {
		return -angleAns + 90
	} else if !isPosiR && !isPosiI {
		return angleAns + 180
	} else {
		return -angleAns + 270
	}
}

func (cNum *complexNumber) Multiplication(c, d float64) ComplexNumber {
	var (
		a float64 = cNum.real
		b float64 = cNum.imaginary
	)
	return &complexNumber{
		real:      a*c - b*d,
		imaginary: a*d + b*c,
	}
}
