package internal

// converts celsius to fahrentheit
func (c Celsius) CelsiusToFahrenheit() Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// converts fahrentheit to celsius
func (f Fahrenheit) FahrenheitToCelsius() Celsius { return Celsius((f - 32) * 5 / 9) }

// converts meters to feet
func (m Meter) MeterToFoot() Foot { return Foot(m) * Foot(MtoFfactor) }

// conversts feet to meters
func (f Foot) FootToMeter() Meter { return Meter(f) / Meter(MtoFfactor) }

// converts kilograms to pounds
func (k Kilogram) KilogramToPound() Pound { return Pound(k) * Pound(KtoPfactor) }

// converts pounds to kilograms
func (p Pound) PoundToKilogram() Kilogram { return Kilogram(p) / Kilogram(KtoPfactor) }
