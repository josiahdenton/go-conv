package internal

import "fmt"

type Celsius float64
type Fahrenheit float64

// returns a representation of celsius in as a string
func (c Celsius) String() string { return fmt.Sprintf("%gC", c) }

// returns a representation of fahrenheit as a string
func (f Fahrenheit) String() string { return fmt.Sprintf("%gF", f) }
