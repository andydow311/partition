package main

import (
	"testing"
)

func TestGetAllPartitionsReturnsCorrectNumberOfSpaces(t *testing.T) {
	items := []string{"A", "B", "C", "D", "E", "F"}
	expected := 203 //Bell No.
	actual := numberOfStructures(getAllPartitions(items), items)
	if expected != actual {
		t.Error("Expected ", expected, "agents, but got ", actual)
	}
}
