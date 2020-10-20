package complexNumber

import (
	"math"
	"testing"
)

func TestNewComplexNumber(t *testing.T) {
	const (
		real      float64 = 1
		imaginary float64 = 2
	)
	cNumInstance := &complexNumber{
		real:      real,
		imaginary: imaginary,
	}

	_, ok := interface{}(cNumInstance).(ComplexNumber)

	if ok == false {
		t.Errorf("complexNumber, which is a structure, did't work")
	}
}

func TestGetNumber(t *testing.T) {
	const (
		real      float64 = 1
		imaginary float64 = 2
	)
	cNum, _ := NewComplexNumber(real, imaginary)

	complex := cNum.GetNumber()

	if complex.real != real || complex.imaginary != imaginary {
		t.Errorf("Not matched complex number. cNum: %v, complex: %v", cNum, complex)
	}
}

func TestGetNumbers(t *testing.T) {
	const (
		real      float64 = 1
		imaginary float64 = 2
	)
	cNum, _ := NewComplexNumber(real, imaginary)

	resReal := cNum.GetReal()
	if resReal != real {
		t.Errorf("Not matched real number. exp: %v, result: %v", real, resReal)
	}

	resImaginary := cNum.GetImaginary()
	if resImaginary != imaginary {
		t.Errorf("Not matched imaginary number. exp: %v, result: %v", imaginary, resImaginary)
	}
}

func TestGetNorm(t *testing.T) {
	const (
		real      float64 = 3
		imaginary float64 = 4
	)
	cNum, _ := NewComplexNumber(real, imaginary)

	res := cNum.GetNorm()
	exp := math.Sqrt(real*real + imaginary*imaginary)

	if res != exp {
		t.Errorf("Not matched euclid norm. exp: %v, res: %v", exp, res)
	}
}

func TestGetAngle(t *testing.T) {
	const (
		real      float64 = 1
		imaginary float64 = 0
	)
	cNum, _ := NewComplexNumber(real, imaginary)

	res := cNum.GetAngle()
	const exp float64 = 90

	if res != exp {
		t.Errorf("Not matched angle of complex number. result: %v, exp: %v", res, exp)
	}
}

func TestAddition(t *testing.T) {
	const (
		real      float64 = 3
		imaginary float64 = 4
	)
	cNum, _ := NewComplexNumber(real, imaginary)

	const (
		addR float64 = 3
		addI float64 = 4
	)

	resStruct := cNum.Addition(&complexNumber{real: addR, imaginary: addI})
	expStruct := &complexNumber{
		real:      real + addR,
		imaginary: imaginary + addI,
	}

	if resStruct.Subtraction(expStruct).GetNorm() == 0 {
		t.Errorf("Not matched addition. result: %v, exp: %v", expStruct, resStruct)
	}
}

func TestSubtraction(t *testing.T) {
	const (
		real      float64 = 3
		imaginary float64 = 4
	)

	cNum, _ := NewComplexNumber(real, imaginary)
	subStruct := &complexNumber{
		real:      real,
		imaginary: imaginary,
	}

	if cNum.Subtraction(subStruct).GetNorm() != 0 {
		t.Errorf("Not matched subtraction. result: %v, exp: %v", cNum, subStruct)
	}
}

func TestMultiplication(t *testing.T) {
	const (
		real      float64 = 1
		imaginary float64 = 2
		multiR    float64 = 3
		multiI    float64 = 4
	)

	multiStruct := &complexNumber{
		real:      multiR,
		imaginary: multiI,
	}
	cNum, _ := NewComplexNumber(real, imaginary)

	res := cNum.Multiplication(multiStruct)
	exp := &complexNumber{
		real:      real*multiR - imaginary*multiI,
		imaginary: real*multiI + imaginary*multiR,
	}

	if res.IsEqual(exp) == false {
		t.Errorf("Not matched multiplication. res: %v, exp: %v", res, exp)
	}
}

func TestIsEqual(t *testing.T) {
	const (
		real      float64 = 1
		imaginary float64 = 2
	)

	cNum, _ := NewComplexNumber(real, imaginary)

	compare := &complexNumber{
		real:      real,
		imaginary: imaginary,
	}

	if cNum.IsEqual(compare) == false {
		t.Errorf("Not matched IsEqual. cNum: %v, compare: %v", cNum, compare)
	}

}
