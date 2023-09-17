package main

func utilitarianSearch(partitionSpace map[int][][][]string, functions []function) ([][]string, float64) {
	//fmt.Println("***utilitarian******")
	optimalValue := 0.0
	optimalPartition := [][]string{}
	for level, partitions := range partitionSpace {
		for _, partition := range partitions {
			if level == 1 {
				optimalValue = UtilitarianPartitionFunction(partition, functions)
				optimalPartition = partition
			} else {
				thisValue := UtilitarianPartitionFunction(partition, functions)
				if optimalValue < thisValue {
					optimalValue = thisValue
					optimalPartition = partition
				}
			}
		}
	}
	return optimalPartition, optimalValue
}

func egalitarianSearch(partitionSpace map[int][][][]string, functions []function) ([][]string, float64) {
	//fmt.Println("***Egalitarian******")
	optimalValue := 0.0
	optimalPartition := [][]string{}
	for level, partitions := range partitionSpace {
		for _, partition := range partitions {
			if level == 1 {
				optimalValue = egalitarianPartitionFunction(partition, functions)
				optimalPartition = partition
			} else {
				thisValue := egalitarianPartitionFunction(partition, functions)
				if optimalValue > thisValue {
					optimalValue = thisValue
					optimalPartition = partition
				}
			}
		}
	}
	return optimalPartition, optimalValue
}
