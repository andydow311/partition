package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type function struct {
	subsets     string
	synergy     float64
	externality float64
}

func getFunctionAtributes(filePath string) []function {
	output := []function{}

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		subsets := strings.Split(fileScanner.Text(), "	")[0]
		synergy := strings.Split(fileScanner.Text(), "	")[1]
		tsynergy, err := strconv.ParseFloat(synergy, 64)
		if err != nil {
			fmt.Println(err)
		}
		externality := strings.Split(fileScanner.Text(), "	")[2]
		texternality, err := strconv.ParseFloat(externality, 64)
		if err != nil {
			fmt.Println(err)
		}

		function := function{
			subsets,
			tsynergy,
			texternality,
		}
		output = append(output, function)
	}

	return output
}

func UtilitarianPartitionFunction(partition [][]string, functions []function) float64 {
	partitionValue := 0.0

	for _, set := range partition {
		//fmt.Println("set is", set)
		setValue := 0.0
		for _, function := range functions {
			setValue = setValue + computeSetValue(function.subsets,
				function.synergy,
				function.externality,
				set)
			//fmt.Println("sunsets:", function.subsets, "setValue", setValue)
		}
		partitionValue = partitionValue + setValue
	}
	//fmt.Println("partition", partition, " with partitionValue", partitionValue)
	return partitionValue
}

func egalitarianPartitionFunction(partition [][]string, functions []function) float64 {

	partitionValue := 0.0

	for index, set := range partition {
		//fmt.Println("index is", index)
		//fmt.Println("set is", set)
		setValue := 0.0
		for _, function := range functions {
			setValue = setValue + computeSetValue(function.subsets,
				function.synergy,
				function.externality,
				set)
			//	fmt.Println("sunsets:", function.subsets, "setValue", setValue)
		}
		if index == 0 {
			//	fmt.Println("index is 0")
			partitionValue = setValue
		} else {
			if partitionValue > setValue {
				partitionValue = setValue
			}
		}
	}
	//fmt.Println("partition", partition, " with partitionValue", partitionValue)
	return partitionValue
}

func computeSetValue(subsets string, synergy float64, externality float64, set []string) float64 {
	setValue := 0.0
	items := toString(set)
	if strings.Contains(items, subsets) {
		setValue = setValue + synergy
	} else if strings.Contains(subsets, items) {
		setValue = setValue + externality
	}
	return setValue
}
