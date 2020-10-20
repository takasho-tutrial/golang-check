package complexnumber

import "testing"

func TestNewComplexNumber(t *testing.T) {
	cNumInstance := &complexNumber{
		real:      1,
		imaginary: 2,
	}

	_, ok := interface{}(cNumInstance).(ComplexNumber)

	if ok == false {
		t.Errorf("complexNumber, which is a structure, did't work")
	}
}

func TestGetReal(t *testing.T) {
	const (
		real      float64 = 1
		imaginary float64 = 2
	)
	cNum, _ := NewComplexNumber(real, imaginary)

	res := cNum.GetReal()

	if res != real {
		t.Errorf("Not matched real number. exp: %v, result: %v", real, res)
	}
}
