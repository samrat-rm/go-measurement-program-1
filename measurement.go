package main

import (
	"fmt"
	"math"
)

type LengthMeasurement struct {
    value float64
    unit  string
}

type WeightMeasurement struct {
    value float64
    unit  string
}

type Measurable interface {
    to(unit string) float64
}

var lengthMeasurements = map[string]float64{
    "millimeter": 1,
    "centimeter": 0.1,
    "decimeter": 0.01,
    "inch": 0.03937,
    "meter": 0.001,
}

var weightMeasurements = map[string]float64{
    "milligram": 10000,
    "gram": 10,
    "kilogram": 0.01,
    "pound": 0.022,
    "ounce": 0.353,
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

func (m LengthMeasurement) to(unit string) float64 {
    if m.unit == unit {
        return m.value
    }

    if factor, ok := lengthMeasurements[m.unit]; ok {
        if factor2, ok := lengthMeasurements[unit]; ok {
            ratio := factor / factor2
            return m.value / ratio
        }
    }
    return math.NaN()
}

func (m WeightMeasurement) to(unit string) float64 {
    if m.unit == unit {
        return m.value
    }

    if factor, ok := weightMeasurements[m.unit]; ok {
        if factor2, ok := weightMeasurements[unit]; ok {
			ratio := factor / factor2
            return m.value / ratio
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





func main() {
    length := NewLengthMeasurement(10, "centimeter")
    length2 := NewLengthMeasurement(10, "meter")
    if length == nil {
        fmt.Println("Invalid length measurement")
    } else {
        fmt.Println(length.to("meter"))
    }

    weight := NewWeightMeasurement(500, "milligram")
    weight2 := NewWeightMeasurement(500, "kilogram")
    if weight == nil {
        fmt.Println("Invalid weight measurement")
    } else {
        fmt.Println(weight.to("gram"))
    }
	fmt.Println(areUnitsOfSameType(weight , weight2))
	fmt.Println(areUnitsOfSameType(length , length2))
}
