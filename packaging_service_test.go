package main

import (
	"reflect"
	services "shipment-service/services"
	"testing"
)

func TestCase1(t *testing.T) {
	var expected = map[int]int{250: 1}
	const itemsCount = 1
	var output, err = services.CalculatePackagesFor(itemsCount)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	} else if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output is not equal to expected result")
	}
}

func TestCase2(t *testing.T) {
	var expected = map[int]int{250: 1}
	const itemsCount = 250
	var output, err = services.CalculatePackagesFor(itemsCount)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	} else if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output is not equal to expected result")
	}
}

func TestCase3(t *testing.T) {
	var expected = map[int]int{500: 1}
	const itemsCount = 251
	var output, err = services.CalculatePackagesFor(itemsCount)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	} else if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output is not equal to expected result")
	}
}

func TestCase4(t *testing.T) {
	var expected = map[int]int{500: 1, 250: 1}
	const itemsCount = 501
	var output, err = services.CalculatePackagesFor(itemsCount)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	} else if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output is not equal to expected result")
	}
}

func TestCase5(t *testing.T) {
	var expected = map[int]int{5000: 2, 2000: 1, 250: 1}
	const itemsCount = 12001
	var output, err = services.CalculatePackagesFor(itemsCount)
	if err != nil {
		t.Errorf("Error: %s", err.Error())
	} else if !reflect.DeepEqual(output, expected) {
		t.Errorf("Output is not equal to expected result")
	}
}

func TestCase6(t *testing.T) {
	var expected = map[int]int{5000: 2, 2000: 1, 250: 1}
	const itemsCount = -1
	var output, err = services.CalculatePackagesFor(itemsCount)
	if err == nil {
		t.Errorf("Error should not be nil")
	} else if reflect.DeepEqual(output, expected) {
		t.Errorf("Output is equal to expected result while input value is not correct")
	}
}
