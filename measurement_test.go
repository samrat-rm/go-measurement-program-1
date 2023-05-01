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
	result := length.convertUnit("meter")

	if math.Abs(expected-result) > 0.00001 {
		t.Errorf("Expected %f, but got %f", expected, result)
	}

	result = length.convertUnit("inch")

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
	result := weight.convertUnit("gram")

	if math.Abs(expected-result) > 0.00001 {
		t.Errorf("Expected %f, but got %f", expected, result)
	}

	result = weight.convertUnit("pound")

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
func TestDiffBetweenMeasurements(t *testing.T) {
	// Test valid case for WeightMeasurement
	weight := NewWeightMeasurement(500, "milligram")
	weight2 := NewWeightMeasurement(500, "gram")
	expectedWeight := NewWeightMeasurement(499500, "milligram")
	resultWeight := diffBetweenMeasurements(weight, weight2)
	if resultWeight == nil || resultWeight.Value() != expectedWeight.Value() || resultWeight.Unit() != expectedWeight.Unit() {
		t.Errorf("Unexpected result for diffBetweenMeasurements with WeightMeasurement inputs. Got: %v, expected: %v", resultWeight, expectedWeight)
	}

	// Test valid case for LengthMeasurement
	length := NewLengthMeasurement(10, "centimeter")
	length2 := NewLengthMeasurement(10, "meter")
	expectedLength := NewLengthMeasurement(990, "centimeter")
	resultLength := diffBetweenMeasurements(length, length2)
	if resultLength == nil || resultLength.Value() != expectedLength.Value() || resultLength.Unit() != expectedLength.Unit() {
		t.Errorf("Unexpected result for diffBetweenMeasurements with LengthMeasurement inputs. Got: %v, expected: %v", resultLength, expectedLength)
	}

	// Test invalid case
	invalidMeasurement := NewLengthMeasurement(5, "candela") // "candela" is not a valid unit for length
	resultInvalid := diffBetweenMeasurements(length, invalidMeasurement)
	if resultInvalid != nil {
		t.Errorf("Unexpected result for diffBetweenMeasurements with invalid inputs. Got: %v, expected: nil", resultInvalid)
	}
}
func TestAddMeasurements(t *testing.T) {
	// Test valid case for WeightMeasurement
	weight := NewWeightMeasurement(500, "milligram")
	weight2 := NewWeightMeasurement(500, "gram")
	expectedWeight := NewWeightMeasurement(500500, "milligram")
	resultWeight := addMeasurements(weight, weight2)
	if resultWeight == nil || resultWeight.Value() != expectedWeight.Value() || resultWeight.Unit() != expectedWeight.Unit() {
		t.Errorf("Unexpected result for addMeasurements with WeightMeasurement inputs. Got: %v, expected: %v", resultWeight, expectedWeight)
	}

	// Test valid case for LengthMeasurement
	length := NewLengthMeasurement(10, "centimeter")
	length2 := NewLengthMeasurement(10, "meter")
	expectedLength := NewLengthMeasurement(1010, "centimeter")
	resultLength := addMeasurements(length, length2)
	if resultLength == nil || resultLength.Value() != expectedLength.Value() || resultLength.Unit() != expectedLength.Unit() {
		t.Errorf("Unexpected result for addMeasurements with LengthMeasurement inputs. Got: %v, expected: %v", resultLength, expectedLength)
	}

	// Test invalid case
	invalidMeasurement := NewLengthMeasurement(5, "candela") // "candela" is not a valid unit for length
	resultInvalid := addMeasurements(length, invalidMeasurement)
	if resultInvalid != nil {
		t.Errorf("Unexpected result for addMeasurements with invalid inputs. Got: %v, expected: nil", resultInvalid)
	}
}
