package main

//Сортировка пузырьком
func bubbleSort(array []int){
	for{
		var temp = true
		for i := 0; i < len(array)-1; i++ {
			
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				temp = false
			}
		}
		if temp {
			break
		}
	}
}