package main

func utilitarianSearch(partitionSpace map[int][][][]string) ([][]string, float64) {
	optimalValue := 0.0
	optimalPartition := [][]string{}
	for level, partitions := range partitionSpace {
		for _, partition := range partitions {
			if level == 1 {
				optimalValue = UtilitarianPartitionFunction(partition)
				optimalPartition = partition
			} else {

			}
		}
	}
	return optimalPartition, optimalValue
}

func egalitarianSearch(partitionSpace map[int][][][]string) ([][]string, float64) {
	optimalValue := 0.0
	optimalPartition := [][]string{}
	for level, partitions := range partitionSpace {
		for _, partition := range partitions {
			if level == 1 {
				optimalValue = egalitarianPartitionFunction(partition)
				optimalPartition = partition
			} else {

			}
		}
	}
	return optimalPartition, optimalValue
}
