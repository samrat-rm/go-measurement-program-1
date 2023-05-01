package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

type LengthMeasurement struct {
	value float64
	unit  string
}

func (m LengthMeasurement) Value() float64 {
	return m.value
}

func (m LengthMeasurement) Unit() string {
	return m.unit
}

type WeightMeasurement struct {
	value float64
	unit  string
}

func (m WeightMeasurement) Value() float64 {
	return m.value
}

func (m WeightMeasurement) Unit() string {
	return m.unit
}

type Measurable interface {
	convertUnit(unit string) float64
	Value() float64
	Unit() string
}

var lengthMeasurements = map[string]float64{
	"millimeter": 1,
	"centimeter": 0.1,
	"decimeter":  0.01,
	"inch":       0.03937,
	"meter":      0.001,
}

var weightMeasurements = map[string]float64{
	"milligram": 10000,
	"gram":      10,
	"kilogram":  0.01,
	"pound":     0.022,
	"ounce":     0.353,
}

func NewLengthMeasurement(value float64, unit string) Measurable {
	if value <= 0 {
		return nil
	}
	if _, ok := lengthMeasurements[unit]; ok {
		return LengthMeasurement{value, unit}
	}
	return nil
}

func NewWeightMeasurement(value float64, unit string) Measurable {
	if value <= 0 {
		return nil
	}
	if _, ok := weightMeasurements[unit]; ok {
		return WeightMeasurement{value, unit}
	}
	return nil
}

func (m LengthMeasurement) convertUnit(unit string) float64 {
	if m.unit == unit {
		return m.value
	}

	if factor, ok := lengthMeasurements[m.unit]; ok {
		if factor2, ok := lengthMeasurements[unit]; ok {
			ratio := factor / factor2
			result := m.value / ratio

			// Format the float to 2 decimal places
			formatted := fmt.Sprintf("%.2f", result)

			// Parse the formatted string back into a float64 type
			parsed, err := strconv.ParseFloat(formatted, 64)
			if err != nil {
				log.Printf("Error parsing measurement: %v", err)
				return 0
			}

			// Print the result
			return parsed
		}
	}
	return math.NaN()
}

func (m WeightMeasurement) convertUnit(unit string) float64 {
	if m.unit == unit {
		return m.value
	}

	if factor, ok := weightMeasurements[m.unit]; ok {
		if factor2, ok := weightMeasurements[unit]; ok {
			ratio := factor / factor2
			result := m.value / ratio

			// Format the float to 2 decimal places
			formatted := fmt.Sprintf("%.2f", result)

			// Parse the formatted string back into a float64 type
			parsed, err := strconv.ParseFloat(formatted, 64)
			if err != nil {
				log.Printf("Error parsing measurement: %v", err)
				return 0
			}

			// Print the result
			return parsed
		}
	}
	return math.NaN()
}

func areUnitsOfSameType(m1, m2 Measurable) bool {
	_, isLengthMeasurement1 := m1.(LengthMeasurement)
	_, isLengthMeasurement2 := m2.(LengthMeasurement)
	_, isWeightMeasurement1 := m1.(WeightMeasurement)
	_, isWeightMeasurement2 := m2.(WeightMeasurement)

	if isLengthMeasurement1 && isLengthMeasurement2 {
		_, ok := lengthMeasurements[m1.(LengthMeasurement).unit]
		if ok {
			_, ok := lengthMeasurements[m2.(LengthMeasurement).unit]
			return ok
		}
	} else if isWeightMeasurement1 && isWeightMeasurement2 {
		_, ok := weightMeasurements[m1.(WeightMeasurement).unit]
		if ok {
			_, ok := weightMeasurements[m2.(WeightMeasurement).unit]
			return ok
		}
	}
	return false
}

func addMeasurements(m1, m2 Measurable) Measurable {
	if !areUnitsOfSameType(m1, m2) {
		return nil
	}

	switch m1 := m1.(type) {
	case LengthMeasurement:
		m2InM1Unit := NewLengthMeasurement(m2.convertUnit(m1.unit), m1.unit)
		newValue := m1.value + m2InM1Unit.convertUnit(m1.unit)
		return NewLengthMeasurement(truncate(newValue), m1.unit)

	case WeightMeasurement:
		m2InM1Unit := NewWeightMeasurement(m2.convertUnit(m1.unit), m1.unit)
		newValue := m1.value + m2InM1Unit.convertUnit(m1.unit)
		return NewWeightMeasurement(truncate(newValue), m1.unit)

	default:
		return nil
	}
}

func diffBetweenMeasurements(m1, m2 Measurable) Measurable {
	if !areUnitsOfSameType(m1, m2) {
		return nil
	}

	switch m1 := m1.(type) {
	case LengthMeasurement:
		m2InM1Unit := NewLengthMeasurement(m2.convertUnit(m1.unit), m1.unit)
		newValue := math.Abs(m1.value - m2InM1Unit.convertUnit(m1.unit))
		return NewLengthMeasurement(truncate(newValue), m1.unit)

	case WeightMeasurement:
		m2InM1Unit := NewWeightMeasurement(m2.convertUnit(m1.unit), m1.unit)
		newValue := math.Abs(m1.value - m2InM1Unit.convertUnit(m1.unit))
		return NewWeightMeasurement(truncate(newValue), m1.unit)

	default:
		return nil
	}
}

func truncate(value float64) float64 {
	return math.Round(value*100) / 100
}

func main() {
	length := NewLengthMeasurement(10, "centimeter")
	length2 := NewLengthMeasurement(10, "meter")
	if length == nil {
		fmt.Println("Invalid length measurement")
	} else {
		fmt.Println(length.convertUnit("meter"))
	}

	weight := NewWeightMeasurement(500, "milligram")
	weight2 := NewWeightMeasurement(500, "gram")
	if weight == nil {
		fmt.Println("Invalid weight measurement")
	} else {
		fmt.Println(weight.convertUnit("gram"))
	}

	fmt.Println(addMeasurements(weight, weight2))
	fmt.Println(addMeasurements(length, length2))
	fmt.Println(diffBetweenMeasurements(weight, weight2))
	fmt.Println(diffBetweenMeasurements(length, length2))
}
