package main

import (
	"reflect"
	"sort"
)

func getBasePartition(items []string) [][][]string {
	output := [][][]string{}
	basePartition := [][]string{}
	for i := 0; i < len(items); i++ {
		set := []string{items[i]}
		basePartition = append(basePartition, set)
	}
	output = append(output, basePartition)
	return output
}

func getTopPartition(items []string) [][][]string {
	output := [][][]string{}
	topPartition := append([][]string{}, items)
	output = append(output, topPartition)
	return output
}

func getNextLevelPartitions(previousLevelPartitions [][]string) [][][]string {
	nextPartitions := [][][]string{}
	for index, _ := range previousLevelPartitions {
		for nextIndex := index + 1; nextIndex < len(previousLevelPartitions); nextIndex++ {
			newSet := merge(previousLevelPartitions[index], previousLevelPartitions[nextIndex])
			newPartition := deleteElements(previousLevelPartitions, index, nextIndex)
			newPartition = append(newPartition, newSet)
			nextPartitions = append(nextPartitions, newPartition)
		}
	}
	return nextPartitions
}

func getAllPartitions(items []string) map[int][][][]string {
	partitionSpace := make(map[int][][][]string)
	level := len(items)
	partitionSpace[level] = getBasePartition(items)
	for level > 1 {
		previousLevel := partitionSpace[level]
		level = level - 1
		nextLevel := [][][]string{}
		for _, partition := range previousLevel {
			nextLevelParitions := getNextLevelPartitions(partition)
			for thisPartition := 0; thisPartition < len(nextLevelParitions); thisPartition++ {
				orderedPartition := orderPartition(nextLevelParitions[thisPartition])
				if !partitionAlreadyConsidered(nextLevel, orderedPartition) {
					nextLevel = append(nextLevel, orderedPartition)
				}
			}
		}
		partitionSpace[level] = nextLevel
	}
	partitionSpace[level] = getTopPartition(items)
	return partitionSpace
}

func partitionAlreadyConsidered(nextPartitions [][][]string, partition [][]string) bool {
	for _, thisPartition := range nextPartitions {
		if reflect.DeepEqual(thisPartition, partition) {
			return true
		}
	}
	return false
}

func orderPartition(partition [][]string) [][]string {
	output := [][]string{}
	tempStorage := []string{}
	for _, set := range partition {
		tempStorage = append(tempStorage, toString(set))
	}
	sort.Strings(tempStorage)
	for _, x := range tempStorage {
		set := [] string{x}
		output = append(output, set)
	}
	return output
}
