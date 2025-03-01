package utils

import (
	"go-fe-fwk/types"
)

func WithoutNulls[T types.StringDom | types.Vdom](arr []T) []T {
	var newArr = make([]T, 0, len(arr))
	for _, value := range arr {

		// if value != nil{
		newArr = append(newArr, value)
		// }
	}
	return newArr
}
