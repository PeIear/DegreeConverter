package main

import (
  "fmt"
)

const (
	str  = "%s\n|  %s        | %s         |\n%s\n%s%s\n"
	form = "|  %5.1f     |  %5.1f     |\n"
)

func main() {
	ColumnA := "°C"
	ColumnB := "°F"
	Table(CelsiusToFahrenheit, FahrenheitToCelsius, ColumnA, ColumnB)
}

type Fn func() string

func Table(a, b Fn, ColumnA, ColumnB string) {

	var feature string
	for i := 0; i < 27; i++ {
		feature += "="
	}

	fmt.Printf(str, feature, ColumnA, ColumnB, feature, a(), feature)
	fmt.Printf(str, feature, ColumnB, ColumnA, feature, b(), feature)
}

func CelsiusToFahrenheit() string {
	var f fahrenheit
	var c celsius = -40
	var result string

	for c = -40; c <= 100; c += 5 {
		f = c.fahrenheit()

		result += fmt.Sprintf(form, c, f)

	}
	return result
}

func FahrenheitToCelsius() string {
	var f fahrenheit = -40
	var c celsius
	var result string

	for f = -40; f <= 100; f += 5 {
		c = f.celsius()

		result += fmt.Sprintf(form, f, c)

	}
  return result
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}


