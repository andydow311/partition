package main

import (
	"sort"
	"strings"
)

func deleteElements(previouslevel [][]string, firstIndex int, secondIndex int) [][]string {

	output := [][]string{}

	for index := range previouslevel {
		if index != firstIndex && index != secondIndex {
			output = append(output, previouslevel[index])
		}
	}

	return output
}

func merge(firstSlice []string, secondSlice []string) []string {
	return orderSet(append(firstSlice, secondSlice...))
}

func orderSet(set []string) []string {
	output := []string{}
	setAsString := toString(set)
	orderedSet:= strings.Split(setAsString, "")
	sort.Strings(orderedSet)
	output = append(output, strings.Join(orderedSet, ""))
	return output
}

func toString(set []string) string {
	output := ""
	for index := range set {
		output = output + set[index]
	}
	return output
}
