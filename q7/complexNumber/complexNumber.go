package complexNumber

import (
	"math"
)

type (
	// ComplexNumber is complex number
	ComplexNumber interface {
		GetReal() float64
		GetImaginary() float64
		GetNorm() float64
		GetAngle() float64

		Addition(arg *complexNumber) ComplexNumber
		Subtraction(arg *complexNumber) ComplexNumber
		Multiplication(arg *complexNumber) ComplexNumber

		IsEqual(arg *complexNumber) bool
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

func (cNum *complexNumber) Addition(arg *complexNumber) ComplexNumber {
	return &complexNumber{
		real:      cNum.real + arg.real,
		imaginary: cNum.imaginary + arg.imaginary,
	}
}

func (cNum *complexNumber) Subtraction(arg *complexNumber) ComplexNumber {
	return &complexNumber{
		real:      cNum.real - arg.real,
		imaginary: cNum.imaginary - arg.imaginary,
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
		return 180 - angleAns
	} else if !isPosiR && !isPosiI {
		return 270 - angleAns
	} else {
		return 360 - angleAns
	}
}

func (cNum *complexNumber) Multiplication(arg *complexNumber) ComplexNumber {
	var (
		a float64 = cNum.real
		b float64 = cNum.imaginary
		c float64 = arg.real
		d float64 = arg.imaginary
	)
	return &complexNumber{
		real:      a*c - b*d,
		imaginary: a*d + b*c,
	}
}

func (cNum *complexNumber) IsEqual(arg *complexNumber) bool {
	return arg.Subtraction(cNum).GetNorm() == 0
}
