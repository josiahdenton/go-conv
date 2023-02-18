package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	cv "github.com/josiahdenton/go-conv/internal"
)

var (
	conversionMode      = flag.String("mode", "", "sets the conversion mode, expected values are: t (temperature), l (length), and w (weight)")
	from                = flag.String("from", "", "sets what we are converting from (i.e. m,ft,kg,lb,c,f)")
	to                  = flag.String("to", "", "sets what we are converting to (i.e. m,ft,kg,lb,c,f)")
	conversionModeShort = flag.String("m", "", "short form for mode")
	fromShort           = flag.String("f", "", "short for from")
	toShort             = flag.String("t", "", "short for to")
)

func main() {
	flag.Parse()
	if *conversionMode == "" && *conversionModeShort == "" {
		log.Fatalln("no conversion mode specified")
	}

	if *from == "" && *fromShort == "" {
		log.Fatalln("no metric for from given")
	}

	if *to == "" && *toShort == "" {
		log.Fatalln("no metric for from given")
	}

	switch {
	case (*conversionMode == "t" || *conversionModeShort == "t") && (*from == "f" || *fromShort == "f") && (*to == "c" || *toShort == "c"):
		fmt.Printf("%g\n", cv.Fahrenheit(value()).FahrenheitToCelsius())
	case (*conversionMode == "t" || *conversionModeShort == "t") && (*from == "c" || *fromShort == "c") && (*to == "f" || *toShort == "f"):
		fmt.Printf("%g\n", cv.Celsius(value()).CelsiusToFahrenheit())
	case (*conversionMode == "l" || *conversionModeShort == "l") && (*from == "ft" || *fromShort == "ft") && (*to == "m" || *toShort == "m"):
		fmt.Printf("%g\n", cv.Foot(value()).FootToMeter())
	case (*conversionMode == "l" || *conversionModeShort == "l") && (*from == "m" || *fromShort == "m") && (*to == "ft" || *toShort == "ft"):
		fmt.Printf("%g\n", cv.Meter(value()).MeterToFoot())
	case (*conversionMode == "w" || *conversionModeShort == "w") && (*from == "kg" || *fromShort == "kg") && (*to == "lb" || *toShort == "lb"):
		fmt.Printf("%g\n", cv.Kilogram(value()).KilogramToPound())
	case (*conversionMode == "w" || *conversionModeShort == "w") && (*from == "lb" || *fromShort == "lb") && (*to == "kg" || *toShort == "kg"):
		fmt.Printf("%g\n", cv.Kilogram(value()).KilogramToPound())
	default:
		log.Fatalln("unsupported conversion!")
	}
}

func value() float64 {
	last := os.Args[len(os.Args)-1]
	value, err := strconv.ParseFloat(last, 64)
	if err != nil {
		log.Fatalln("not a valid argument")
	}
	return value
}
