package reflect_db

import (
	"github.com/ElPeque/reflect-db/conv"
	"github.com/ElPeque/reflect-db/types"

	"fmt"
	"strings"
)

func leafPrintPath() types.WalkCallback {
	return func(path []string, obj interface{}) bool {
		fmt.Println(strings.Join(append(path, fmt.Sprint(obj)), "."))
		return true
	}
}

func leafSum(sum *int) types.WalkCallback {
	return func(path []string, obj interface{}) bool {
		*sum += conv.ToInt(obj)
		return true
	}
}

func leafAccumulate(avg *int) types.WalkCallback {
	var vals []int

	return func(path []string, obj interface{}) bool {
		val, ok := obj.(int)
		if ok {
			vals = append(vals, val)
		}
		return true
	}
}
