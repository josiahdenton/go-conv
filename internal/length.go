package internal

import "fmt"

const (
	MtoFfactor float64 = 3.281
)

type Meter float64
type Foot float64

// returns a representation of meters as a string
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// returns a representation of feet as a string
func (f Foot) String() string { return fmt.Sprintf("%g", f) }
