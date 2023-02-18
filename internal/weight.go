package internal

import "fmt"

type Pound float64
type Kilogram float64

const (
	KtoPfactor float64 = 2.2046
)

// returns pounds as a string
func (p Pound) String() string { return fmt.Sprintf("%glb", p) }

// returns kilograms as a string
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
