package main

import "fmt"

func main() {
	items := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	partitionSpace := getAllPartitions(items)
	fmt.Println("There are", numberOfStructures(partitionSpace, items), "partitions of", len(items), "items.")
	logStructureSpace(partitionSpace, items)
}

func numberOfStructures(partitionSpace map[int][][][]string, items []string ) int {
	number :=0
	for i := 1; i <= len(items); i++ {
		number = number + len(partitionSpace[i])
	}
	return number
}

func logStructureSpace(partitionSpace map[int][][][]string, items []string) {
	for level := 1; level <= len(items); level++ {
		fmt.Println("Level: ", level)
		for _, cs := range partitionSpace[level]{
			fmt.Println("cs: ", cs)
		}
	}
}