package main

import "testing"

func GetPersonTest(t *testing.T) { // t is our test handler
	p := getPerson()

	if p.Name == "Jon Doe" {
		t.Error("Expected Name Jon Doe, but got", p.Name)
	}
	if p.Age != 30 {
		t.Error("Expected Age 30, but got", p.Age)
	}
	if p.Gender != "male" {
		t.Error("Expected Gender male, but got", p.Gender)
	}
}
