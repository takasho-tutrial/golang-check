package number

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

// Execute this function only loading
func init() {
	fmt.Println("number!!")
}

var (
	// Logger is what we call Logger, actually.
	logging = log.New(os.Stdout,
		"Number: ",
		log.Ldate|log.Lmicroseconds|log.LUTC|log.Lshortfile,
	)
	logMu sync.Mutex
)

// SetLogger sets logger
func SetLogger(l *log.Logger) {
	logMu.Lock()
	logging = l
	logMu.Unlock()
}

func logf(msg string, v ...interface{}) {
	logMu.Lock()
	logging.Printf(msg, v...)
	logMu.Unlock()
}

type (
	// Number is what we call number, surely, indeed.
	Number interface {
		Addition(float64) float64
		Division(float64) (float64, error)
		GetValue() float64
		SetValue(float64) error
	}
	number struct {
		val float64
	}
)

// NewNumber is just for creating new number, probably.
func NewNumber(arg float64) (Number, error) {
	if arg == 0 {
		logf("Invalid number: 0")
		return nil, errors.New("Invalid number: 0")
	}
	logf("Number created")
	return &number{
		val: arg,
	}, nil
}

func (n *number) Addition(arg float64) float64 {
	return arg + n.val
}

func (n *number) Division(arg float64) (float64, error) {
	if arg == 0 {
		return 0, errors.New("Invalid argument: 0")
	}
	return n.val / arg, nil
}

func (n *number) GetValue() float64 {
	return n.val
}

func (n *number) SetValue(arg float64) error {
	if arg == 0 {
		return errors.New("Invalid argument: 0")
	}
	n.val = arg
	return nil
}
