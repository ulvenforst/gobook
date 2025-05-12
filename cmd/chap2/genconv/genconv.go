// Package genconv performs Celsius and Fahrenheit, Feet and Meter, and Pound and Kilogram conversions.
package genconv

import (
	"fmt"
)

const (
	unitMeterInFeet     = 3.28084
	unitKilogramInPound = 2.204623
)

// Temperature types
type Celsius float64
type Fahrenheit float64

// Length types
type Feet float64
type Meter float64

// Weight types
type Pound float64
type Kilogram float64

// String methods for temperature types
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// String methods for length types
func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }

// String methods for weight types
func (p Pound) String() string    { return fmt.Sprintf("%g lb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }

// Conversion functions for temperature

// CToF converts Celsius to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts Fahrenheit to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FeetToMeter converts Feet to Meter.
func FeetToMeter(f Feet) Meter { return Meter(f / unitMeterInFeet) }

// MeterToFeet converts Meter to Feet.
func MeterToFeet(m Meter) Feet { return Feet(m * unitMeterInFeet) }

// PoundToKilogram converts Pound to Kilogram.
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p / unitKilogramInPound) }

// KilogramToPound converts Kilogram to Pound.
func KilogramToPound(k Kilogram) Pound { return Pound(k * unitKilogramInPound) }
