package arrays

func CreateArray() ([5]int, [10]int, [15]int) {
	var array1 [5]int

	array1[0] = 1
	array1[1] = 1
	array1[2] = 2
	array1[3] = 3
	array1[4] = 5

	array2 := [...]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	array3 := [15]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}

	return array1, array2, array3
}
