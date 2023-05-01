package main

import (
	"math"
	"testing"
)

func TestLengthMeasurementTo(t *testing.T) {
    length := NewLengthMeasurement(10, "centimeter")

    if length == nil {
        t.Error("Expected valid length measurement, got nil")
    }

    expected := 0.1
    result := length.to("meter")

    if math.Abs(expected-result) > 0.00001 {
        t.Errorf("Expected %f, but got %f", expected, result)
    }

    result = length.to("inch")

    if math.Abs(expected-result) < 0.00001 {
        t.Errorf("Expected NaN, but got %f", result)
    }
}

func TestWeightMeasurementTo(t *testing.T) {
    weight := NewWeightMeasurement(500, "milligram")

    if weight == nil {
        t.Error("Expected valid weight measurement, got nil")
    }

    expected := 0.5
    result := weight.to("gram")

    if math.Abs(expected-result) > 0.00001 {
        t.Errorf("Expected %f, but got %f", expected, result)
    }

    result = weight.to("pound")

    if math.Abs(expected-result) < 0.00001 {
        t.Errorf("Expected NaN, but got %f", result)
    }
}

func TestInvalidMeasurement(t *testing.T) {
    length := NewLengthMeasurement(-10, "meter")

    if length != nil {
        t.Error("Expected invalid length measurement, but got valid measurement")
    }

    weight := NewWeightMeasurement(1000, "unknown")

    if weight != nil {
        t.Error("Expected invalid weight measurement, but got valid measurement")
    }
}

func TestAreUnitsOfSameType(t *testing.T) {
	// Test length measurements
	length1 := NewLengthMeasurement(10, "centimeter")
	length2 := NewLengthMeasurement(5, "meter")
	if !areUnitsOfSameType(length1, length2) {
		t.Errorf("Expected true but got false")
	}

	length3 := NewLengthMeasurement(10, "centimeter")
	weight1 := NewWeightMeasurement(500, "milligram")
	if areUnitsOfSameType(length3, weight1) {
		t.Errorf("Expected false but got true")
	}

	// Test weight measurements
	weight2 := NewWeightMeasurement(100, "gram")
	weight3 := NewWeightMeasurement(50, "kilogram")
	if !areUnitsOfSameType(weight2, weight3) {
		t.Errorf("Expected true but got false")
	}

	weight4 := NewWeightMeasurement(100, "gram")
	length4 := NewLengthMeasurement(5, "meter")
	if areUnitsOfSameType(weight4, length4) {
		t.Errorf("Expected false but got true")
	}
}

