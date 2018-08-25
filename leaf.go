package reflect_db

import (
	"fmt"
	"strings"
)

func leafPrintPath() walkCallback {
	return func(path []string, obj interface{}) bool {
		fmt.Println(strings.Join(append(path, fmt.Sprint(obj)), "."))
		return true
	}
}

func leafSum(sum *int) walkCallback {
	return func(path []string, obj interface{}) bool {
		val, ok := obj.(int)
		if ok {
			*sum += val
		}
		return true
	}
}

func leafAccumulate(avg *int) walkCallback {
	var vals []int

	return func(path []string, obj interface{}) bool {
		val, ok := obj.(int)
		if ok {
			vals = append(vals, val)
		}
		return true
	}
}
