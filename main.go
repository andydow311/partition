package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := os.Args[1]
	functions := getFunctionAtributes(filePath)
	items := []string{"A", "B" , "C", "D", "E"}
	partitionSpace := getAllPartitions(items)
	//fmt.Println("There are", numberOfStructures(partitionSpace, items), "partitions of", len(items), "items.")
	//logStructureSpace(partitionSpace, items)
	utilitarianOptimal, utilValue  := utilitarianSearch(partitionSpace, functions)
	fmt.Println("utilitarianOptimal is", utilitarianOptimal, "with value", utilValue)
	egaitarianOptimal, egalValue  := egalitarianSearch(partitionSpace, functions)
	fmt.Println("egaitarianOptimal is", egaitarianOptimal, "with value", egalValue)
}

func numberOfStructures(partitionSpace map[int][][][]string, items []string) int {
	number := 0
	for i := 1; i <= len(items); i++ {
		number = number + len(partitionSpace[i])
	}
	return number
}

func logStructureSpace(partitionSpace map[int][][][]string, items []string) {
	for level := 1; level <= len(items); level++ {
		fmt.Println("Level: ", level)
		for _, cs := range partitionSpace[level] {
			fmt.Println("cs: ", cs)
		}
	}
}
