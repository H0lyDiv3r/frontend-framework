package utils

import "reflect"

// func WithoutNulls(arr []internals.Vdom) []internals.Vdom {
// 	var newArr = make([]internals.Vdom, 0, len(arr))
// 	for _, value := range arr {

//			// if value != nil{
//			newArr = append(newArr, value)
//			// }
//		}
//		return newArr
//	}
func WithoutNulls() {}
func FuncExists(funcArray []func(payload ...any), target func(payload ...any)) bool {
	for _, f := range funcArray {
		if reflect.ValueOf(f).Pointer() == reflect.ValueOf(target).Pointer() {
			return true
		}
	}
	return false
}

func IndexOfFunction(arr []func(payload ...any), target func(payload ...any)) int {
	for i, f := range arr {
		if reflect.ValueOf(f).Pointer() == reflect.ValueOf(target).Pointer() {
			return i
		}
	}
	return -1
}
