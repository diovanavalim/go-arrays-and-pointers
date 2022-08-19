package slices

import (
	"arrays-slices/arrays"
	"fmt"
	"reflect"
)

func CreateSlice() {
	array1, array2, array3 := arrays.CreateArray()

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)

	var slice []string

	slice = append(slice, "Isto", "é", "um", "slice", "de", "palavras")

	sliceString := []string{"Slice", "inicializado", "por", "inferência", "de", "tipos"}

	sliceString = append(sliceString, "Adicionando mais um elemento!")

	slice1 := array1[0:5]
	slice2 := array2[0:10]
	slice3 := array3[0:15]

	fmt.Println(reflect.TypeOf(slice1), reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice2), reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(slice3), reflect.TypeOf(array3))
}
